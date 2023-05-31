package image

import "github.com/bwmarrin/discordgo"

// コマンドが実行されたら画像を送信します
func Basic(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content != CMD_Basic {
		return
	}

	// メッセージ作成
	msg := &discordgo.MessageSend{
		Content: "画像を送信します:",
		Embed: &discordgo.MessageEmbed{
			Image: &discordgo.MessageEmbedImage{
				URL: ImageURL,
			},
		},
	}

	// 画像を送信
	_, err := s.ChannelMessageSendComplex(m.ChannelID, msg)
	if err != nil {
		panic(err)
	}
}
