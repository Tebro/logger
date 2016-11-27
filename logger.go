package logger

import (
	"fmt"
	"log"
	"os"
)

var _stderr = log.New(os.Stderr, "Error: ", log.LstdFlags)
var _stdout = log.New(os.Stdout, "Log: ", log.LstdFlags)
var _debug = false

// SetDebug enables the Debug printing
func SetDebug() {
	_debug = true
}

// Log is meant for normal logging
func Log(msg interface{}) {
	_stdout.Println(msg)
}

// Error is ment for logging errors, prints to stderr by default
func Error(msg interface{}) {
	_stderr.Println(msg)
}

// Fatal is ment for errors that exit the program, logs to stderr by default
func Fatal(msg interface{}) {
	_stderr.Fatal(msg)
}

// Debug only prints to stdout if SetDebug has been called
func Debug(msg interface{}) {
	if _debug {
		Log(fmt.Sprintf("DEBUG: %s", msg))
	}
}
