package postgres

import (
	"database/sql"
	"fmt"
	"net"
	"net/url"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

type Database struct {
	// pool - containers pool.
	pool *dockertest.Pool
	// config - database configuration.
	config config
	// resource - containers container.
	resource *dockertest.Resource
	// DSN - DB connection string.
	DSN *url.URL
}

func NewDatabase(options ...Option) (*Database, error) {
	var database = &Database{
		config: defaultConfig,
	}

	for _, option := range options {
		if err := option(&database.config); err != nil {
			return nil, fmt.Errorf("%w: %s", ErrInvalidOption, err.Error())
		}
	}

	return database, nil
}

func MustNewDatabase(options ...Option) *Database {
	database, err := NewDatabase(options...)
	if err != nil {
		panic(err)
	}

	return database
}

func (d *Database) Start() (err error) {
	if d.resource != nil {
		return nil
	}

	if err = d.runDatabase(); err != nil {
		return err
	}

	return nil
}

func (d *Database) runDatabase() (err error) {
	d.pool, err = dockertest.NewPool("")
	if err != nil {
		return fmt.Errorf("dockertest.NewPool: %w", err)
	}

	if err = d.pool.Client.Ping(); err != nil {
		return fmt.Errorf("pool.Client.Ping: %w", err)
	}

	d.pool.MaxWait = d.config.maxWait

	d.resource, err = d.pool.RunWithOptions(&dockertest.RunOptions{
		Hostname:     d.config.host,
		Repository:   d.config.repository,
		Tag:          d.config.tag,
		ExposedPorts: d.config.exposes(),
		Env: []string{
			env("POSTGRES_USER", d.config.username),
			env("POSTGRES_DATABASE", d.config.database),
			env("POSTGRES_PASSWORD", d.config.password),
		},
		PortBindings: d.config.bindings(),
		Auth:         d.config.auth,
		Networks:     d.config.networks,
	}, func(hc *docker.HostConfig) {
		hc.AutoRemove = true
		hc.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		return fmt.Errorf("pool.RunWithOptions: %w", err)
	}

	d.buildDSN()

	if err = d.pool.Retry(d.retry); err != nil {
		return fmt.Errorf("pool.Retry: %w", err)
	}

	return nil
}

func (d *Database) buildDSN() {
	d.DSN = new(url.URL)

	d.DSN.Scheme = "postgres"
	d.DSN.User = url.UserPassword(d.config.username, d.config.password)
	d.DSN.Host = net.JoinHostPort(d.config.host, d.resource.GetPort(fmt.Sprintf("%d/tcp", d.config.exposePort)))
	d.DSN.Path = d.config.database

	var values = make(url.Values)

	for key, value := range d.config.query {
		values.Add(key, value)
	}

	d.DSN.RawQuery = values.Encode()
}

func (d *Database) retry() error {
	db, err := sql.Open("postgres", d.DSN.String())
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("db.Ping: %w", err)
	}

	return nil
}

func (d *Database) Close() error {
	if d.resource == nil {
		return nil
	}

	if err := d.resource.Close(); err != nil {
		return err
	}

	d.DSN = nil
	d.resource = nil

	return nil
}
