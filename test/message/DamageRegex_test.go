package message_test

import (
	"fmt"
	"testing"
	"whisperingdice/internal/pkg/message"

	"github.com/stretchr/testify/assert"
)

func TestDamageRegex(t *testing.T) {

	damageRegex := message.NewDamageRegex()

	testData := []struct {
		name         string
		input        string
		actualParsed bool
		dice         string
		capSign      string
		cap          string
	}{
		{"", "damage 3", true, "3", "", ""},
		{"", "damage 3d", true, "3", "", ""},
		{"", "damage 3d cap 2", true, "3", "", "2"},
		{"", "damage 3 cap 2", true, "3", "", "2"},
		{"", "damage 32 cap 12", true, "32", "", "12"},
		{"", "damage 9 cap-2", true, "9", "-", "2"},
		{"", "damage 11 cap -2", true, "11", "-", "2"},
		{"", "damage 45 cap - 2", true, "45", "-", "2"},
		{"", "damage 56 cap- 2", true, "56", "-", "2"},
		{"", "damage 999 cap -3", true, "999", "-", "3"},
		{"", "damage 5 cap -12#stuff", true, "5", "-", "12"},
		{"", "damage 5 cap 3 #stuff with spaces", true, "5", "", "3"},
		{"", "damage 5 cap 4 # stuuf with 1234", true, "5", "", "4"},
		{"", "damage 1000 cap -3", true, "1000", "-", "3"},
		{"", "damage 5cap 4", true, "5", "", "4"},
		{"", "damage 5 cap4", true, "5", "", "4"},
		{"", "damage 5 cap4", true, "5", "", "4"},
		{"", "damage5 cap 4", true, "5", "", "4"},
		{"", "damage 5 cap +4", true, "5", "+", "4"},
	}

	for _, test := range testData {
		actual := damageRegex.FindAllStringSubmatch(message.CommandPrefix+" "+test.input, -1)
		fmt.Println(actual)
		if test.actualParsed {
			assert.NotNil(t, actual, "The match was nil when it should have matched (did not match)")
		} else {
			assert.Nil(t, actual, "The match happened when it should not have matched (not nil)")

		}
		// if the result is not here (missing, nil), there is nothing else to check.
		if actual == nil {
			continue
		}
		assert.Equal(t, 1, len(actual), test.input)
		assert.Equal(t, 4, len(actual[0]), test.input)
		assert.Equal(t, test.dice, actual[0][1], test.input)
		assert.Equal(t, test.capSign, actual[0][2], test.input)
		assert.Equal(t, test.cap, actual[0][3], test.input)
	}
}

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
