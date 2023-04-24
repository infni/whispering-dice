package message_test

import (
	"fmt"
	"testing"
	"whisperingdice/internal/pkg/message"

	"github.com/stretchr/testify/assert"
)

func TestCreateDamageMessage_WhenNilParams_ThenError(t *testing.T) {

	var params [][]string = nil

	_, _, actual := message.ValidateDamageMessageParams(params)

	assert.Contains(t, actual, "Regex params for damage message missing", fmt.Sprintf("'%s'", actual))
}

func TestCreateDamageMessage_WhenBlankDice_ThenError(t *testing.T) {

	params := [][]string{{"", "", "", ""}}

	_, _, actual := message.ValidateDamageMessageParams(params)

	assert.Contains(t, actual, "Cannot roll ", fmt.Sprintf("'%s'", actual))
}

func TestCreateDamageMessage_WhenZeroDice_ThenError(t *testing.T) {

	params := [][]string{{"", "0", "", ""}}

	_, _, actual := message.ValidateDamageMessageParams(params)

	assert.Contains(t, actual, "cannot create a result with dice count of ", fmt.Sprintf("'%s'", actual))
}

func TestCreateDamageMessage_WhenTooManyDice_ThenError(t *testing.T) {

	params := [][]string{{"", "1000", "", ""}}

	_, _, actual := message.ValidateDamageMessageParams(params)

	assert.Contains(t, actual, "_Unwilling_ to roll ", fmt.Sprintf("'%s'", actual))
}

func TestCreateDamageMessage_WhenCapModifierTooLarge_ThenError(t *testing.T) {

	params := [][]string{{"", "1", "-", "10"}}

	_, _, actual := message.ValidateDamageMessageParams(params)

	assert.Contains(t, actual, "Die cap modifer of ", fmt.Sprintf("'%s'", actual))
}

func TestCreateDamageMessage_WhenCapModifierTooSmall_ThenError(t *testing.T) {

	paramsArray := [][][]string{
		{{"", "1", "", "0"}},
		{{"", "1", "-", "6"}},
		{{"", "1", "-", "7"}},
	}

	for _, params := range paramsArray {
		_, _, actual := message.ValidateDamageMessageParams(params)

		assert.Contains(t, actual, "cannot create a result with a die cap of ", fmt.Sprintf("'%s'", actual))
	}
}

func TestCreateDamageMessage_WhenCapModifierNegativeAndRollsAreOverCap_THenResultsHaveFormattedFrollsAndNonFormattedRolls(t *testing.T) {

	paramsArray := [][][]string{
		{{"", "3", "", "0"}},
		{{"", "1", "-", "6"}},
		{{"", "1", "-", "7"}},
	}

	for _, params := range paramsArray {
		_, _, actual := message.ValidateDamageMessageParams(params)

		assert.Contains(t, actual, "cannot create a result with a die cap of ", fmt.Sprintf("'%s'", actual))
	}
}
