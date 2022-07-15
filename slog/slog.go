package slog

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/EwenLan/silicon-server/utils"
)

type slog struct {
	log.Logger
	standardLogOutputOption bool
}

// Init
func (l *slog) Init(path string) {
	directory := utils.GetDirectoryFromPath(path)
	err0 := os.MkdirAll(directory, logPermit)
	if err0 != nil {
		fmt.Printf("make directory failed, err = %s", err0)
	}
	logFile, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, logPermit)
	if err != nil {
		fmt.Printf("open log file failed, err: %s\n", err)
		return
	}
	l.SetOutput(logFile)
}

func (l *slog) output(level string, format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(outputDepth)
	if !ok || file == "" {
		file = "???"
		line = 999
	}
	prefix := fmt.Sprintf("[%s][%s][%s:%d]", level, time.Now().Format(time.RFC3339), utils.GetFilenameFromPath(file), line)
	row := fmt.Sprintf("["+prefix+"["+format+"]]", args...)
	l.Output(0, row)
	if l.standardLogOutputOption {
		fmt.Println(row)
	}
}

func (l *slog) SetStandardLogOutput(option bool) {
	l.standardLogOutputOption = option
	l.Debugf("set standard log option = %t", option)
}

// Debugf
func (l *slog) Debugf(format string, args ...interface{}) {
	l.output(debug, format, args...)
}

// Infof
func (l *slog) Infof(format string, args ...interface{}) {
	l.output(info, format, args...)
}

// Errorf
func (l *slog) Errorf(format string, args ...interface{}) {
	l.output(error, format, args...)
}
