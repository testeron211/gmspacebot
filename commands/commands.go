package commands

import (
	"log"

	"github.com/Lukaesebrot/dgc"
)

func Init(router *dgc.Router) {
	log.Printf("Initialize commands...")

	defer log.Printf("Commands initialized!")

	router.RegisterCmd(&dgc.Command{
		Name:        "ping",
		Description: "Вернёт pong",
		Usage:       "ping",
		Example:     "ping",
		IgnoreCase:  true,
		Handler:     Ping,
	})

	router.RegisterCmd(&dgc.Command{
		Name:        "echo",
		Description: "Вернёт аргументы",
		Usage:       "echo [a]...",
		Example:     "echo Hello World!",
		IgnoreCase:  true,
		Handler:     Echo,
	})

	router.RegisterCmd(&dgc.Command{
		Name:        "status",
		Description: "Вернёт статус Garry's Mod сервера",
		Usage:       "status",
		Example:     "status",
		IgnoreCase:  true,
		Handler:     Status,
	})

	router.RegisterCmd(&dgc.Command{
		Name:        "players",
		Description: "Вернёт игроков на сервере.",
		Usage:       "players",
		Example:     "players",
		IgnoreCase:  true,
		Handler:     Players,
	})
}
