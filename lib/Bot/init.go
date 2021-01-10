package Bot

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	commands "../../commands"
	config "../Config"

	"github.com/bwmarrin/discordgo"
	"github.com/lus/dgc"
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

	rate := dgc.NewRateLimiter(5*time.Second, 3*time.Second, func(ctx *dgc.Ctx) {
		ctx.RespondText("Нельзя использовать бота так часто!")
	})

	Router.RegisterDefaultHelpCommand(Session, rate)

	Router.Initialize(Session)

	commands.Init(Router)

	Session.AddHandler(playing)

	if config.Option.StatusCh == true {

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
