package debug

import "log"

// Mode indicates whether the application is running in debug mode.
var Mode bool

// Log logs a message if Mode is true.
func Log(format string, v ...interface{}) {
	if Mode {
		log.Printf(format, v...)
	}
}
