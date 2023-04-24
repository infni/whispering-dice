package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"whisperingdice/internal/pkg/log"
	"whisperingdice/internal/pkg/message"

	"github.com/bwmarrin/discordgo"
)

var version = "defined at build time"

const (
	defaultDiscordToken string = "collectatruntime"
)

func main() {

	// get configuration
	cfg := &Config{}
	flag.BoolVar(&cfg.DisplayVesion, "version", false, "display the version and exit")
	flag.StringVar(&cfg.DiscordToken, "token", defaultDiscordToken, "the discord tolken for this API account")

	flag.Parse()

	if discordToken := os.Getenv("TOKEN"); len(discordToken) > 0 && cfg.DiscordToken == defaultDiscordToken {
		cfg.DiscordToken = discordToken
	}

	if cfg.DisplayVesion {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}

	if execute(cfg) {
		fmt.Print("\nExited\n")
		os.Exit(0)
	}

	// the error has previously been written out.  Just exit.
	os.Exit(1)
}

type customLogger struct{}

func (_ customLogger) Print(v ...interface{}) {
	fmt.Println(v...)
}

func execute(cfg *Config) bool {

	logger := log.NewFileLogger(customLogger{})

	dg, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return false
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(message.CreateResponseFunc(logger))

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return false
	}
	// Cleanly close down the Discord session.
	defer dg.Close()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	return true
}
