package events

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"duckess-bot/constants"
	"duckess-bot/events/commands"
	"duckess-bot/types"

	"github.com/bwmarrin/discordgo"
)



func sendFile(discordSession *discordgo.Session, channelId string, fileName string, data io.Reader){
    _, err := discordSession.ChannelFileSend(channelId, fileName, data)
    if err != nil {
        fmt.Println(err)
    }
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func MessageCreate(discordSession *discordgo.Session, message *discordgo.MessageCreate) {

    // Ignore all messages created by the bot itself
    // This isn't required in this specific example but it's a good practice.
    if message.Author.ID == discordSession.State.User.ID {
        return
    }

    if message.Content == "!gopher" {

        //Call the KuteGo API and retrieve our cute Dr Who types.Gopher
        response, err := http.Get(constants.KuteGoAPIURL + "/gopher/" + "dr-who")
        if err != nil {
            fmt.Println(err)
        }
        defer response.Body.Close()

        if response.StatusCode == 200 {
            _, err = discordSession.ChannelFileSend(message.ChannelID, "dr-who.png", response.Body)
            if err != nil {
                fmt.Println(err)
            }
        } else {
            fmt.Println("Error: Can't get dr-who types.Gopher! :-(")
        }
    }

    if message.Content == "!random" {

        result := commands.Random()
        fmt.Println(result)
        //Call the KuteGo API and retrieve a random types.Gopher
        response, err := http.Get(constants.KuteGoAPIURL + "/gopher/random/")
        if err != nil {
            fmt.Println(err)
        }
        defer response.Body.Close()

        if response.StatusCode == 200 {
            sendFile(discordSession,message.ChannelID,"random-gopher.png",response.Body)
        } else {
            fmt.Println("Error: Can't get random types.Gopher! :-(")
        }
    }

    if message.Content == "!gophers" {

        //Call the KuteGo API and display the list of available Gophers
        response, err := http.Get(constants.KuteGoAPIURL + "/gophers/")
        if err != nil {
            fmt.Println(err)
        }
        defer response.Body.Close()

        if response.StatusCode == 200 {
            // Transform our response to a []byte
            body, err := ioutil.ReadAll(response.Body)
            if err != nil {
                fmt.Println(err)
            }

            // Put only needed informations of the JSON document in our array of types.Gopher
            var data []types.Gopher
            err = json.Unmarshal(body, &data)
            if err != nil {
                fmt.Println(err)
            }

            // Create a string with all of the types.Gopher's name and a blank line as separator
            var gophers strings.Builder
            for _, gopher := range data {
                gophers.WriteString(gopher.Name + "\n")
            }

            // Send a text message with the list of Gophers
            _, err = discordSession.ChannelMessageSend(message.ChannelID, gophers.String())
            if err != nil {
                fmt.Println(err)
            }
        } else {
            fmt.Println("Error: Can't get list of Gophers! :-(")
        }
    }
}