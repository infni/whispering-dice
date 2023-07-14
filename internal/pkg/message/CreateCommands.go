package message

import (
	"fmt"
	"whisperingdice/internal/pkg/exalted"
	"whisperingdice/internal/pkg/log"
	"whisperingdice/internal/pkg/whisperingvault"

	"github.com/bwmarrin/discordgo"
)

func CreateCommands(logger log.Logger) map[*discordgo.ApplicationCommand]func(s *discordgo.Session, i *discordgo.InteractionCreate) {

	makeHandlerFunction := func(fn func(optionMap map[string]*discordgo.ApplicationCommandInteractionDataOption) (message string, addl log.AddlInfo)) func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		return func(s *discordgo.Session, i *discordgo.InteractionCreate) {

			// re-orginize the options in a map for easier access
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(i.ApplicationCommandData().Options))
			for _, opt := range i.ApplicationCommandData().Options {
				optionMap[opt.Name] = opt
			}

			// run the provided funtion
			message, addl := fn(optionMap)

			content := fmt.Sprintf("<@%s> %s", i.Member.User.ID, message)

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: content,
				},
			})

			addl["user"] = i.Member.User.Username
			logger.Log(content, addl)
		}
	}

	return map[*discordgo.ApplicationCommand]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		whisperingvault.NewChallengeCommand(): makeHandlerFunction(whisperingvault.ChallengeCommandHandler),
		whisperingvault.NewDamageCommand():    makeHandlerFunction(whisperingvault.DamageCommandHandler),
		exalted.NewChallengeCommand():         makeHandlerFunction(exalted.ActionCommandHandler),
	}
}
