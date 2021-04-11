package main

import (
	"log"

	"gmspacebot/lib/bot"
	"gmspacebot/lib/config"
)

func main() {
	log.Println("Init bot config...")
	config.Init()

	log.Println("Init Bot...")
	bot.Init()
}
