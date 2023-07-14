package exalted_test

import (
	"fmt"
	"testing"
	"whisperingdice/internal/pkg/exalted"
	"whisperingdice/internal/pkg/log"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

type TestingLogger struct{}

func (l *TestingLogger) Print(v ...interface{}) {
	fmt.Println(v...)
}

func NewTestingLogger() log.Logger {
	return log.NewFileLogger(&TestingLogger{})
}

func TestGetMessage(t *testing.T) {

	optionMap := map[string]*discordgo.ApplicationCommandInteractionDataOption{
		exalted.OptionPool: {
			Name:  exalted.OptionPool,
			Value: float64(15),
			Type:  discordgo.ApplicationCommandOptionInteger,
		},
		exalted.OptionAutomaticSuccesses: {
			Name:  exalted.OptionPool,
			Value: float64(2),
			Type:  discordgo.ApplicationCommandOptionInteger,
		},
	}

	actual, addl := exalted.ActionCommandHandler(optionMap)

	fmt.Println(actual)
	fmt.Println(addl)
	assert.NotEmpty(t, actual, addl)
}
