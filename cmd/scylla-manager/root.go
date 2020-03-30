// Copyright (C) 2017 ScyllaDB

package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
	"text/template"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/gops/agent"
	"github.com/pkg/errors"
	"github.com/scylladb/go-log"
	"github.com/scylladb/go-log/gocqllog"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/migrate"
	"github.com/scylladb/mermaid"
	"github.com/scylladb/mermaid/internal/fsutil"
	"github.com/scylladb/mermaid/schema/cql"
	"github.com/spf13/cobra"
)

var (
	cfgConfigFile []string
	cfgVersion    bool
)

func init() {
	rootCmd.Flags().StringSliceVarP(&cfgConfigFile, "config-file", "c", []string{"/etc/scylla-manager/scylla-manager.yaml"}, "configuration file `path`")
	rootCmd.Flags().BoolVar(&cfgVersion, "version", false, "print product version and exit")
}

var rootCmd = &cobra.Command{
	Use:           "scylla-manager",
	Short:         "Scylla Manager server",
	Args:          cobra.NoArgs,
	SilenceUsage:  true,
	SilenceErrors: true,

	RunE: func(cmd *cobra.Command, args []string) (runError error) {
		// print version and return
		if cfgVersion {
			fmt.Fprintf(cmd.OutOrStdout(), "%s\n", mermaid.Version())
			return
		}

		// read configuration
		config, err := newConfigFromFile(cfgConfigFile...)
		if err != nil {
			runError = errors.Wrapf(err, "configuration %q", cfgConfigFile)
			fmt.Fprintf(cmd.OutOrStderr(), "%s\n", runError)
			return
		}
		if err := config.validate(); err != nil {
			runError = errors.Wrapf(err, "configuration %q", cfgConfigFile)
			fmt.Fprintf(cmd.OutOrStderr(), "%s\n", runError)
			return
		}

		// launch gops agent
		if err := agent.Listen(agent.Options{Addr: config.Gops, ShutdownCleanup: false}); err != nil {
			return errors.Wrapf(err, "gops agent startup")
		}
		defer agent.Close()

		// get a base context
		ctx := log.WithNewTraceID(context.Background())

		// create logger
		logger, err := logger(config)
		if err != nil {
			return errors.Wrapf(err, "logger")
		}
		defer func() {
			if runError != nil {
				logger.Error(ctx, "Bye", "error", runError)
			} else {
				logger.Info(ctx, "Bye")
			}
			logger.Sync() // nolint
		}()
		logger.Info(ctx, "Using config", "config", obfuscatePasswords(config))

		// set gocql logger
		gocql.Logger = gocqllog.StdLogger{
			BaseCtx: ctx,
			Logger:  logger.Named("gocql"),
		}

		// wait for database
		if err := waitForDatabase(ctx, config, logger); err != nil {
			return err
		}

		// create manager keyspace
		logger.Info(ctx, "Using keyspace",
			"keyspace", config.Database.Keyspace,
			"template", config.Database.KeyspaceTplFile,
		)
		if err := createKeyspace(config); err != nil {
			return errors.Wrapf(err, "database")
		}

		// migrate schema
		logger.Info(ctx, "Migrating schema", "dir", config.Database.MigrateDir)
		if err := migrateSchema(config, logger); err != nil {
			return errors.Wrapf(err, "database migration")
		}
		logger.Info(ctx, "Migrating schema done")

		// start server
		s, err := newServer(config, logger)
		if err != nil {
			return errors.Wrapf(err, "server init")
		}
		if err := s.initDatabase(ctx); err != nil {
			return errors.Wrapf(err, "database init")
		}
		if err := s.startServices(ctx); err != nil {
			return errors.Wrapf(err, "server start")
		}
		s.startHTTPServers(ctx)
		defer s.close()

		logger.Info(ctx, "Service started")

		// wait signal
		signalCh := make(chan os.Signal, 1)
		signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
		select {
		case err := <-s.errCh:
			if err != nil {
				logger.Error(ctx, "Server error", "error", err)
			}
		case sig := <-signalCh:
			{
				logger.Info(ctx, "Received signal", "signal", sig)
			}
		}

		// close
		s.shutdownServers(ctx, 30*time.Second)

		return
	},
}

func logger(config *serverConfig) (log.Logger, error) {
	if config.Logger.Development {
		return log.NewDevelopmentWithLevel(config.Logger.Level), nil
	}

	if config.Logger.Mode != log.StderrMode {
		f, err := redirectStdErrAndStdOutToFile()
		if err != nil {
			return log.NopLogger, err
		}
		defer f.Close()
	}

	return log.NewProduction(log.Config{
		Mode:  config.Logger.Mode,
		Level: config.Logger.Level,
	})
}

func redirectStdErrAndStdOutToFile() (*os.File, error) {
	p := path.Join(fsutil.HomeDir(), "stdout")

	f, err := os.OpenFile(p, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		return nil, err
	}

	os.Stdout = f
	os.Stderr = f
	return f, nil
}

func waitForDatabase(ctx context.Context, config *serverConfig, logger log.Logger) error {
	const (
		wait        = 5 * time.Second
		maxAttempts = 60
	)

	for i := 0; i < maxAttempts; i++ {
		if _, err := tryConnect(config); err != nil {
			logger.Info(ctx, "Could not connect to database",
				"sleep", wait,
				"error", err,
			)
			time.Sleep(wait)
		} else {
			return nil
		}
	}

	return errors.New("could not connect to database, max attempts reached")
}

func tryConnect(config *serverConfig) (string, error) {
	for _, host := range config.Database.Hosts {
		conn, err := net.Dial("tcp", net.JoinHostPort(host, "9042"))
		if conn != nil {
			conn.Close()
		}
		if err == nil {
			return host, nil
		}
	}
	return "", errors.Errorf("tried all hosts %s", config.Database.Hosts)
}

func createKeyspace(config *serverConfig) error {
	c := gocqlConfig(config)
	c.Keyspace = "system"
	c.Timeout = config.Database.MigrateTimeout
	c.MaxWaitSchemaAgreement = config.Database.MigrateMaxWaitSchemaAgreement

	session, err := c.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()

	var cnt int
	q := session.Query("SELECT COUNT(keyspace_name) FROM system_schema.keyspaces WHERE keyspace_name = ?").Bind(config.Database.Keyspace)
	if err := q.Scan(&cnt); err != nil {
		return err
	}
	if cnt == 1 {
		return nil
	}

	// Auto upgrade replication factor if needed. RF=1 with multiple hosts means
	// data loss when one of the nodes is down. This is understood with a single
	// node deployment but must be avoided if we have more nodes.
	if config.Database.ReplicationFactor == 1 {
		var peers int
		q := session.Query("SELECT COUNT(*) FROM system.peers")
		if err := q.Scan(&peers); err != nil {
			return err
		}
		if peers > 0 {
			rf := peers + 1
			if rf > 3 {
				rf = 3
			}
			config.Database.ReplicationFactor = rf
		}
	}

	stmt, err := readKeyspaceTplFile(config)
	if err != nil {
		return err
	}

	return gocqlx.Query(session.Query(stmt), nil).ExecRelease()
}

func readKeyspaceTplFile(config *serverConfig) (stmt string, err error) {
	b, err := ioutil.ReadFile(config.Database.KeyspaceTplFile)
	if err != nil {
		return "", errors.Wrapf(err, "could not read file %s", config.Database.KeyspaceTplFile)
	}

	t := template.New("")
	if _, err := t.Parse(string(b)); err != nil {
		return "", errors.Wrapf(err, "template error file %s", config.Database.KeyspaceTplFile)
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, config.Database); err != nil {
		return "", errors.Wrapf(err, "template error file %s", config.Database.KeyspaceTplFile)
	}

	return buf.String(), err
}

func migrateSchema(config *serverConfig, logger log.Logger) error {
	host, err := tryConnect(config)
	if err != nil {
		return err
	}

	c := gocqlConfig(config)
	c.Timeout = config.Database.MigrateTimeout
	c.MaxWaitSchemaAgreement = config.Database.MigrateMaxWaitSchemaAgreement
	c.DisableInitialHostLookup = true
	c.Hosts = []string{host}

	session, err := c.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()

	cql.Logger = logger
	migrate.Callback = cql.MigrateCallback

	return migrate.Migrate(context.Background(), session, config.Database.MigrateDir)
}

func gocqlConfig(config *serverConfig) *gocql.ClusterConfig {
	c := gocql.NewCluster(config.Database.Hosts...)

	// Chose consistency level, for a single node deployments use ONE, for
	// multi-dc deployments use LOCAL_QUORUM, otherwise use QUORUM.
	if config.Database.LocalDC != "" {
		c.Consistency = gocql.LocalQuorum
	} else if config.Database.ReplicationFactor == 1 {
		c.Consistency = gocql.One
	} else {
		c.Consistency = gocql.Quorum
	}

	c.Keyspace = config.Database.Keyspace
	c.Timeout = config.Database.Timeout

	// ssl
	if config.Database.SSL {
		c.SslOpts = &gocql.SslOptions{
			CaPath:                 config.SSL.CertFile,
			CertPath:               config.SSL.UserCertFile,
			KeyPath:                config.SSL.UserKeyFile,
			EnableHostVerification: config.SSL.Validate,
		}
	}

	// authentication
	if config.Database.User != "" {
		c.Authenticator = gocql.PasswordAuthenticator{
			Username: config.Database.User,
			Password: config.Database.Password,
		}
	}

	// enable token aware host selection policy
	if config.Database.TokenAware {
		fallback := gocql.RoundRobinHostPolicy()
		if config.Database.LocalDC != "" {
			fallback = gocql.DCAwareRoundRobinPolicy(config.Database.LocalDC)
		}
		c.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(fallback)
	}

	return c
}

func obfuscatePasswords(config *serverConfig) serverConfig {
	cfg := *config
	cfg.Database.Password = strings.Repeat("*", len(cfg.Database.Password))
	return cfg
}