package message

import (
	"fmt"
	"regexp"
)

func NewChallengeRegex() *regexp.Regexp {
	// matched to;
	// the dice - "dice"
	// the cap - "cap"
	// everything after a '#' sign can be used to annotate a roll.  it's all ignored.
	// NOTE:  the match "sign" is optional.  if present, and negative, the cap is a negative number.
	str := `\s+(?P<dice>\d+)\s?(?:[dD](?:6)?)?\s?(?:(?P<skillsign>[+-])\s?(?P<skill>\d+))?(?:\s?#.*)*`
	return regexp.MustCompile(fmt.Sprintf("^%s%s", CommandPrefix, str))
}
