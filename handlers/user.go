package handlers

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

func UserJoin(s *discordgo.Session, member *discordgo.GuildMemberAdd) {
	welcome_channel := os.Getenv("WELCOME_CHANNEL")
	welcome_message := "Bienvenue " + member.User.Username
	s.ChannelMessageSend(welcome_channel, welcome_message)
}

func UserLeave(s *discordgo.Session, member *discordgo.GuildMemberRemove) {
	leave_channel := os.Getenv("WELCOME_CHANNEL")
	leave_message := "Ciao " + member.User.Username
	s.ChannelMessageSend(leave_channel, leave_message)
}
