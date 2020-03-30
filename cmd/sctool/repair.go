// Copyright (C) 2017 ScyllaDB

package main

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/scylladb/mermaid/internal/duration"
	"github.com/scylladb/mermaid/mermaidclient"
	"github.com/spf13/cobra"
)

type tokenRangesKind string

const (
	pr  tokenRangesKind = "pr"
	npr tokenRangesKind = "npr"
	all tokenRangesKind = "all"
)

func (r tokenRangesKind) String() string {
	return string(r)
}

func (r *tokenRangesKind) Set(s string) error {
	switch tokenRangesKind(s) {
	case pr:
		*r = pr
	case npr:
		*r = npr
	case all:
		*r = all
	default:
		return errors.New("valid values are: pr, npr, all")
	}
	return nil
}

func (tokenRangesKind) Type() string {
	return "token ranges"
}

var repairTokenRanges = tokenRangesKind("")

var repairCmd = &cobra.Command{
	Use:   "repair",
	Short: "Schedule repair",
	RunE: func(cmd *cobra.Command, args []string) error {
		props := make(map[string]interface{})

		t := &mermaidclient.Task{
			Type:       "repair",
			Enabled:    true,
			Schedule:   new(mermaidclient.Schedule),
			Properties: props,
		}

		f := cmd.Flag("start-date")
		startDate, err := mermaidclient.ParseStartDate(f.Value.String())
		if err != nil {
			return printableError{err}
		}
		t.Schedule.StartDate = startDate

		i, err := cmd.Flags().GetString("interval")
		if err != nil {
			return printableError{err}
		}
		if _, err := duration.ParseDuration(i); err != nil {
			return printableError{err}
		}
		t.Schedule.Interval = i

		t.Schedule.NumRetries, err = cmd.Flags().GetInt64("num-retries")
		if err != nil {
			return printableError{err}
		}

		if f = cmd.Flag("keyspace"); f.Changed {
			keyspace, err := cmd.Flags().GetStringSlice("keyspace")
			if err != nil {
				return printableError{err}
			}
			props["keyspace"] = unescapeFilters(keyspace)
		}

		if f = cmd.Flag("dc"); f.Changed {
			dc, err := cmd.Flags().GetStringSlice("dc")
			if err != nil {
				return printableError{err}
			}
			props["dc"] = unescapeFilters(dc)
		}

		if f = cmd.Flag("host"); f.Changed {
			host, err := cmd.Flags().GetString("host")
			if err != nil {
				return printableError{err}
			}
			props["host"] = host
		}

		if f = cmd.Flag("with-hosts"); f.Changed {
			hosts, err := cmd.Flags().GetStringSlice("with-hosts")
			if err != nil {
				return printableError{err}
			}
			props["with_hosts"] = hosts
		}

		if f = cmd.Flag("token-ranges"); f.Changed {
			if !cmd.Flag("host").Changed && !cmd.Flag("with-hosts").Changed {
				return printableError{errors.New("token-ranges is only available with \"host\" and \"with-hosts\" flags")}
			}
			props["token_ranges"] = repairTokenRanges.String()
		}

		failFast, err := cmd.Flags().GetBool("fail-fast")
		if err != nil {
			return printableError{err}
		}
		if failFast {
			t.Schedule.NumRetries = 0
			props["fail_fast"] = true
		}

		force, err := cmd.Flags().GetBool("force")
		if err != nil {
			return printableError{err}
		}

		dryRun, err := cmd.Flags().GetBool("dry-run")
		if err != nil {
			return printableError{err}
		}
		if dryRun {
			fmt.Fprintf(cmd.OutOrStderr(), "NOTICE: dry run mode, repair is not scheduled\n\n")
			res, err := client.GetTarget(ctx, cfgCluster, t)
			if err != nil {
				return printableError{err}
			}
			return res.Render(cmd.OutOrStdout())
		}

		id, err := client.CreateTask(ctx, cfgCluster, t, force)
		if err != nil {
			return printableError{err}
		}

		fmt.Fprintln(cmd.OutOrStdout(), mermaidclient.TaskJoin("repair", id))

		return nil
	},
}

func init() {
	cmd := repairCmd
	withScyllaDocs(cmd, "/sctool/#repair")
	register(repairCmd, rootCmd)

	fs := cmd.Flags()
	fs.StringSliceP("keyspace", "K", nil, "comma-separated `list` of keyspace/tables glob patterns, e.g. keyspace,!keyspace.table_prefix_*")
	fs.StringSlice("dc", nil, "comma-separated `list` of data centers glob patterns, e.g. dc1,!otherdc*")
	fs.String("host", "", "host to repair, by default all hosts are repaired")
	fs.StringSlice("with-hosts", nil, "comma-separated `list` of hosts to repair with")
	fs.Var(&repairTokenRanges, "token-ranges", "`kind` of token-ranges valid values are: pr, npr, all")
	fs.Bool("fail-fast", false, "stop repair on first error")
	fs.Bool("force", false, "force repair to skip database validation and schedule repair even if there no matching keyspaces/tables")
	fs.Bool("dry-run", false, "validate and print repair information without scheduling a repair")
	taskInitCommonFlags(fs)
}

// accommodate for escaping of bash expansions, we can safely remove '\'
// as it's not a valid char in keyspace or table name
func unescapeFilters(strs []string) []string {
	for i := range strs {
		strs[i] = strings.Replace(strs[i], "\\", "", -1)
	}
	return strs
}