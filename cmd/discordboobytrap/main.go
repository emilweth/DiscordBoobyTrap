package main

import (
	"DiscordBoobyTrap/internal/config"
	"DiscordBoobyTrap/modules/discord"
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
)

var Version = "dev"

func init() {
	log.WithField("Version", Version).Info("Starting DiscordBoobyTrap")
	level, err := log.ParseLevel(config.GetVerbosity())
	if err != nil {
		log.Fatalf("Invalid verbosity level: %v", err)
	}

	// Set the Logrus verbosity level
	log.SetLevel(level)
}

func main() {
	versionFlag := flag.Bool("version", false, "Print the version and exit")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("DiscordBoobyTrap version: %s\n", Version)
		return
	}

	discord.Connect()
}
