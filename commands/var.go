package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	author = &discordgo.MessageEmbedAuthor{
		URL:     "https://gm.troit5ky.ru",
		Name:    "GM.space",
		IconURL: "https://troit5ky.ru/assets/img/avatar.png",
	}

	footer = &discordgo.MessageEmbedFooter{
		Text: "made by troit5ky.ru",
	}

	embed = &discordgo.MessageEmbed{
		Author: author,
	}
)
