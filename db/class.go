package db

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func CreateClass(classID int64, name string) (map[string]interface{}, error) {
	var class map[string]interface{}

	result, err := Session.Run(
		"MERGE (c:Class { tg_id: $tg_id, name: $name }) RETURN c",
		map[string]interface{}{
			"tg_id": classID,
			"name":  name,
		},
	)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		class = result.Record().GetByIndex(0).(neo4j.Node).Props()
	}

	return class, result.Err()
}

func QueryClassStudents(studentID int) ([]map[string]interface{}, error) {
	var students []map[string]interface{}

	result, err := Session.Run(
		"MATCH (:Student { tg_id: $tg_id })-[:STUDYING_IN]->(c:Class)"+
			"\n"+
			"MATCH (s:Student)-[:STUDYING_IN]->(c)"+
			"\n"+
			"RETURN s",
		map[string]interface{}{
			"tg_id": studentID,
		},
	)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		students = append(students, result.Record().Values()[0].(neo4j.Node).Props())
	}

	return students, result.Err()
}
