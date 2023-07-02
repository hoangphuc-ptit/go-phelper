package go_phelper

import "log"

var LoggerInstance Logger = simpleLogger{}

type Logger interface {
	LogTrace(format string, v ...any)
	LogDebug(format string, v ...any)
	LogError(format string, v ...any)
}

type simpleLogger struct{}

func (simpleLogger) LogTrace(format string, v ...any) {
	log.Printf("[TRACE] "+format, v...)
}

func (simpleLogger) LogDebug(format string, v ...any) {
	log.Printf("[DEBUG] "+format, v...)
}

func (simpleLogger) LogError(format string, v ...any) {
	log.Printf("[Error] "+format, v...)
}
