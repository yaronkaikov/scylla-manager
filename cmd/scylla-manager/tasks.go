// Copyright (C) 2017 ScyllaDB

package main

import (
	"github.com/scylladb/mermaid/internal/timeutc"
	"github.com/scylladb/mermaid/sched"
	"github.com/scylladb/mermaid/uuid"
)

func autoRepairTask(clusterID uuid.UUID) *sched.Task {
	return &sched.Task{
		ClusterID: clusterID,
		Type:      sched.RepairTask,
		Enabled:   true,
		Sched: sched.Schedule{
			IntervalDays: 7,
			StartDate:    timeutc.TodayMidnight(),
			NumRetries:   3,
		},
		Properties: []byte{'{', '}'},
	}
}
