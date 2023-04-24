package message

import (
	"fmt"
	"strings"
	"whisperingdice/internal/pkg/log"
	"whisperingdice/internal/pkg/rolling"

	"github.com/bwmarrin/discordgo"
)

func CreateResponseFunc(logger log.Logger) func(s *discordgo.Session, m *discordgo.MessageCreate) {

	damage := NewDamageRegex()

	challenge := NewChallengeRegex()

	return func(s *discordgo.Session, m *discordgo.MessageCreate) {

		send := func(res string) {
			_, err := s.ChannelMessageSend(m.ChannelID, res)
			if err != nil {
				logger.LogCritical(err.Error(), nil)
			}
		}

		// Ignore all messages created by the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		// if it's not a command for us, ignore it
		if !strings.HasPrefix(m.Content, CommandPrefix) {
			return
		}

		// if the command has no content, then just return the help message.
		if strings.Replace(m.Content, CommandPrefix, "", 1) == "" {
			send(createHelpResponse())
			return
		}

		// look for different command formats in sequence.  the first one that matches returns.  If none match, return the help message.

		// Damage response
		if match := damage.FindAllStringSubmatch(m.Content, -1); match != nil {

			// validate match inputs
			diceCount, dieCap, errStr := ValidateDamageMessageParams(match)
			if len(errStr) > 0 {
				send(fmt.Sprintf("<@%s> %s", m.Author.ID, errStr))
				return
			}

			// roll dice
			rolls, total := rolling.RollDamageAllRolls(diceCount, dieCap)

			// format rolls
			formattedRolls := make([]string, len(rolls))
			for i, val := range rolls {
				if val > dieCap {
					formattedRolls[i] = fmt.Sprintf("~~_%d_~~", val)
				} else {
					formattedRolls[i] = fmt.Sprintf("%d", val)
				}
			}

			// send formatted message
			send(fmt.Sprintf("<@%s> Damage: **%d** Cap: %d [%s]", m.Author.ID, total, dieCap, strings.Join(formattedRolls, ",")))

			// log
			logger.Log("responded with a damage roll message.", log.AddlInfo{
				"input":     m.Content,
				"dice":      diceCount,
				"capNumber": dieCap,
				"author":    m.Author.Username,
				"total":     total,
				"rolls":     rolls,
			})

			// do not check any other commands. one response per message
			return
		}

		// Challenge rsponse
		if match := challenge.FindAllStringSubmatch(m.Content, -1); match != nil {

			// validate match inputs
			diceCount, skill, errStr := ValidateChallengeMessageParams(match)
			if len(errStr) > 0 {
				send(fmt.Sprintf("<@%s> %s", m.Author.ID, errStr))
				return
			}

			// roll dice
			rolls, total, selectedSide := rolling.RollChallengeAllRolls(diceCount)

			// format rolls
			formattedRolls := make([]string, len(rolls))
			for i := 0; i < len(rolls); i++ {
				if rolls[i] == selectedSide {
					formattedRolls[i] = fmt.Sprintf("**%d**", rolls[i])
				} else {
					formattedRolls[i] = fmt.Sprintf("%d", rolls[i])
				}
			}

			// send formatted message
			send(fmt.Sprintf("<@%s> Roll: [%s] Total: _(%d%+d)_= **%d**", m.Author.ID, strings.Join(formattedRolls, ","), total, skill, (total + skill)))

			// log
			logger.Log("responded with a challenge roll message.", log.AddlInfo{
				"input":        m.Content,
				"dice":         diceCount,
				"skill":        skill,
				"author":       m.Author.Username,
				"selectedSide": selectedSide,
				"total":        total,
				"rolls":        rolls,
				"final":        (total + skill),
			})

			// do not check any other commands. one response per message
			return
		}

		// did not mastch to another command.  Default to responding with the help response.
		send(createHelpResponse())
	}
}

type DamageParams struct {
	Dice    int
	MaxRoll int
}

func createHelpResponse() string {
	return "Available formats:\n" +
		"```" +
		"<number of dice>[dD]                                         Example: 6d\n" +
		"<number of dice>[dD]+<skill value>                           Example: 8d+10\n" +
		"damage <number of dice>                                      Example: damage 16\n" +
		"damage <number of dice> cap <maximum roll on a 6-sided die>  Example: damage 9 cap 4\n" +
		"damage <number of dice> cap <number to subtract from 6>      Example: damage 9 cap -2\n" +
		"```"
}

/*

!wd 2d+5 - roll 2 dice, and add 5 skill.

!wd 6d - roll 6 dice

!wd damage 4 cap 4 - roll 4 damage dice, with a die cap of 4 (out of 6)

!wd damage 4 cap -2 - roll 4 damage dice, with a die cap of 4 (out of 6)


*/
