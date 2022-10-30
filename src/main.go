package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"duckess-bot/events"
    "github.com/joho/godotenv"
)

// Variables used for command line parameters
var (
    Token string
)

func init() {
    flag.StringVar(&Token, "t", "", "Bot Token")
    flag.Parse()
}

func main() {

    
    // load .env file
    err := godotenv.Load(".env")
  
    if err != nil {
        fmt.Println("error loading env variables,", err)
        return 
    }


    fmt.Println("LOAD ENV", os.Getenv("TESTE"))

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