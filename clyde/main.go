package main

import (
  "flag"
  "fmt"
  "os"
  "os/signal"
  "strings"
  "syscall"

  "github.com/bwmarrin/discordgo"
  "github.com/mkdillard/clyde/dicebag/dicebag"
)

// Variables used for command line parameters
var (
  Token string
)

func init() {
  flag.StringVar(&Token, "t", "", "Bot Token.")
  flag.Parse()
}

func main() {
  // Create new discord
  dg, err := discordgo.New("Bot " + Token)
  if err != nil {
    fmt.Println("error creating Discord session: ", err)
    return
  }

  //Register the messageCreate func as a callback for MessageCreate events.
  dg.AddHandler(messageCreate)

  //Open a websocket connection to Discord and begin listening.
  err = dg.Open()
  if err != nil {
    fmt.Println("error opening Discord session: ", err)
    return
  }

  // Wait here until CTRL-C or other term signal is received.
  fmt.Println("Bot is now running. Press CTRL-C to exit.")
  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
  <-sc

  //Cleanly close down the Discord session.
  dg.Close()
}

//This function will be called (due to the AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  // Ignore all messages created by the bot itself
  // This isn't required in this specific example but it's a good practice.
  if m.Author.ID == s.State.User.ID {
    return
  }
  // If the message is "ping" reply with "Pong!"
  if m.Content == "ping" {
    s.ChannelMessageSend(m.ChannelID, "Pong!")
  }

  // If the message is "pong" reply with "Ping!"
  if m.Content == "pong" {
    s.ChannelMessageSend(m.ChannelID, "Ping!")
  }

  // Find if the message starts with "/r"
  if strings.HasPrefix(m.Content, "/r") {
      testResult := ""
      if strings.Contains(m.Content, "1d6") || strings.Contains(m.Content, "rd") || strings.Contains(m.Content, "disadvantage") || strings.Contains(m.Content, "td") {
        testResult = dicebag.rollTest("disadvantage")
      } else if strings.Contains(m.Content, "2d6") || strings.Contains(m.Content, "rt") || strings.Contains(m.Content, "test") || strings.HasSuffix(m.Content, "t") {
        testResult = dicebag.rollTest("test")
      } else if strings.Contains(m.Content, "3d6") || strings.Contains(m.Content, "ra") || strings.Contains(m.Content, "advantage") || strings.Contains(m.Content, "ta") {
        testResult = dicebag.rollTest("advantage")
      }
      s.ChannelMessageSend(m.ChannelID, testResult)
  }
}
