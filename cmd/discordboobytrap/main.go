package main

import (
	"DiscordBoobyTrap/internal/config"
	"DiscordBoobyTrap/modules/discord"
	log "github.com/sirupsen/logrus"
)

func init() {
	level, err := log.ParseLevel(config.GetVerbosity())
	if err != nil {
		log.Fatalf("Invalid verbosity level: %v", err)
	}

	// Set the Logrus verbosity level
	log.SetLevel(level)
}

func main() {
	discord.Connect()
}
