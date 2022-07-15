package slog

import (
	"fmt"
	"time"
)

var globalLogger slog

func getTodayLogFilename() string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf("%04d%02d%02d.log", year, month, day)
}

func init() {
	globalLogger.Init(defaultLogDirectory + getTodayLogFilename())
}

// SetStandardLogOutput
func SetStandardLogOutput(option bool) {
	globalLogger.SetStandardLogOutput(option)
	Debugf("set global standard log output = %t", option)
}

// Debugf
func Debugf(format string, args ...interface{}) {
	globalLogger.Debugf(format, args...)
}

// Infof
func Infof(format string, args ...interface{}) {
	globalLogger.Infof(format, args...)
}

// Errorf
func Errorf(format string, args ...interface{}) {
	globalLogger.Errorf(format, args...)
}
