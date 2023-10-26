package postgres

import (
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"strconv"
	"time"
)

type config struct {
	username    string
	password    string
	database    string
	query       map[string]string
	host        string
	exposePort  int
	portBinding int
	repository  string
	tag         string
	auth        docker.AuthConfiguration
	networks    []*dockertest.Network
	maxWait     time.Duration
}

var defaultConfig = config{
	username:   "postgres",
	password:   "password",
	database:   "postgres",
	host:       "localhost",
	exposePort: 5432,
	query: map[string]string{
		"sslmode": "disable",
	},
	repository: "postgres",
	tag:        "latest",
	maxWait:    time.Minute,
}

func env(key, value string) string {
	return fmt.Sprintf("%s=%s", key, value)
}

func (c *config) exposes() []string {
	return []string{
		fmt.Sprintf("%d/tcp", c.exposePort),
	}
}

func (c *config) bindings() map[docker.Port][]docker.PortBinding {
	if c.portBinding != 0 {
		return map[docker.Port][]docker.PortBinding{
			docker.Port(
				fmt.Sprintf("%d/tcp", c.exposePort),
			): {
				{
					HostIP:   c.host,
					HostPort: strconv.Itoa(c.portBinding),
				},
			},
		}
	}

	return nil
}
