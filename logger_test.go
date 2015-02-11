package logger_test

import (
	"os"

	"github.com/ironsmile/logger"
)

func ExampleLogger() {

	// Create a new logger
	lg := logger.New()

	lg.Debugger.SetPrefix("")
	lg.Debugger.SetFlags(0)

	// Now all types of messages will be outputed.
	lg.LogLevel = logger.LEVEL_DEBUG

	lg.Debugf("debug myself and you will be %d time stronger", 42)
	// Output:
	// debug myself and you will be 42 time stronger
}

func ExampleLogger_configuration() {
	lg := logger.New()

	// Increase this logger's verbosity
	lg.LogLevel = logger.LEVEL_DEBUG

	// Lets be sure all error messages have the '[ERROR] ' string in front of them
	lg.Errorer.SetPrefix("[ERROR] ")

	// We do not want anything in the message except our prefix.
	lg.Errorer.SetFlags(0)

	// We want errors to be seen in the standard output.
	lg.SetErrorOutput(os.Stdout)

	lg.Error("debug to stdout")
	// Output: [ERROR] debug to stdout
}

func Example_default_logger() {
	logger.Default().Debugger.SetFlags(0)
	logger.Default().Debugger.SetPrefix("level.debug: ")
	logger.SetLevel(logger.LEVEL_DEBUG)
	logger.SetDebugOutput(os.Stdout)

	logger.Debugln("the default may be more than you need")
	// Output: level.debug: the default may be more than you need
}
