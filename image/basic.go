package image

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/discord-image-keeper/image/shared"
)

// コマンドが実行されたら画像を送信します
func Basic(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content != shared.CMD_Basic {
		return
	}

	for _, url := range shared.ImageURLs {
		// メッセージ作成
		msg := &discordgo.MessageSend{
			Content: "画像を送信します:",
			Embed: &discordgo.MessageEmbed{
				Image: &discordgo.MessageEmbedImage{
					URL: url,
				},
			},
		}

		// 画像を送信
		_, err := s.ChannelMessageSendComplex(m.ChannelID, msg)
		if err != nil {
			panic(err)
		}
	}
}
