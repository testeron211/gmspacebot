package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	Ver = "0.12"

	RateErr = "Нельзя использовать бота так часто!"

	author = &discordgo.MessageEmbedAuthor{
		IconURL: "https://troit5ky.space/bot/assets/img/bot.gif",
		URL:     "https://troit5ky.space",
		Name:    "GM.space",
	}
	footer = &discordgo.MessageEmbedFooter{
		Text: Ver,
	}
	embed = &discordgo.MessageEmbed{
		Author:    author,
		Timestamp: time.Now().Format(time.RFC3339),
		Footer:    footer,
		Color:     4868863,
	}
)
