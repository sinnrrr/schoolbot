package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"strconv"
)

var (
	personalInviteButton = tb.InlineButton{
		Unique: "personInvite",
		Text:   l.Gettext("Go to personal chat"),
		URL:    "http://t.me/schoolhelperTheBot/?start=",
	}

	groupInviteButton = tb.InlineButton{
		Unique: "groupInvite",
		Text:   l.Gettext("Add to group"),
		URL:    "http://t.me/schoolhelperTheBot?startgroup=true",
	}

	groupInviteKeys = [][]tb.InlineButton{
		{groupInviteButton},
	}
)

func generatePersonalInviteKeys(groupID int64) [][]tb.InlineButton {
	personalInviteButton.URL += strconv.FormatInt(groupID, 10)

	return [][]tb.InlineButton{
		{personalInviteButton},
	}
}
