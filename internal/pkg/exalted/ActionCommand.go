package exalted

import (
	"fmt"
	"sort"
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
				Description: "The number of dice in the pool for this action (max=50)",
				Required:    true,
				MinValue:    &IntegerOptionValueOne,
				MaxValue:    50,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        OptionAutomaticSuccesses,
				Description: "The qauntity of automatic success to add to this roll (<25)",
				Required:    false,
				MinValue:    &IntegerOptionValueOne,
				MaxValue:    25,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        OptionDoubleSuccesses,
				Description: "The number (and higher) that will double successes (default=10, 0=double nothing)",
				Required:    false,
				MinValue:    &IntegerOptionValueZero,
				MaxValue:    10,
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        OptionRerollOnes,
				Description: "Reroll all 1 until they fail to appear",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        OptionRerollTwos,
				Description: "Reroll all 2 until they fail to appear",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        OptionRerollThrees,
				Description: "Reroll all 3 until they fail to appear",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        OptionRerollFours,
				Description: "Reroll all 4 until they fail to appear",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        OptionRerollFives,
				Description: "Reroll all 5 until they fail to appear",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionBoolean,
				Name:        OptionRerollSixes,
				Description: "Reroll all 6 until they fail to appear",
				Required:    false,
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

func ActionCommandHandler(optionMap map[string]*discordgo.ApplicationCommandInteractionDataOption) (message string, addl log.AddlInfo) {

	comment := ""
	if s, ok := optionMap[OptionComment]; ok {
		comment = " ## " + s.StringValue()
	}
	pool := int(optionMap[OptionPool].IntValue())
	as := 0
	if v, ok := optionMap[OptionAutomaticSuccesses]; ok {
		as = int(v.IntValue())
	}
	db := 10
	if v, ok := optionMap[OptionDoubleSuccesses]; ok {
		db = int(v.IntValue())
	}
	re := make(map[int]struct{}, 0)
	if _, ok := optionMap[OptionRerollOnes]; ok {
		re[1] = struct{}{}
	}
	if _, ok := optionMap[OptionRerollTwos]; ok {
		re[2] = struct{}{}
	}
	if _, ok := optionMap[OptionRerollThrees]; ok {
		re[3] = struct{}{}
	}
	if _, ok := optionMap[OptionRerollFours]; ok {
		re[4] = struct{}{}
	}
	if _, ok := optionMap[OptionRerollFives]; ok {
		re[5] = struct{}{}
	}
	if _, ok := optionMap[OptionRerollSixes]; ok {
		re[6] = struct{}{}
	}

	// roll dice
	rolls, total := rolling.RollActionAllRolls(pool, db, as, re)

	sort.Ints(rolls)

	// format rolls
	formattedRolls := make([]string, len(rolls))
	for i, val := range rolls {
		if _, ok := re[val]; ok {
			formattedRolls[i] = fmt.Sprintf("~~_%d_~~", val)
		} else if val >= db {
			formattedRolls[i] = fmt.Sprintf("**%d**", val)
		} else if val >= 7 {
			formattedRolls[i] = fmt.Sprintf("%d", val)
		} else {
			formattedRolls[i] = fmt.Sprintf("_%d_", val)
		}
	}

	botchText := ""
	_, re1 := re[1]
	if total == 0 && !re1 && rolls[0] == 1 {
		botchText = " **_BOTCH!!_** "
	}

	rearray := make([]string, len(re))
	i := 0
	for k := range re {
		rearray[i] = fmt.Sprintf("%d", k)
		i++
	}
	sort.Strings(rearray)
	addl = log.AddlInfo{
		OptionPool:            pool,
		OptionDoubleSuccesses: db,
		"re":                  strings.Join(rearray, ","),
		ResultTotal:           total,
		ResultRolls:           rolls,
	}

	return fmt.Sprintf("Pool: %d Rolls: [%s] Successes: **%d**%s%s", pool, strings.Join(formattedRolls, ","), total, botchText, comment), addl
}
