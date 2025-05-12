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
	&welcomeData,
	&leaveData,
	&kickData,
	&banData,
	&unbanData,
	&muteData,
	&unmuteData,
}

var guildID = flag.String("guild", "", "Test uild ID")

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"welcome": welcome,
	"leave":   leave,
	"kick":    kick,
	"ban":     ban,
	"unban":   unban,
	"mute":    mute,
	"unmute":  unmute,
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
