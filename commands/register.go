package commands

import (
	"flag"
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var removeCommands = flag.Bool("rmcmd", true, "remove all commands after shuting down")
var registeredCommands = make([]*discordgo.ApplicationCommand, len(commands))
var defaultPermissions int64 = discordgo.PermissionAdministrator

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
		DefaultMemberPermissions: &defaultPermissions,
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
		DefaultMemberPermissions: &defaultPermissions,
	},
	{
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
	},
	{
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
	},
	{
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
	"kick": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

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
	},
	"ban": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

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
	},
	"unban": func(s *discordgo.Session, i *discordgo.InteractionCreate) {

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
	},
}

func RegisterCommands(discord *discordgo.Session) {
	discord.AddHandler(CommandsRegister)

	fmt.Println("Adding commands...")
	for i, v := range commands {
		cmd, err := discord.ApplicationCommandCreate(discord.State.User.ID, *guildID, v)
		if err != nil {
			fmt.Println("A command cannot be created")
			fmt.Println(err.Error())
		}
		fmt.Println("Adding a command")
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
