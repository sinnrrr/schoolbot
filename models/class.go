package models

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/sinnrrr/schoolbot/db"
)

type Class struct {
	ID   int64 `json:"id"`
	Name string `json:"name"`
}

func (Class) Create(id int64, name string) (neo4j.Node, error) {
	var class neo4j.Node

	result, err := db.Session.Run(
		"CREATE (c:Class { id: $id, name: $name }) RETURN c",
		map[string]interface{}{
			"id": id,
			"name": name,
		},
	)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		class = result.Record().GetByIndex(0).(neo4j.Node)
	}

	return class, nil
}
