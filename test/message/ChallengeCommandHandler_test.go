package message_test

import (
	"fmt"
	"whisperingdice/internal/pkg/log"
)

type TestingLogger struct{}

func (l *TestingLogger) Print(v ...interface{}) {
	fmt.Println(v...)
}

func NewTestingLogger() log.Logger {
	return log.NewFileLogger(&TestingLogger{})
}

// func TestGetMessage(t *testing.T) {
// 	regex := message.NewChallengeRegex()
// 	inputStr := message.CommandPrefix + " " + "8+6"
// 	inputMatch := regex.FindAllStringSubmatch(inputStr, -1)

// 	handler := message.NewChallengeCommandHandler()

// 	handler.Validate(inputMatch)

// 	actual := handler.GetMessage(inputMatch)

// 	assert.NotEmpty(t, actual, inputStr)
// }

// func TestGetMessageLogging(t *testing.T) {
// 	regex := message.NewChallengeRegex()
// 	inputStr := message.CommandPrefix + " " + "8+6"
// 	inputMatch := regex.FindAllStringSubmatch(inputStr, -1)

// 	handler := message.NewChallengeCommandHandler()

// 	handler.Validate(inputMatch)

// 	actual := handler.GetMessage(inputMatch)

// 	logger := NewTestingLogger()

// 	logger.Log("test Message", handler.AddlInfo())

// 	assert.NotEmpty(t, actual, inputStr)
// }
