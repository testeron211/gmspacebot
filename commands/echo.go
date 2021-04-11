package commands

import (
	"github.com/Lukaesebrot/dgc"
)

func Echo(ctx *dgc.Ctx) {
	embed.Color = 4915018
	embed.Title = "Echo"

	if ctx.Arguments.Amount() < 1 {
		embed.Description = "Введи аргументы!"
		embed.Color = 16730698
		ctx.RespondEmbed(embed)
	} else {
		embed.Description = "```\n" + ctx.Arguments.Raw() + "\n```"
		ctx.RespondEmbed(embed)
	}

}
