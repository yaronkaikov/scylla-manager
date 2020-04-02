// Copyright (C) 2017 ScyllaDB

package restapi

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"github.com/scylladb/mermaid/pkg/service/backup"
	"github.com/scylladb/mermaid/pkg/service/scheduler"
)

type backupHandler struct {
	svc      BackupService
	schedSvc SchedService
}

func newBackupHandler(services Services) *chi.Mux {
	m := chi.NewMux()
	h := backupHandler{
		svc:      services.Backup,
		schedSvc: services.Scheduler,
	}

	m.Use(
		h.locationsCtx,
		h.listFilterCtx,
	)
	m.Get("/", h.list)
	m.Get("/files", h.listFiles)

	return m
}

func (h backupHandler) locationsCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			locations []backup.Location
			err       error
		)

		// Read locations from the request
		if v := r.FormValue("locations"); v != "" {
			for _, v := range r.Form["locations"] {
				var l backup.Location
				if err := l.UnmarshalText([]byte(v)); err != nil {
					respondBadRequest(w, r, err)
					return
				}
				locations = append(locations, l)
			}
		}

		// Fallback read locations from scheduler
		if len(locations) == 0 {
			locations, err = h.extractLocations(r)
			if err != nil {
				respondError(w, r, err)
				return
			}
		}

		// Report error if no locations can be found
		if len(locations) == 0 {
			respondBadRequest(w, r, errors.New("missing locations"))
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, ctxBackupLocations, locations)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h backupHandler) extractLocations(r *http.Request) ([]backup.Location, error) {
	tasks, err := h.schedSvc.ListTasks(r.Context(), mustClusterIDFromCtx(r), scheduler.BackupTask)
	if err != nil {
		return nil, err
	}
	properties := make([]json.RawMessage, len(tasks))
	for _, t := range tasks {
		if t.Enabled {
			properties = append(properties, t.Properties)
		}
	}
	return h.svc.ExtractLocations(r.Context(), properties), nil
}

func (h backupHandler) mustLocationsFromCtx(r *http.Request) []backup.Location {
	v, ok := r.Context().Value(ctxBackupLocations).([]backup.Location)
	if !ok {
		panic("missing locations in context")
	}
	return v
}

func (h backupHandler) listFilterCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var filter = backup.ListFilter{
			Keyspace:    r.Form["keyspace"],
			SnapshotTag: r.FormValue("snapshot_tag"),
		}

		c := mustClusterFromCtx(r)
		if v := r.FormValue("cluster_id"); v != "" {
			if v == c.Name || v == c.ID.String() {
				filter.ClusterID = c.ID
			} else if err := filter.ClusterID.UnmarshalText([]byte(v)); err != nil {
				respondBadRequest(w, r, errors.Wrap(err, "invalid cluster_id"))
				return
			}
		}

		if v := r.FormValue("min_date"); v != "" {
			if err := filter.MinDate.UnmarshalText([]byte(v)); err != nil {
				respondBadRequest(w, r, errors.Wrap(err, "invalid min_date"))
				return
			}
		}
		if v := r.FormValue("max_date"); v != "" {
			if err := filter.MaxDate.UnmarshalText([]byte(v)); err != nil {
				respondBadRequest(w, r, errors.Wrap(err, "invalid max_date"))
				return
			}
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, ctxBackupListFilter, filter)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (h backupHandler) mustListFilterFromCtx(r *http.Request) backup.ListFilter {
	v, ok := r.Context().Value(ctxBackupListFilter).(backup.ListFilter)
	if !ok {
		panic("missing filter in context")
	}
	return v
}

func (h backupHandler) list(w http.ResponseWriter, r *http.Request) {
	v, err := h.svc.List(
		r.Context(),
		mustClusterIDFromCtx(r),
		h.mustLocationsFromCtx(r),
		h.mustListFilterFromCtx(r),
	)
	if err != nil {
		respondError(w, r, err)
		return
	}

	render.Respond(w, r, v)
}

func (h backupHandler) listFiles(w http.ResponseWriter, r *http.Request) {
	v, err := h.svc.ListFiles(
		r.Context(),
		mustClusterIDFromCtx(r),
		h.mustLocationsFromCtx(r),
		h.mustListFilterFromCtx(r),
	)
	if err != nil {
		respondError(w, r, err)
		return
	}

	render.Respond(w, r, v)
}