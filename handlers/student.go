package handlers

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/sinnrrr/schoolbot/db"
	tb "gopkg.in/tucnak/telebot.v2"
)

func CreateStudent(user *tb.User, classID int64) (neo4j.Node, error) {
	var student neo4j.Node

	result, err := db.Session.Run(
		"MATCH (c:Class { id: $class_id })"+
			"\n"+
			"CREATE (s:Student {"+
			"id: $id,"+
			"first_name: $first_name,"+
			"last_name: $last_name,"+
			"username: $username,"+
			"language_code: $language_code"+
			"state: $dialogue_state"+
			"})-[:STUDYING_IN]->(c)"+
			"\n"+
			"RETURN s",
		map[string]interface{}{
			"id":            user.ID,
			"first_name":    user.FirstName,
			"last_name":     user.LastName,
			"username":      user.Username,
			"language_code": user.LanguageCode,
			"class_id":      classID,
			"state":         0,
		},
	)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		student = result.Record().GetByIndex(0).(neo4j.Node)
	}

	return student, nil
}

func GetDialogueState(studentID int) (int8, error) {
	var dialogueState int8

	result, err := db.Session.Run(
		"MATCH (s:Student { id: $id } RETURN s.state",
		map[string]interface{}{
			"id": studentID,
		},
	)
	if err != nil {
		return -1, err
	}

	for result.Next() {
		dialogueState = result.Record().GetByIndex(0).(int8)
	}

	return dialogueState, nil
}

func SetDialogueState(studentID int, state int8) error {
	_, err := db.Session.Run(
		"MATCH (s:Student { id: $id } SET s.state = $dialogue_state",
		map[string]interface{}{
			"id": studentID,
			"dialogue_state": state,
		},
	)
	if err != nil {
		return err
	}

	return nil
}
