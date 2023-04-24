package message_test

import (
	"fmt"
	"testing"
	"whisperingdice/internal/pkg/message"

	"github.com/stretchr/testify/assert"
)

func TestChallengeRegex(t *testing.T) {

	challengeRegex := message.NewChallengeRegex()

	testData := []struct {
		input        string
		actualParsed bool
		dice         string
		skillSign    string
		skill        string
	}{
		{"4", true, "4", "", ""},
		{"4+3", true, "4", "+", "3"},
		{"5d", true, "5", "", ""},
		{"6d6", true, "6", "", ""},
		{"7d", true, "7", "", ""},
		{"8d7", true, "8", "", ""},
		{"8d7+3", true, "8", "", ""},
		{"9d+4", true, "9", "+", "4"},
		{"10d-2", true, "10", "-", "2"},
		{"11d+ 4", true, "11", "+", "4"},
		{"12d -2", true, "12", "-", "2"},
		{"13d - 2", true, "13", "-", "2"},
		{"14 d - 2", true, "14", "-", "2"},
		{"15 d-2", true, "15", "-", "2"},
		{"16 d - 2 # as", true, "16", "-", "2"},
		{"17d - 2 #", true, "17", "-", "2"},
		{"18d - 2 # retad ml;kasdf /asd", true, "18", "-", "2"},
	}

	for _, test := range testData {
		actual := challengeRegex.FindAllStringSubmatch(message.CommandPrefix+" "+test.input, -1)
		fmt.Println(actual)
		if test.actualParsed {
			assert.NotNil(t, actual, "The match was nil when it should have matched (did not match) "+test.input)
		} else {
			assert.Nil(t, actual, "The match happened when it should not have matched (not nil) "+test.input)

		}
		// if the result is not here (missing, nil), there is nothing else to check.
		if actual == nil {
			continue
		}
		assert.Equal(t, 1, len(actual), test.input)
		assert.Equal(t, 4, len(actual[0]), test.input)
		assert.Equal(t, test.dice, actual[0][1], test.input)
		assert.Equal(t, test.skillSign, actual[0][2], test.input)
		assert.Equal(t, test.skill, actual[0][3], test.input)
	}
}

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
