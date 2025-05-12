package handlers

import (
	messageFormater "neon-nexus/discord/controllers"
	"os"

	"github.com/bwmarrin/discordgo"
)

func UserJoin(s *discordgo.Session, member *discordgo.GuildMemberAdd) {
	welcome_channel := os.Getenv("WELCOME_CHANNEL")
	welcome_message := messageFormater.ProcessMessage(s, member.Member, "Bienvenue {user}")
	s.ChannelMessageSend(welcome_channel, welcome_message)
}

func UserLeave(s *discordgo.Session, member *discordgo.GuildMemberRemove) {
	leave_channel := os.Getenv("WELCOME_CHANNEL")
	leave_message := messageFormater.ProcessMessage(s, member.Member, "Ciao {user}")
	s.ChannelMessageSend(leave_channel, leave_message)
}
