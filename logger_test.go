package logger_test

import (
	"os"

	"github.com/ironsmile/logger"
)

func ExampleLogger() {

	// Create a new logger
	lg := logger.New()

	// Using the underlying standard's library log.Logger instances
	// to reconfiger the debug stream.
	lg.Debugger.SetPrefix("")
	lg.Debugger.SetFlags(0)

	// Now all types of messages will be outputed.
	lg.Level = logger.LevelDebug

	// Emit something to the debug stream
	lg.Debugf("debug yourself and you will be %d time stronger", 42)
	// Output:
	// debug yourself and you will be 42 time stronger
}

func ExampleLogger_configuration() {
	lg := logger.New()

	// Increase this logger's verbosity
	lg.Level = logger.LevelDebug

	// Lets be sure all error messages have the '[ERROR] ' string in front of them
	lg.Errorer.SetPrefix("[ERROR] ")

	// We do not want anything in the message except our prefix.
	lg.Errorer.SetFlags(0)

	// We want errors to be seen in the standard output.
	lg.SetErrorOutput(os.Stdout)

	// Emit something to the error stream
	lg.Error("error to stdout")
	// Output: [ERROR] error to stdout
}

func ExampleDefault() {

	// Reconfiguring the default logger
	logger.Default().Debugger.SetFlags(0)
	logger.Default().Debugger.SetPrefix("level.debug: ")

	// Setting its level
	logger.SetLevel(logger.LevelDebug)

	// Setting its output destination
	logger.SetDebugOutput(os.Stdout)

	// Actually printing something in the output destination
	logger.Debugln("the default may be more than you need")
	// Output: level.debug: the default may be more than you need
}
