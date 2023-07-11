package message

import (
	"fmt"
	"whisperingdice/internal/pkg/log"

	"github.com/bwmarrin/discordgo"
)

// making GO provide a *float64 is strange .. .and hard.  this constant is a shim to acomplish that.
var integerOptionValueTwo = 2.0
var integerOptionValueOne = 1.0
var integerOptionNegative99 = -99.0

const (
	OptionPool  string = "pool"
	OptionSkill string = "skill"
	OptionCap   string = "cap"

	ResultSelectedSide string = "selectedside"
	ResultRolls        string = "rolls"
	ResultTotal        string = "total"
)

func CreateCommands(logger log.Logger) map[*discordgo.ApplicationCommand]func(s *discordgo.Session, i *discordgo.InteractionCreate) {

	makeFunction := func(fn func(optionMap map[string]*discordgo.ApplicationCommandInteractionDataOption) (message string, addl log.AddlInfo)) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		return func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(i.ApplicationCommandData().Options))
			for _, opt := range i.ApplicationCommandData().Options {
				optionMap[opt.Name] = opt
			}

			message, addl := fn(optionMap)
			addl["user"] = i.User.Username

			content := fmt.Sprintf("<@%s> %s", i.User.ID, message)

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				// Ignore type for now, they will be discussed in "responses"
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content,
				},
			})

			logger.Log(content, addl)
		}
	}

	return map[*discordgo.ApplicationCommand]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		NewChallengeCommand(): makeFunction(NewChallengeCommandHandler),
		NewDamageCommand():    makeFunction(NewDamageCommandHandler),
	}
}
