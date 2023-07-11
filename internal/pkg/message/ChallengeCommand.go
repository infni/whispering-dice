package message

import (
	"fmt"
	"strings"
	"whisperingdice/internal/pkg/log"
	"whisperingdice/internal/pkg/rolling"

	"github.com/bwmarrin/discordgo"
)

func NewChallengeCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "challenge",
		Description: "make a challenge roll (attack, defend, evoke, skill, etc.)",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        OptionPool,
				Description: "The number of dice in the pool for this challenge roll (max=25)",
				Required:    true,
				MinValue:    &integerOptionValueOne,
				MaxValue:    25,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        OptionSkill,
				Description: "The skill bonus added to the final roll (range -99,99). defaults to 0.",
				Required:    false,
				MinValue:    &integerOptionNegative99,
				MaxValue:    99,
			},
		},
	}
}

func NewChallengeCommandHandler(optionMap map[string]*discordgo.ApplicationCommandInteractionDataOption) (message string, addl log.AddlInfo) {

	pool := int(optionMap[OptionPool].IntValue())
	skill := 0
	if val, ok := optionMap[OptionSkill]; ok {
		skill = int(val.IntValue())
	}

	// roll dice
	rolls, total, selectedSide := rolling.RollChallengeAllRolls(pool)

	// format rolls
	formattedRolls := make([]string, len(rolls))
	for i := 0; i < len(rolls); i++ {
		if rolls[i] == selectedSide {
			formattedRolls[i] = fmt.Sprintf("**%d**", rolls[i])
		} else {
			formattedRolls[i] = fmt.Sprintf("%d", rolls[i])
		}
	}

	addl = log.AddlInfo{
		OptionPool:         pool,
		OptionSkill:        skill,
		ResultSelectedSide: selectedSide,
		ResultTotal:        total,
		ResultRolls:        rolls,
	}

	return fmt.Sprintf("Roll: [%s] Total: _(%d%+d)_= **%d**", strings.Join(formattedRolls, ","), total, skill, (total + skill)), addl
}
