package noreturn

type Logger interface {
	Log(message string)
	LogWithLevel(level string, message string)
}
