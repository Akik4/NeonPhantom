package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var leaveData = discordgo.ApplicationCommand{
	Name:        "leave",
	Description: "leave message",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionChannel,
			Name:        "channel",
			Description: "Define a channel to leave event",
			Required:    false,
		},
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "message",
			Description: "Define a message to leave event",
			Required:    false,
		},
	},
	DefaultMemberPermissions: &defaultPermissions,
}

func leave(s *discordgo.Session, i *discordgo.InteractionCreate) {

	options := i.ApplicationCommandData().Options

	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	margs := make([]interface{}, 0, len(options))
	msgformat := ""

	if opt, ok := optionMap["channel"]; ok {
		margs = append(margs, opt.ChannelValue(nil).ID)
		msgformat += "Channel defined on %s..."
		//@TODO
	}

	if opt, ok := optionMap["message"]; ok {
		margs = append(margs, opt.StringValue())
		msgformat += "Message defined by %s\n"
		//@TODO
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf(
				msgformat,
				margs...,
			),
		},
	})
}
