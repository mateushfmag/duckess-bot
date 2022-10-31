package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"duckess-bot/events"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)


func main() {

    
    // load .env file
    err := godotenv.Load(".env")
  
    if err != nil {
        fmt.Println("error loading env variables,", err)
        return 
    }

    Token := os.Getenv("DISCORD_TOKEN")


    // Create a new Discord session using the provided bot token.
    discordGo, err := discordgo.New("Bot " + Token)
    if err != nil {
        fmt.Println("error creating Discord session,", err)
        return
    }

    // Register the messageCreate func as a callback for MessageCreate events.
    discordGo.AddHandler(events.MessageCreate)

    // In this example, we only care about receiving message events.
    discordGo.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentMessageContent

    // Open a websocket connection to Discord and begin listening.
    err = discordGo.Open()
    if err != nil {
        fmt.Println("error opening connection,", err)
        return
    }

    // Wait here until CTRL-C or other term signal is received.
    fmt.Println("Bot is now running. Press CTRL-C to exit.")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    // Cleanly close down the Discord session.
    discordGo.Close()
}