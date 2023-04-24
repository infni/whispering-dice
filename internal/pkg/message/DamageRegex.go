package message

import (
	"fmt"
	"regexp"
)

func NewDamageRegex() *regexp.Regexp {
	// matched to;
	// the dice - "dice"
	// the cap - "cap"
	// everything after a '#' sign can be used to annotate a roll.  it's all ignored.
	// NOTE:  the match "sign" is optional.  if present, the cap is a negative number.
	str := `\s+damage\s*(?P<dice>\d+)(?:[dD])?(?:\s*cap\s*(?:(?P<sign>[-+]))?\s*(?P<cap>\d+))?(?:\s?#.*)*`
	return regexp.MustCompile(fmt.Sprintf("^%s%s", CommandPrefix, str))
}
