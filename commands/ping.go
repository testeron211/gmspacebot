package commands

import (
	"github.com/Lukaesebrot/dgc"
)

func Ping(ctx *dgc.Ctx) {
	ctx.RespondText(":ping_pong: **Pong!**")
}
