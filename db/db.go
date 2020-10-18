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

func Init() {
	Driver, err = neo4j.NewDriver(config.URI, config.Auth, config.DB(neo4j.INFO))
	if err != nil {
		panic(err)
	}
	defer Driver.Close()

	Session, err = Driver.NewSession(config.Session)
	if err != nil {
		panic(err)
	}
	defer Session.Close()
}
