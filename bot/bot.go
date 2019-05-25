package bot

import (
	// System Modules
	"os"
	"os/signal"
	"syscall"

	// External Modules
	configLoader "bitbucket.org/noxdew/config"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

// Bot structure
type Bot struct {
	Logger   *zap.Logger
	Config   *ServiceConfig
	handlers []interface{}
}

// Create a new bot instance
func Create() *Bot {
	customConfig := &ServiceConfig{}
	config := configLoader.Load(customConfig)

	logger := config.Logging.Create()

	b := &Bot{
		Config: customConfig,
		Logger: logger,
	}

	return b
}

// Close the bot instance
func (b *Bot) Close() {
	b.Logger.Sync()
}

// AddHandler to the bot
func (b *Bot) AddHandler(i interface{}) {
	b.handlers = append(b.handlers, i)
}

// Start the bot
func (b *Bot) Start() {
	var s *discordgo.Session

	// Login
	s, err := discordgo.New("Bot " + b.Config.Token)
	if err != nil {
		b.Logger.Panic(err.Error(), zap.Error(err))
	}

	// Add Handlers
	for _, handler := range b.handlers {
		s.AddHandler(handler)
	}

	// Open bot session
	err = s.Open()
	if err != nil {
		b.Logger.Panic(err.Error(), zap.Error(err))
	}

	// Wait here until CTRL-C or other term signal is received.
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}
