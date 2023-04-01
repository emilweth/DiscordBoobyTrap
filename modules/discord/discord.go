package discord

import (
	"DiscordBoobyTrap/internal/config"
	"DiscordBoobyTrap/modules/discord/messageHandler"
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var (
	dg           *discordgo.Session
	discordToken string
	err          error
)

func init() {
	discordToken, err = config.GetDiscordToken()
	if err != nil {
		log.WithError(err).Panic()
	}
}

func Connect() {
	var err error
	dg, err = discordgo.New("Bot " + discordToken)
	if err != nil {
		log.WithError(err).Panic("Error creating discord session")
	}
	dg.AddHandler(messageHandler.Handle)
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		log.WithError(err).Panic("Error while opening connection")
	}

	// Wait here until CTRL-C or other term signal is received.
	log.Info(fmt.Sprintf("Connected to Discord as %s#%s", dg.State.User.Username, dg.State.User.Discriminator))
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	err = dg.Close()
	if err != nil {
		return
	}
}
