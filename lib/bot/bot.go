package bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"gmspacebot/commands"
	"gmspacebot/lib/config"

	"github.com/Lukaesebrot/dgc"
	"github.com/bwmarrin/discordgo"
)

var (
	Router  *dgc.Router
	Session *discordgo.Session
)

func Init() {
	Session, err := discordgo.New("Bot " + config.Option.Token)
	if err != nil {
		log.Fatal(err)
	}

	Router = dgc.Create(&dgc.Router{
		Prefixes: []string{config.Option.Prefix},
	})

	Router.RegisterDefaultHelpCommand(Session, nil)

	Router.Initialize(Session)

	commands.Init(Router)

	Session.AddHandler(playing)

	if config.Option.StatusCh == true {
		/*
			Постоянное сообщение со статусом
			Представьте, что оно есть :shrug:
		*/
	}

	err = Session.Open()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bot initialized!")

	log.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	Session.Close()
}
