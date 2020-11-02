package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"github.com/sinnrrr/schoolbot/db"
	"github.com/sinnrrr/schoolbot/templates"
	"net/http"
	"os"
	"time"
)

func cronAlerts() {
	var (
		alerts            = make(map[int64]map[string]interface{})
		users             = make(map[int64][]map[string]interface{})
	)

	result, err := db.Session.Run(
		"MATCH (a:Alert)-[:BELONGS_TO]->(:Class)<-[:STUDYING_IN]-(s:Student)"+
			"\n"+
			"RETURN a, s",
		map[string]interface{}{},
	)
	if err != nil {
		panic(err)
	}

	for result.Next() {
		alertID := result.Record().Values()[0].(neo4j.Node).Id()

		alerts[alertID] = result.Record().Values()[0].(neo4j.Node).Props()
		users[alertID] = append(users[alertID], result.Record().Values()[1].(neo4j.Node).Props())
	}

	for alertID, alert := range alerts {
		for _, user := range users[alertID] {
			currentDate := time.Now()
			alertDate := time.Unix(alert["date"].(int64), 0)

			if alertDate.Hour() == currentDate.Hour() && alertDate.Minute() == currentDate.Minute() {
				fmt.Println(templates.GenerateMessageURL(
					os.Getenv("BOT_TOKEN"),
					user["id"].(int),
					templates.GenerateCronAlert(alert["content"].(string)),
				))

				_, err := http.Get(
					templates.GenerateMessageURL(
						os.Getenv("BOT_TOKEN"),
						user["id"].(int),
						templates.GenerateCronAlert(alert["content"].(string)),
					),
				)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
