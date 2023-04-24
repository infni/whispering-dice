package log

type Logger interface {
	Log(message string, addlInfo AddlInfo)
	LogCritical(message string, addlInfo AddlInfo)
}

const (
	Action         = "action"
	Timestamp      = "timestamp"
	LogLevel       = "level"
	Message        = "message"
	AdditionalInfo = "additional-information"
)
