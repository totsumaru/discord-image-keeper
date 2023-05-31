package image

import (
	"github.com/bwmarrin/discordgo"
)

// パネルを作成します
func PanelCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Content != CMDNamePanelCreate {
		return
	}

	btn1 := discordgo.Button{
		Label:    "Next",
		Style:    discordgo.PrimaryButton,
		CustomID: CustomID_Next,
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
