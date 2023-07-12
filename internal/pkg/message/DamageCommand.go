package message

import (
	"fmt"
	"strings"
	"whisperingdice/internal/pkg/log"
	"whisperingdice/internal/pkg/rolling"

	"github.com/bwmarrin/discordgo"
)

func NewDamageCommand() *discordgo.ApplicationCommand {
	return &discordgo.ApplicationCommand{
		Name:        "damage",
		Description: "make a challenge roll (attack, defend, evoke, skill, etc.)",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        OptionPool,
				Description: "The number of dice in the pool for this damage roll (max=25)",
				Required:    true,
				MinValue:    &integerOptionValueOne,
				MaxValue:    25,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        OptionCap,
				Description: "The maximum number that will be counted when rolling damage (2 <= value <= 6). defaults to 6.",
				Required:    false,
				MinValue:    &integerOptionValueTwo,
				MaxValue:    6,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        OptionComment,
				Description: "A comment.  (<99 characters)",
				Required:    false,
				MaxLength:   99,
			},
		},
	}
}

func NewDamageCommandHandler(optionMap map[string]*discordgo.ApplicationCommandInteractionDataOption) (message string, addl log.AddlInfo) {

	comment := optionMap[OptionComment].StringValue()
	if len(comment) > 0 {
		comment = " ## " + comment
	}
	pool := int(optionMap[OptionPool].IntValue())
	cap := 6
	if val, ok := optionMap[OptionCap]; ok {
		cap = int(val.IntValue())
	}

	// roll dice
	rolls, total := rolling.RollDamageAllRolls(pool, cap)

	// format rolls
	formattedRolls := make([]string, len(rolls))
	for i, val := range rolls {
		if val > cap {
			formattedRolls[i] = fmt.Sprintf("~~_%d_~~", val)
		} else {
			formattedRolls[i] = fmt.Sprintf("%d", val)
		}
	}

	addl = log.AddlInfo{
		OptionPool:  pool,
		OptionCap:   cap,
		ResultTotal: total,
		ResultRolls: rolls,
	}

	return fmt.Sprintf("Damage: **%d** Cap: %d [%s]%s", total, cap, strings.Join(formattedRolls, ","), comment), addl
}
