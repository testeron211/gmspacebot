package commands

import (
	"log"

	"github.com/Lukaesebrot/dgc"
)

func connectionTimeOut(ctx *dgc.Ctx, err error) {
	log.Println(err)
	embed.Color = 16730698
	embed.Description = `Сервер недоступен ¯\_(ツ)_/¯`
	ctx.RespondEmbed(embed)
	return
}

func notEnoughPlayers(ctx *dgc.Ctx) {
	embed.Description = `Никого нету ¯\_(ツ)_/¯`
	embed.Color = 16730698
	ctx.RespondEmbed(embed)
	return
}

func notEnoughArgs(ctx *dgc.Ctx) {
	embed.Description = "Введи аргументы!"
	ctx.RespondEmbed(embed)
	return
}
