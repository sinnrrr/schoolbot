package main

import (
	"github.com/sinnrrr/schoolbot/db"
	"golang.org/x/text/language"
	tb "gopkg.in/tucnak/telebot.v2"
	"io/ioutil"
)

var supportedLocale []language.Tag

func syncSupportedLanguage() error {
	files, err := ioutil.ReadDir("locale")
	if err != nil {
		return err
	}

	for _, f := range files {
		tag, err := language.Parse(f.Name())
		if err != nil {
			return err
		}

		supportedLocale = append(supportedLocale, tag)
	}

	return nil
}

func syncUserSession(student *tb.User, dialogueState int) (map[string]interface{}, error) {
	session, err := db.StudentSession(student.ID)
	if err != nil {
		return nil, err
	}

	if session["language_code"].(string) != student.LanguageCode {
		l.SetLanguage(student.LanguageCode)

		var data = map[string]interface{}{
			"tg_id": student.ID,
			"language_code": student.LanguageCode,
		}
		if dialogueState >= 0 {
			data["dialogue_state"] = dialogueState
		}

		result, err := db.UpdateStudentSession(data)
		if err != nil {
			return nil, err
		}

		return result, nil
	}

	return session, nil
}
