package bot

import (
	"gmspacebot/lib/config"

	"github.com/bwmarrin/discordgo"
)

func playing(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateGameStatus(0, config.Option.Status)
}
