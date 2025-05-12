package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var kickData = discordgo.ApplicationCommand{
	Name:        "kick",
	Description: "For kicking an user",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionUser,
			Name:        "user",
			Description: "Define the user to kick",
			Required:    true,
		},
	},
	DefaultMemberPermissions: &defaultPermissions,
}

func kick(s *discordgo.Session, i *discordgo.InteractionCreate) {

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
		msgformat = "<@%s> is kicked"
	}

	s.GuildMemberDelete(i.GuildID, user.ID)

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
