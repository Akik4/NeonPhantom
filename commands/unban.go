package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var unbanData = discordgo.ApplicationCommand{
	Name:        "unban",
	Description: "For unbanning an user",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "userid",
			Description: "Define user id to unban",
			Required:    true,
		},
	},
	DefaultMemberPermissions: &defaultPermissions,
}

func unban(s *discordgo.Session, i *discordgo.InteractionCreate) {

	options := i.ApplicationCommandData().Options

	optionsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionsMap[opt.Name] = opt
	}

	margs := make([]interface{}, 0, len(options))
	msgformat := ""
	userID := ""

	if opt, ok := optionsMap["userid"]; ok {
		margs = append(margs, opt.StringValue())
		userID = opt.StringValue()
		msgformat = "<@%s> is unbanned"
	}

	s.GuildBanDelete(i.GuildID, userID)

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
