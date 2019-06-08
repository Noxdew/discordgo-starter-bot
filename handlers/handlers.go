package handlers

import (
	// App Modules
	"github.com/Noxdew/discordgo-starter-bot/bot"

	// External Modules
	"github.com/bwmarrin/discordgo"
)

// Handlers structure
type Handlers struct {
	Bot *bot.Bot
}

// Ready handler called when a `Ready` event is fired
func (h *Handlers) Ready(s *discordgo.Session, r *discordgo.Ready) {
	s.UpdateStatus(0, h.Bot.Config.Msg.Ready)
	h.Bot.Logger.Info("Bot has successfully logged in")
}
