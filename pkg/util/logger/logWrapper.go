package logger

import (
	"github.com/google/logger"
	"io/ioutil"
	"os"
)

var logPath = "."

func SetLogPath(path string) {
	logPath = path
}

type logLevel struct {
	name  string
	value int
}

var levelInfo = logLevel{name: "INFO", value: 0}
var levelWarn = logLevel{name: "WARN", value: 1}
var levelError = logLevel{name: "ERROR", value: 2}
var levelFatal = logLevel{name: "FATAL", value: 3}

var supportedLogLevels = []*logLevel{
	&levelInfo,
	&levelWarn,
	&levelError,
	&levelFatal,
}

type logWrapper struct {
	log          *logger.Logger
	cliLogger    *logger.Logger
	minimumLevel *logLevel
	forceCli     bool
}

// Initializes the logger and creates a new wrapper object
func newLogWrapper() *logWrapper {
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}

	l := logger.Init("FNMSLogger", true, false, logFile)
	cliL := logger.Init("FNMS_CLI_Logger", true, false, ioutil.Discard)

	return &logWrapper{
		log:          l,
		cliLogger:    cliL,
		minimumLevel: &levelInfo,
		forceCli:     false,
	}
}

// Returns a pointer to the available log level or nil
func GetLogLevelForName(logLevelName string) *logLevel {
	for _, availableLevel := range supportedLogLevels {
		if availableLevel.name == logLevelName {
			return availableLevel
		}
	}

	return nil
}

// Set if the logger should always (independent of the log level output to the cli)
func (lw *logWrapper) SetForceCli(forceCli bool) {
	lw.forceCli = forceCli
}

// Sets the log level
func (lw *logWrapper) SetLogLevel(logLevelName string) {
	ll := GetLogLevelForName(logLevelName)
	if ll != nil {
		lw.minimumLevel = ll
		return
	}

	availableLevelsString := ""

	for i, availableLevel := range supportedLogLevels {
		availableLevelsString += availableLevel.name
		if i != len(supportedLogLevels)-1 {
			availableLevelsString += ", "
		}
	}

	lw.log.Fatalf("No log level %v. Possible levels are %v", logLevelName, availableLevelsString)
}

// Checks if the passed log level is over the minimum and therefore allowed to print the log
func (lw *logWrapper) isOverMinimum(level *logLevel) bool {
	return lw.minimumLevel.value <= level.value
}

// Returns the log instance for the passed level
func (lw *logWrapper) getLogInstance(level *logLevel) *logger.Logger {
	if lw.isOverMinimum(level) {
		return lw.log
	} else if lw.forceCli {
		return lw.cliLogger
	}

	return nil
}

// Info logs with the Info severity.
// Arguments are handled in the manner of fmt.Print.
func (lw *logWrapper) Info(v ...interface{}) {
	logInstance := lw.getLogInstance(&levelInfo)

	if logInstance != nil {
		logInstance.Info(v...)
	}
}

// InfoDepth acts as Info but uses depth to determine which call frame to log.
// InfoDepth(0, "msg") is the same as Info("msg").
func (lw *logWrapper) InfoDepth(depth int, v ...interface{}) {
	logInstance := lw.getLogInstance(&levelInfo)

	if logInstance != nil {
		logInstance.InfoDepth(depth, v...)
	}
}

// Infoln logs with the Info severity.
// Arguments are handled in the manner of fmt.Println.
func (lw *logWrapper) Infoln(v ...interface{}) {
	logInstance := lw.getLogInstance(&levelInfo)

	if logInstance != nil {
		logInstance.Infoln(v...)
	}
}

// Infof logs with the Info severity.
// Arguments are handled in the manner of fmt.Printf.
func (lw *logWrapper) Infof(format string, v ...interface{}) {
	logInstance := lw.getLogInstance(&levelInfo)

	if logInstance != nil {
		logInstance.Infof(format, v...)
	}
}

// Warning logs with the Warning severity.
// Arguments are handled in the manner of fmt.Print.
func (lw *logWrapper) Warning(v ...interface{}) {
	logInstance := lw.getLogInstance(&levelWarn)

	if logInstance != nil {
		logInstance.Warning(v...)
	}
}

// WarningDepth acts as Warning but uses depth to determine which call frame to log.
// WarningDepth(0, "msg") is the same as Warning("msg").
func (lw *logWrapper) WarningDepth(depth int, v ...interface{}) {
	logInstance := lw.getLogInstance(&levelWarn)

	if logInstance != nil {
		logInstance.WarningDepth(depth, v...)
	}
}

// Warningln logs with the Warning severity.
// Arguments are handled in the manner of fmt.Println.
func (lw *logWrapper) Warningln(v ...interface{}) {
	logInstance := lw.getLogInstance(&levelWarn)

	if logInstance != nil {
		logInstance.Warningln(v...)
	}
}

// Warningf logs with the Warning severity.
// Arguments are handled in the manner of fmt.Printf.
func (lw *logWrapper) Warningf(format string, v ...interface{}) {
	logInstance := lw.getLogInstance(&levelWarn)

	if logInstance != nil {
		logInstance.Warningf(format, v...)
	}
}

// Error logs with the ERROR severity.
// Arguments are handled in the manner of fmt.Print.
func (lw *logWrapper) Error(v ...interface{}) {
	logInstance := lw.getLogInstance(&levelError)

	if logInstance != nil {
		logInstance.Error(v...)
	}
}

// ErrorDepth acts as Error but uses depth to determine which call frame to log.
// ErrorDepth(0, "msg") is the same as Error("msg").
func (lw *logWrapper) ErrorDepth(depth int, v ...interface{}) {
	logInstance := lw.getLogInstance(&levelError)

	if logInstance != nil {
		logInstance.ErrorDepth(depth, v...)
	}
}

// Errorln logs with the ERROR severity.
// Arguments are handled in the manner of fmt.Println.
func (lw *logWrapper) Errorln(v ...interface{}) {
	logInstance := lw.getLogInstance(&levelError)

	if logInstance != nil {
		logInstance.Errorln(v...)
	}
}

// Errorf logs with the Error severity.
// Arguments are handled in the manner of fmt.Printf.
func (lw *logWrapper) Errorf(format string, v ...interface{}) {
	logInstance := lw.getLogInstance(&levelError)

	if logInstance != nil {
		logInstance.Errorf(format, v...)
	}
}

// Fatal logs with the Fatal severity, and ends with os.Exit(1).
// Arguments are handled in the manner of fmt.Print.
func (lw *logWrapper) Fatal(v ...interface{}) {
	logInstance := lw.getLogInstance(&levelFatal)

	if logInstance != nil {
		logInstance.Fatal(v...)
	}
}

// FatalDepth acts as Fatal but uses depth to determine which call frame to log.
// FatalDepth(0, "msg") is the same as Fatal("msg").
func (lw *logWrapper) FatalDepth(depth int, v ...interface{}) {
	logInstance := lw.getLogInstance(&levelFatal)

	if logInstance != nil {
		logInstance.FatalDepth(depth, v...)
	}
}

// Fatalln logs with the Fatal severity, and ends with os.Exit(1).
// Arguments are handled in the manner of fmt.Println.
func (lw *logWrapper) Fatalln(v ...interface{}) {
	logInstance := lw.getLogInstance(&levelFatal)

	if logInstance != nil {
		logInstance.Fatalln(v...)
	}
}

// Fatalf logs with the Fatal severity, and ends with os.Exit(1).
// Arguments are handled in the manner of fmt.Printf.
func (lw *logWrapper) Fatalf(format string, v ...interface{}) {
	logInstance := lw.getLogInstance(&levelFatal)

	if logInstance != nil {
		logInstance.Fatalf(format, v...)
	}
}

// Close closes all the underlying log writers, which will flush any cached logs.
// Any errors from closing the underlying log writers will be printed to stderr.
// Once Close is called, all future calls to the logger will panic.
func (lw *logWrapper) Close() {
	lw.log.Close()
}
