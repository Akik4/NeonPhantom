package commands

import (
	"flag"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var removeCommands = flag.Bool("rmcmd", true, "remove all commands after shuting down")
var registeredCommands = make([]*discordgo.ApplicationCommand, len(commands))

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "welcome",
		Description: "welcome message",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionChannel,
				Name:        "channel",
				Description: "Define a channel to welcome event",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "message",
				Description: "Define a message to welcome event",
				Required:    false,
			},
		},
	},
	{
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
	},
}

var guildID = flag.String("guild", "", "Test uild ID")

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"welcome": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

		options := i.ApplicationCommandData().Options

		optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionMap[opt.Name] = opt
		}

		margs := make([]interface{}, 0, len(options))
		msgformat := ""

		if opt, ok := optionMap["channel"]; ok {
			margs = append(margs, opt.ChannelValue(nil).ID)
			msgformat += "Channel defined on %s\n"
			//@TODO
		}

		if opt, ok := optionMap["message"]; ok {
			margs = append(margs, opt.StringValue())
			msgformat += "Message defined by %s"
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
	},
	"leave": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

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
	},
}

func RegisterCommands(discord *discordgo.Session) {
	discord.AddHandler(CommandsRegister)

	fmt.Println("Adding commands...")
	for i, v := range commands {
		cmd, err := discord.ApplicationCommandCreate(discord.State.User.ID, *guildID, v)
		if err != nil {
			fmt.Println("A command cannot be created")
		}
		registeredCommands[i] = cmd
	}
}

func CommandsRegister(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}

func RemoveCommands(discord *discordgo.Session) {
	if *removeCommands {
		log.Println("Removing commands...")

		for _, v := range registeredCommands {
			err := discord.ApplicationCommandDelete(discord.State.User.ID, *guildID, v.ID)
			if err != nil {
				fmt.Println("Cannot delete a command")
			}
		}
	}
}
