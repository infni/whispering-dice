package message

import (
	"fmt"
	"strconv"
)

// ValidateDamageMessageParams validates the results of a Regex match against the needed values for a damage message
// params = the results of a regexp.Regexp.FindAllStringSubmatch() call.
// params[0] contains the results of the Regex application.
// params[0] - the order of theses params is dictated by the regex that was used to make then.  see the DamageRegex function for more details.
func ValidateDamageMessageParams(params [][]string) (diceCount int, dieCap int, errStr string) {

	if len(params) == 0 {
		errStr = "Regex params for damage message missing (len=0)."
		return
	}

	// the regex has a bunch of garentee's that we can lean on here to simplify validation
	// like, the '\d'  matcher will never return anything that isn't a number, so Atoi() will never throw error.
	// and, there will always be 4 entries in the array thanks to how the regex package works, due to asking for 3 capture groups.

	match := params[0]

	dice := match[1]
	capSign := match[2]
	capNumnber := match[3]

	if len(dice) == 0 {
		errStr = fmt.Sprintf("Cannot roll '%s' dice.", dice)
		return
	}
	if len(dice) > 3 {
		errStr = fmt.Sprintf("_Unwilling_ to roll '%s' dice.", dice)
		return
	}

	if len(capSign) > 0 && len(capNumnber) > 1 {
		errStr = fmt.Sprintf("Die cap modifer of '%s' makes no sense. ", capNumnber)
		return
	}

	// SET dieCount and dieCap here
	diceCount, _ = strconv.Atoi(dice)
	if len(capSign) > 0 {
		capAdj, _ := strconv.Atoi(capNumnber)
		if capSign == "-" {
			dieCap = 6 - capAdj
		} else {
			dieCap = 6 + capAdj
		}
	} else if len(capNumnber) > 0 {
		dieCap, _ = strconv.Atoi(capNumnber)
	} else {
		dieCap = 6
	}

	if dieCap < 1 {
		errStr = fmt.Sprintf("cannot create a result with a die cap of '%d'", dieCap)
		return
	}

	if diceCount < 1 {
		errStr = fmt.Sprintf("cannot create a result with dice count of '%d'", diceCount)
		return
	}

	// errStr has not been set, so just return the valid count and cap
	return
}
