package commands

import (
	"gmspacebot/lib/config"
	"log"
	"strconv"

	"github.com/Lukaesebrot/dgc"
	"github.com/rumblefrog/go-a2s"
)

func Players(ctx *dgc.Ctx) {
	embed.Color = 4915018
	embed.Title = "Игроки на сервере"

	client, err := a2s.NewClient(config.Option.IP)
	if err != nil {
		log.Println(err)
		embed.Color = 16730698
		embed.Description = "Сервер недоступен"
		ctx.RespondEmbed(embed)
		return
	}

	defer client.Close()

	info, err := client.QueryPlayer()
	if err != nil {
		log.Println(err)
		embed.Color = 16730698
		embed.Description = "Сервер недоступен"
		ctx.RespondEmbed(embed)
		return
	}

	var result string

	for i, player := range info.Players {
		if player.Name == "" {
			player.Name = "Подключается..."
			continue
		}

		plytimeraw := int(player.Duration)
		s := strconv.Itoa(plytimeraw % 60)
		m := strconv.Itoa(plytimeraw / 60 % 60)
		h := strconv.Itoa(plytimeraw / 3600 % 24)
		result = result + strconv.Itoa(i+1) + ". [" + h + ":" + m + ":" + s + "] " + player.Name + "\n"
	}

	embed.Description = "```\n" + result + "\n```"
	ctx.RespondEmbed(embed)
}
