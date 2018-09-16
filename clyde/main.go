package main

import (
  "flag"
  "fmt"
  "io"
  "log"
  "os"
  "os/signal"
  "regexp"
  "strings"
  "syscall"
  "time"

  "github.com/bwmarrin/discordgo"
  "github.com/boltdb/bolt"
  "github.com/mkdillard/clyde/dicebag"
  "github.com/mkdillard/clyde/scroll"
)

// Variables used for command line parameters
var (
  Token string
)

var db *bolt.DB

func init() {
  flag.StringVar(&Token, "t", "", "Bot Token.")
  flag.Parse()
}

func main() {

  //open boltdb database
  var err error //This is a hack to force global variable
  db, err = bolt.Open("clyde.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

  //Make sure db buckets exist
  err = scroll.CreateBuckets(db)
  if err != nil {
    log.Fatal("error creating db buckets: ", err)
  }

  // Create new discord
  dg, err := discordgo.New("Bot " + Token)
  if err != nil {
    log.Fatal("error creating Discord session: ", err)
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

  genericRollFinderRegex := regexp.MustCompile(".*[0-9]+d[0-9]+.*")
  genericRollParamRegex := regexp.MustCompile("[0-9]+d[0-9]+")

  // Find if the message starts with "/r"
  if strings.HasPrefix(m.Content, "/r") {
    testResult := ""
    if strings.Contains(m.Content, "1d6") || strings.Contains(m.Content, "rd") || strings.Contains(m.Content, "disadvantage") || strings.Contains(m.Content, "td") {
      testResult = dicebag.RollTest("disadvantage")
    } else if strings.Contains(m.Content, "2d6") || strings.Contains(m.Content, "rt") || strings.Contains(m.Content, "test") || strings.HasSuffix(m.Content, "t") {
      testResult = dicebag.RollTest("test")
    } else if strings.Contains(m.Content, "3d6") || strings.Contains(m.Content, "ra") || strings.Contains(m.Content, "advantage") || strings.Contains(m.Content, "ta") {
      testResult = dicebag.RollTest("advantage")
    } else if strings.Contains(m.Content, "hw") || strings.Contains(m.Content, "hurthaan") {
      testResult = dicebag.HurthaansWill()
    } else if genericRollFinderRegex.MatchString(m.Content) {
      words := strings.Split(m.Content, " ")
      for _, word := range words{
        if genericRollParamRegex.MatchString(word){
          testResult = dicebag.GenericRoll(word)
          break
        }
      }
    }
    if testResult != ""{
      s.ChannelMessageSend(m.ChannelID, testResult)
    }
  } else if strings.HasPrefix(m.Content, "/f") {
    focusResult := ""
    if strings.Contains(m.Content, "focus test") || strings.Contains(m.Content, "ft") {
      focusResult = dicebag.RollTest("focusTest")
    } else if strings.Contains(m.Content, "focus disadvantage") || strings.Contains(m.Content, "fd") {
      focusResult = dicebag.RollTest("focusDisadvantage")
    } else if strings.Contains(m.Content, "focus advantage") || strings.Contains(m.Content, "fa") {
      focusResult = dicebag.RollTest("focusAdvantage")
    } else if strings.Contains(m.Content, "focus marksman") || strings.Contains(m.Content, "fmt") {
      focusResult = dicebag.RollTest("focusMarksman")
    } else if strings.Contains(m.Content, "focus marksman disadvantage") || strings.Contains(m.Content, "fmd") {
      focusResult = dicebag.RollTest("focusMarksmanDisadvantage")
    } else if strings.Contains(m.Content, "focus marksman advantage") || strings.Contains(m.Content, "fma") {
      focusResult = dicebag.RollTest("focusMarksmanAdvantage")
    }

    if focusResult != ""{
      s.ChannelMessageSend(m.ChannelID, focusResult)
    }
  } else if strings.HasPrefix(m.Content, "/scribe") {
    response := ""
    if strings.Contains(m.Content, "record") {

      inputs := strings.SplitN(m.Content, " ", 4)
      k := inputs[2]
      v := inputs[3]
      err := scroll.Write(db, k, v)
      if err != nil {
        response = fmt.Sprintf("I'm sorry I was unable to record your information about %s", k)
        log.Printf("Error writing key: %s and value: %s to database: %v", k, v, err)
      } else {
        response = fmt.Sprintf("%s successfully added to my scrolls", k)
      }
    } else if strings.Contains(m.Content, "read") {
      inputs := strings.Split(m.Content, " ")
      k := inputs[2]
      v := ""
      err := scroll.Read(db, k, &v)
      if err != nil {
        response = "I am sorry I seem to have misplaced my scrolls I couldn't find what you wanted"
        log.Printf("Error reading from database: %v", err)
      }
      if v == "" {
        response = fmt.Sprintf("Sorry I was unable to find %s in my records", k)
      } else {
        response = fmt.Sprintf("I found information about %s:\n%s", k, v)
      }
    }

    if response != ""{
      s.ChannelMessageSend(m.ChannelID, response)
    }
  } else if strings.Compare(m.Content, "/backupScrolls") == 0 {
    log.Printf("Creating database backup")
    reader, writer := io.Pipe()
    var err error
    go func() {
      defer writer.Close()
      err = scroll.BackupScroll(db, writer)
      if err != nil {
        s.ChannelMessageSend(m.ChannelID, "Error creating a backup copy of my scrolls")
        log.Printf("Error creating backup of scrolls: %v", err)
      }
    }()

    log.Printf("after scroll thread started")
    currentTime := time.Now()
    filename := currentTime.Format("2006_01_02") + "_clyde.db"
    s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{File: &discordgo.File{Name: filename, Reader: reader}, Content:"Here is a backup copy of my scrolls"})
    log.Printf("Finished database backup")
  }
}
