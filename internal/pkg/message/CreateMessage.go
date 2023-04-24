package message

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"whisperingdice/internal/pkg/log"
	"whisperingdice/internal/pkg/rolling"

	"github.com/bwmarrin/discordgo"
)

func CreateResponseFunc(logger log.Logger) func(s *discordgo.Session, m *discordgo.MessageCreate) {

	commandPrefix := "!wd"

	// matched to;
	// the dice - "dice"
	// the cap - "cap"
	// everything after a '#' sign can be used to annotate a roll.  it's all ignored.
	// NOTE:  the match "sign" is optional.  if present, the cap is a negative number.
	// original: ^!wd\s+damage\s+(?P<dice>\d+)(?:\s+cap\s+(?:(?P<sign>[-]))?(?P<cap>\d+))?(?:\s?#.*)*
	/*
		Samples:
			!wd damage 3
			!wd damage 3d
			!wd damage 3d cap 2
			!wd damage 3 cap 2
			!wd damage 32 cap 12
			!wd damage 32 cap -2
			!wd damage 5 cap -12
			!wd damage 5 cap -12 # akosd 12 pa oksk d
			!wd damage 5 cap -12# akosd 12 pa oksk d
			!wd damage 5 cap +3 # akosd 12 pa oksk d
	*/
	damage := regexp.MustCompile(fmt.Sprintf("^%s\\s+damage\\s+(?P<dice>\\d+)(?:\\s+cap\\s+(?:(?P<capsign>[-]))?(?P<cap>\\d+))?(?:\\s?#.*)*", commandPrefix))

	// matched to;
	// the dice - "dice"
	// the cap - "cap"
	// everything after a '#' sign can be used to annotate a roll.  it's all ignored.
	// NOTE:  the match "sign" is optional.  if present, and negative, the cap is a negative number.
	// original: ^!wd\s+(?P<dice>\d+)\s?(?:[dD](?:6)?)?\s?(?:(?P<skillsign>[+-])\s?(?P<skill>\d+))?(?:\s?#.*)*
	/*
		Samples:
			!wd 4
			!wd 4+3
			!wd 4d
			!wd 4d
			!wd 4d6
			!wd 4d
			!wd 4d7
			!wd 4d+4
			!wd 4d-2
			!wd 4d+ 4
			!wd 4d -2
			!wd 4d - 2
			!wd 4 d - 2
			!wd 4 d-2
			!wd 4 d - 2 # as
			!wd 4d - 2 #
			!wd 4d - 2 # retad ml;kasdf /asd
	*/
	challenge := regexp.MustCompile(fmt.Sprintf("^%s\\s+(?P<dice>\\d+)\\s?(?:[dD](?:6)?)?\\s?(?:(?P<skillsign>[+-])\\s?(?P<skill>\\d+))?(?:\\s?#.*)*", commandPrefix))

	return func(s *discordgo.Session, m *discordgo.MessageCreate) {

		send := func(res string) {
			_, err := s.ChannelMessageSend(m.ChannelID, res)
			if err != nil {
				fmt.Println(err)
			}
		}

		// Ignore all messages created by the bot itself
		if m.Author.ID == s.State.User.ID {
			return
		}

		// if it's not a command for us, ignore it
		if !strings.HasPrefix(m.Content, commandPrefix) {
			return
		}

		logger.Log("Recieved message dispatch", log.AddlInfo{})

		// if the command has no content, then just return the help message.
		if strings.Replace(m.Content, commandPrefix, "", 1) == "" {
			send(createHelpResponse())
			return
		}

		// look for different command formats in sequence.  the first one that matches returns.  If nonw, match, return the help message.
		if match := damage.FindAllStringSubmatch(m.Content, -1); match != nil {

			// the regex has a bunch of garentee's that we can lean on here to simplify validation
			// like, the '\d'  matcher will never return anything that isn't a number, so Atoi() will never throw error.
			// and, there will always be 4 entries in the array thanks to how the regex package works, due to asking for 3 capture groups.

			dice := match[0][1]
			capSign := match[0][2]
			capNumnber := match[0][3]

			diceVal, _ := strconv.Atoi(dice)
			var maxSide int
			if len(capSign) > 0 {
				capAdj, _ := strconv.Atoi(capNumnber)
				if capSign == "-" {
					maxSide = 6 - capAdj
				} else {
					maxSide = 6 + capAdj
				}
			} else if len(capNumnber) > 0 {
				maxSide, _ = strconv.Atoi(capNumnber)
			} else {
				maxSide = 6
			}

			if maxSide < 1 {
				// return error about max die cap being less than 1
			}

			if diceVal < 1 {
				// return error about rolling dice being less than 1
			}

			send(fmt.Sprintf("<@%s> %s", m.Author.ID, rolling.RollDamageMessage(diceVal, maxSide)))

			logger.Log("matched to a damage roll.", log.AddlInfo{
				"dice":      dice,
				"capSign":   capSign,
				"capNumber": capNumnber,
				"author":    m.Author.Username,
				"maxSide":   maxSide,
			})

			return
		}

		if match := challenge.FindAllStringSubmatch(m.Content, -1); match != nil {
			// the regex has a bunch of garentee's that we can lean on here to simplify validation
			// like, the '\d'  matcher will never return anything that isn't a number, so Atoi() will never throw error.
			// and, there will always be 4 entries in the array thanks to how the regex package works, due to asking for 3 capture groups.

			dice := match[0][1]
			skillSign := match[0][2]
			skillNumber := match[0][3]

			diceVal, _ := strconv.Atoi(dice)
			skillNumberVal, _ := strconv.Atoi(skillNumber)
			if skillSign == "-" {
				skillNumberVal = 0 - skillNumberVal
			}

			send(fmt.Sprintf("<@%s> %s", m.Author.ID, rolling.RollChallengeMessage(diceVal, skillNumberVal)))

			logger.Log("matched to a challenge roll.", log.AddlInfo{
				"dice":           dice,
				"skillSign":      skillSign,
				"skillNumber":    skillNumber,
				"author":         m.Author.Username,
				"skillNumberVal": skillNumberVal,
			})

			return
		}

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
