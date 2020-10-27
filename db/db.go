package db

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/sinnrrr/schoolbot/config"
)

var (
	err     error

	Driver  neo4j.Driver
	Session neo4j.Session
)

func Init() error {
	Driver, err = neo4j.NewDriver(config.URI, config.Auth, config.DB())
	if err != nil {
		return err
	}

	Session, err = Driver.NewSession(config.Session)
	return err
}
