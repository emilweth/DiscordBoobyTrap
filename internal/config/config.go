package config

import (
	"errors"

	"github.com/gookit/config/v2"
)

func init() {

	config.LoadOSEnvs(map[string]string{
		"DISCORD_TOKEN":      "discord_token",
		"DISCORD_CHANNEL_ID": "discord_channel_id",
		"VERBOSITY":          "verbosity",
	})
}

func GetDiscordToken() (string, error) {
	token := config.String("discord_token")
	if len(token) == 0 {
		return "", errors.New("missing discord token")
	}
	return token, nil
}

func GetDiscordChannelId() (string, error) {
	channelId := config.String("discord_channel_id")
	if len(channelId) == 0 {
		return "", errors.New("missing discord channel ID")
	}
	return channelId, nil
}

func GetVerbosity() string {
	verbosity := config.String("verbosity")

	if len(verbosity) == 0 {
		return "info"
	} else {
		return verbosity
	}

}
