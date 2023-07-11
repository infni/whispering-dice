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
	defaultAppId        string = "defaultAppId"
	defaultGuidId       string = ""
)

func main() {

	// get configuration
	cfg := &Config{}
	flag.BoolVar(&cfg.DisplayVesion, "version", false, "display the version and exit")
	flag.StringVar(&cfg.DiscordToken, "token", defaultDiscordToken, "the discord tolken for this API account")
	flag.StringVar(&cfg.AppId, "appid", defaultAppId, "the application ID for this API account")
	flag.StringVar(&cfg.GuildId, "guildid", defaultGuidId, "the application ID for this API account")

	flag.Parse()

	if discordToken := os.Getenv("TOKEN"); len(discordToken) > 0 && cfg.DiscordToken == defaultDiscordToken {
		cfg.DiscordToken = discordToken
	}
	if appId := os.Getenv("APPID"); len(appId) > 0 && cfg.AppId == defaultAppId {
		cfg.AppId = appId
	}
	if guildId := os.Getenv("GUILDID"); len(guildId) > 0 && cfg.GuildId == defaultGuidId {
		cfg.GuildId = guildId
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

func execute(cfg *Config) bool {

	logger := log.NewFileLogger(customLogger{})

	dg, err := discordgo.New("Bot " + cfg.DiscordToken)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return false
	}

	commands := message.CreateCommands(logger)

	commandArray := make([]*discordgo.ApplicationCommand, len(commands))
	commandHandlers := make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate), len(commands))
	i := 0
	for k, v := range commands {
		commandArray[i] = k
		commandHandlers[k.Name] = v
		i++
	}

	_, err = dg.ApplicationCommandBulkOverwrite(cfg.AppId, cfg.GuildId, commandArray)
	if err != nil {
		fmt.Println("error bulk registering command,", err)
		return false
	}

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

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
