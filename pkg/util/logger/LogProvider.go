package logger

import (
	"sync"
)

type LogProvider struct {
	log *logWrapper
}

var loggerInstance *LogProvider
var once sync.Once

// Return the logger instance
func GetLogger() *logWrapper {
	once.Do(func() {
		loggerInstance = newLogProvider()
	})

	return loggerInstance.log
}

// Create and initialize a new log provider
func newLogProvider() *LogProvider {
	return &LogProvider{
		log: newLogWrapper(),
	}
}
