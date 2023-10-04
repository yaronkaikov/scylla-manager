// Copyright (C) 2017 ScyllaDB

//go:build all || integration
// +build all integration

package dynamoping

import (
	"context"
	"github.com/scylladb/scylla-manager/v3/pkg/testutils/testconfig"
	"testing"
	"time"

	_ "github.com/scylladb/scylla-manager/v3/pkg/testutils"
)

func TestPingIntegration(t *testing.T) {
	config := Config{
		Addr:    "http://" + testconfig.ManagedClusterHost() + ":8000",
		Timeout: 250 * time.Millisecond,
	}

	t.Run("simple", func(t *testing.T) {
		d, err := SimplePing(context.Background(), config)
		if err != nil {
			t.Error(err)
		}
		t.Logf("simplePing() = %s", d)
	})

	t.Run("query", func(t *testing.T) {
		d, err := QueryPing(context.Background(), config)
		if err != nil {
			t.Error(err)
		}
		t.Logf("QueryPing() = %s", d)
	})
}
