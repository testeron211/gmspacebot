package commands

import (
	"strconv"

	"gmspacebot/lib/config"

	"github.com/Lukaesebrot/dgc"
	"github.com/rumblefrog/go-a2s"
)

func Status(ctx *dgc.Ctx) {
	embed.Color = 4915018
	embed.Title = "Статус"

	client, err := a2s.NewClient(config.Option.IP)
	if err != nil {
		connectionTimeOut(ctx, err)
	}

	defer client.Close()

	info, err := client.QueryInfo() // QueryInfo, QueryPlayer, QueryRules

	if err != nil {
		connectionTimeOut(ctx, err)
	}

	Players := strconv.Itoa(int(info.Players))
	MaxPlayers := strconv.Itoa(int(info.MaxPlayers))
	Map := info.Map
	Game := info.Game

	embed.Description = "**" + info.Name + "**\n IP: `" + config.Option.IP + "`\n Карта: `" + Map + "`\n Игроков: `" + Players + "/" + MaxPlayers + "`\n Режим: `" + Game + "`"

	ctx.RespondEmbed(embed)
}
