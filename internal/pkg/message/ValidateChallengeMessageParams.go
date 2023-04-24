package message

import (
	"fmt"
	"strconv"
)

// ValidateDamageMessageParams validates the results of a Regex match against the needed values for a damage message
// params = the results of a regexp.Regexp.FindAllStringSubmatch() call.
// params[0] contains the results of the Regex application.
// params[0] - the order of theses params is dictated by the regex that was used to make then.  see the DamageRegex function for more details.
func ValidateChallengeMessageParams(params [][]string) (diceCount int, skill int, errStr string) {

	if len(params) == 0 {
		errStr = "Regex params for damage message missing (len=0)."
		return
	}

	// the regex has a bunch of garentee's that we can lean on here to simplify validation
	// like, the '\d'  matcher will never return anything that isn't a number, so Atoi() will never throw error.
	// and, there will always be 4 entries in the array thanks to how the regex package works, due to asking for 3 capture groups.

	match := params[0]

	dice := match[1]
	skillSign := match[2]
	skillNumber := match[3]

	if len(dice) == 0 {
		errStr = fmt.Sprintf("Cannot roll '%s' dice.", dice)
		return
	}
	if len(dice) > 3 {
		errStr = fmt.Sprintf("_Unwilling_ to roll '%s' dice.", dice)
		return
	}
	if len(skillNumber) > 3 {
		errStr = fmt.Sprintf("_Unwilling_ to roll any dice with a skill modifer of '%s%s'.", skillSign, skillNumber)
		return
	}

	diceCount, _ = strconv.Atoi(dice)
	skill, _ = strconv.Atoi(skillNumber)
	if skillSign == "-" {
		skill = 0 - skill
	}

	if diceCount < 1 {
		errStr = fmt.Sprintf("cannot create a result with dice count of '%d'", diceCount)
		return
	}

	// errStr has not been set, so just return the valid count and cap
	return
}
