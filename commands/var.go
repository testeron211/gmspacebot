package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	author = &discordgo.MessageEmbedAuthor{
		URL:     "https://gm.troit5ky.ru",
		Name:    "GM.space",
		IconURL: "https://troit5ky.ru/assets/img/avatar.png",
	}

	footer = &discordgo.MessageEmbedFooter{
		Text:         "",
		IconURL:      "",
		ProxyIconURL: "",
	}

	embed = &discordgo.MessageEmbed{
		Author:    author,
		Timestamp: time.Now().Format(time.RFC3339),
	}
)
