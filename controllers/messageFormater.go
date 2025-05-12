package messageFormater

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var HookMap = map[string]func(s *discordgo.Session, m *discordgo.Member) string{
	"{user}": func(s *discordgo.Session, m *discordgo.Member) string {
		return m.User.Username
	},
	"{membercount}": func(s *discordgo.Session, m *discordgo.Member) string {
		guild, _ := s.State.Guild(m.GuildID)
		return fmt.Sprintf("%d", guild.MemberCount)
	},
	"{servername}": func(s *discordgo.Session, m *discordgo.Member) string {
		guild, _ := s.State.Guild(m.GuildID)
		return guild.Name
	},
}

func ProcessMessage(s *discordgo.Session, m *discordgo.Member, msg string) string {
	for hook, fn := range HookMap {
		msg = strings.ReplaceAll(msg, hook, fn(s, m))
	}
	return msg
}
