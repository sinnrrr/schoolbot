package main

import tb "gopkg.in/tucnak/telebot.v2"

var (
	personalInviteButton = tb.InlineButton{
		Text:   "Go",
		Unique: "personInvite",
		URL:    "http://t.me/schoolhelperTheBot/?start=1432",
	}

	groupInviteButton = tb.InlineButton{
		Text: "Add to group",
		Unique: "groupInvite",
		URL: "http://t.me/schoolhelperTheBot?startgroup=true",
	}

	personalInviteKeys = [][]tb.InlineButton{
		{personalInviteButton},
	}

	groupInviteKeys = [][]tb.InlineButton{
		{groupInviteButton},
	}
)
