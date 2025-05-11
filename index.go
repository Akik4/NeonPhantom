package main

import (
	"fmt"
	"neon-nexus/discord/commands"
	"neon-nexus/discord/handlers"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	discord, err := discordgo.New("Bot " + os.Getenv("TOKEN"))
	if err != nil {
		fmt.Println("Error when creating discord bot")
		return
	}

	discord.Identify.Intents = discordgo.IntentGuildMembers

	discord.AddHandler(handlers.UserJoin)
	discord.AddHandler(handlers.UserLeave)

	err = discord.Open()
	if err != nil {
		fmt.Println("Error when openning session : " + err.Error())
		return
	}
	fmt.Println("Session connected and oppened.")

	commands.RegisterCommands(discord)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	commands.RemoveCommands()

	discord.Close()
}
