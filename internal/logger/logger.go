package logger

import (
	"log"
	"os"
)

// Adapted from https://stackoverflow.com/a/66735891
type Logger struct {
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
	Debug   *log.Logger
}

func (l *Logger) Setup(output *os.File) {
	// Info writes logs in the color blue with "INFO: " as prefix
	l.Info = log.New(output, "\u001b[34mINFO: \u001B[0m", log.LstdFlags|log.Lshortfile)
	// Warning writes logs in the color yellow with "WARNING: " as prefix
	l.Warning = log.New(output, "\u001b[33mWARNING: \u001B[0m", log.LstdFlags|log.Lshortfile)
	// Error writes logs in the color red with "ERROR: " as prefix
	l.Error = log.New(output, "\u001b[31mERROR: \u001b[0m", log.LstdFlags|log.Lshortfile)
	// Debug writes logs in the color cyan with "DEBUG: " as prefix
	l.Debug = log.New(output, "\u001b[36mDEBUG: \u001B[0m", log.LstdFlags|log.Lshortfile)
}
