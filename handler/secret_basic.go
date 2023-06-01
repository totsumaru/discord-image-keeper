package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/discord-image-keeper/handler/shared"
)

// ここは単体の画像を指定
const imageURL = ""

func SecretBasic(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type == discordgo.InteractionMessageComponent &&
		i.MessageComponentData().CustomID == shared.CustomID_Next {

		embed := &discordgo.MessageEmbed{
			Description: "埋め込みメッセージ",
			Image: &discordgo.MessageEmbedImage{
				URL: imageURL,
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
