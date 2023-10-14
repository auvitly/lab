package postgres

import (
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"net"
	"net/url"
)

type Database struct {
	// config - database configuration.
	config config
	// resource - container.
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

func (d *Database) Init() (err error) {
	if d.resource != nil {
		return nil
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		return fmt.Errorf("dockertest.NewPool: %w", err)
	}

	if err = pool.Client.Ping(); err != nil {
		return fmt.Errorf("pool.Client.Ping: %w", err)
	}

	d.resource, err = pool.RunWithOptions(&dockertest.RunOptions{
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

	return nil
}

func (d *Database) Close() error {
	if d.resource == nil {
		return nil
	}

	return d.resource.Close()
}
