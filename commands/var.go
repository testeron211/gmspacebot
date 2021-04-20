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

	embed = &discordgo.MessageEmbed{
		Author: author,
	}
)
