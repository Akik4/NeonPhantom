package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var banData = discordgo.ApplicationCommand{
	Name:        "ban",
	Description: "For banning an user",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionUser,
			Name:        "user",
			Description: "Define user to ban",
			Required:    true,
		},
	},
	DefaultMemberPermissions: &defaultPermissions,
}

func ban(s *discordgo.Session, i *discordgo.InteractionCreate) {

	options := i.ApplicationCommandData().Options

	optionsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionsMap[opt.Name] = opt
	}

	margs := make([]interface{}, 0, len(options))
	msgformat := ""
	user := new(discordgo.User)

	if opt, ok := optionsMap["user"]; ok {
		margs = append(margs, opt.UserValue(nil).ID)
		user = opt.UserValue(nil)
		msgformat = "<@%s> is banned"
	}

	s.GuildBanCreate(i.GuildID, user.ID, -1)

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
