package log

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type FileLogger struct {
	Logger GoLoggerInterface
}

func NewFileLogger(logger GoLoggerInterface) *FileLogger {
	r := new(FileLogger)

	r.Logger = logger

	return r
}

func (logger *FileLogger) Log(message string, addlInfo AddlInfo) {
	logger.logMessage(message, "INFO", addlInfo)
}

func (logger *FileLogger) LogCritical(message string, addlInfo AddlInfo) {
	logger.logMessage(message, "CRITICAL", addlInfo)
}

func (logger *FileLogger) logMessage(message string, logLevel string, addlInfo AddlInfo) {

	data := map[string]interface{}{
		Action:    "file-processing",
		LogLevel:  logLevel,
		Message:   strings.Replace(strings.Replace(message, "\"", "\\\"", -1), "\n", "\\n", -1),
		Timestamp: time.Now().Format(time.RFC3339Nano),
	}
	if len(addlInfo) != 0 {
		data[AdditionalInfo] = addlInfo.ToJson()
	}

	if jsonMsg, jsonErr := json.Marshal(data); jsonErr != nil {
		data["message"] = fmt.Errorf("error serilizing json : %w", jsonErr).Error()
		if jsonMsg2, jsonErr2 := json.Marshal(data); jsonErr != nil {
			fmt.Print(fmt.Errorf("Failed to LOG!! error serilizing json : %w\nOriginal message: %s", jsonErr2, message).Error())
		} else {
			logger.Logger.Print(string(jsonMsg2))
		}
	} else {
		logger.Logger.Print(string(jsonMsg))
	}
}
