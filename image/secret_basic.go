package image

import (
	"github.com/bwmarrin/discordgo"
)

func SecretBasic(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionMessageComponent &&
		i.MessageComponentData().CustomID == CustomID_Next {

		embed := &discordgo.MessageEmbed{
			Description: "埋め込みメッセージ",
			Image: &discordgo.MessageEmbedImage{
				URL: ImageURL,
			},
		}

		resp := &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				//Content: "hello",
				Embeds: []*discordgo.MessageEmbed{embed},
				Flags:  discordgo.MessageFlagsEphemeral,
			},
		}

		if err := s.InteractionRespond(i.Interaction, resp); err != nil {
			panic(err)
		}
	}
}
