package image

import (
	"github.com/bwmarrin/discordgo"
	"github.com/techstart35/discord-image-keeper/image/shared"
)

// パネルを作成します
func PanelCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content != shared.CMDNamePanelCreate {
		return
	}

	btn1 := discordgo.Button{
		Label:    "Next",
		Style:    discordgo.PrimaryButton,
		CustomID: shared.CustomID_Next,
	}

	actions := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{btn1},
	}

	description := `
パネルです。
`

	embed := &discordgo.MessageEmbed{
		Description: description,
	}

	data := &discordgo.MessageSend{
		Components: []discordgo.MessageComponent{actions},
		Embed:      embed,
	}

	_, err := s.ChannelMessageSendComplex(m.ChannelID, data)
	if err != nil {
		panic(err)
	}
}
