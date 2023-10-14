package postgres

import (
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

type Option func(cfg *config) error

func WithUsername(username string) Option {
	return func(cfg *config) error {
		if len(username) != 0 {
			cfg.username = username
		}

		return nil
	}
}

func WithPassword(password string) Option {
	return func(cfg *config) error {
		if len(password) != 0 {
			cfg.password = password
		}

		return nil
	}
}

func WithDatabaseName(name string) Option {
	return func(cfg *config) error {
		if len(name) != 0 {
			cfg.database = name
		}

		return nil
	}
}

func WithStaticPort(port int) Option {
	return func(cfg *config) error {
		if port < 0 && port > (1<<32-1) {
			return fmt.Errorf("%w: actual %d", ErrPortOutOfRange, port)
		}

		cfg.exposePort = port

		return nil
	}
}

func WithQuery(query map[string]string) Option {
	return func(cfg *config) error {
		cfg.query = query

		return nil
	}
}

func WithRepository(repository string) Option {
	return func(cfg *config) error {
		if len(repository) != 0 {
			cfg.repository = repository
		}

		return nil
	}
}

func WithTag(tag string) Option {
	return func(cfg *config) error {
		if len(tag) != 0 {
			cfg.tag = tag
		}

		return nil
	}
}

func WithAuthConfiguration(auth docker.AuthConfiguration) Option {
	return func(cfg *config) error {
		switch {
		case len(auth.IdentityToken) > 0 && len(auth.ServerAddress) > 0,
			len(auth.Username) > 0 && len(auth.Password) > 0 && len(auth.ServerAddress) > 0:
			cfg.auth = auth

			return nil
		default:
			return ErrInvalidAuthParameters
		}
	}
}

func WithNetwork(network *dockertest.Network) Option {
	return func(cfg *config) error {
		if network == nil {
			return ErrNotFoundNetwork
		}

		for _, item := range cfg.networks {
			if item.Network.ID == network.Network.ID {
				return nil
			}
		}

		cfg.networks = append(cfg.networks, network)

		return nil
	}
}
