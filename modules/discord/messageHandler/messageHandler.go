package messageHandler

import (
	"DiscordBoobyTrap/internal/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var (
	channelId string
	err       error
)

func init() {
	channelId, err = config.GetDiscordChannelId()
	if err != nil {
		log.WithError(err).Panic()
	}
}

func Handle(s *discordgo.Session, m *discordgo.MessageCreate) {
	fullUsername := fmt.Sprintf("%s#%s", m.Author.Username, m.Author.Discriminator)

	if m.ChannelID == channelId {
		log.WithFields(log.Fields{
			"user":    fullUsername,
			"userID":  m.Author.ID,
			"message": m.Message.Content,
		}).Info("Boobytrap triggered")
		log.Info(fmt.Sprintf("Trying to ban %s", fullUsername))
		err := s.GuildBanCreate(m.GuildID, m.Author.ID, 7)
		if err != nil {
			log.WithError(err).Error(fmt.Sprintf("Could not ban %s", fullUsername))
			return
		}
	}
}