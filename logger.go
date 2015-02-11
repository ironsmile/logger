/*
   Package logger delivers a tiered loggin mechanics. It is a thin wrapper above the
   standard library's log package. Logger gives you the ability to separate your logs
   in 3 types: debug, log and error.

   For fine tuning your loggers you will need to import the log package and use the
   Logger.Debugger, Logger.Logger and Logger.Errorer objects which are just instances
   of log.Logger.

   The package supplies a default logger and a shorthand functions for using it.
   They are similar to the log's default package level functions.
*/
package logger

import (
	"io"
	"log"
	"os"
)

/*
   These constants are used to configer a Logger's output level. They are ordered in
   a strict ascending order. If the level of a logger is set at particular constant
   it will only emit messages in the streams for this constant's level and above.

   For example if a logger has its LogLevel set to LEVEL_LOG it will emit messages
   in its log and error streams but will not emit anything in the debug stream.

   A special level LEVEL_NOLOG may be used in order to silence the logger completely.
*/
const (
	LEVEL_DEBUG = iota
	LEVEL_LOG
	LEVEL_ERROR
	LEVEL_NOLOG // No logs will be emitted whatsoever
)

var (
	defaultLogger *Logger
)

func init() {
	defaultLogger = New()
}

/*
   Logger is a type which actually consistes of 3 log.Logger object. One for every
   stream - debug, log and error.
*/
type Logger struct {

	/*
	   The debug stream. It is a pointer to the underlying log.Logger strcture
	   which is used when Debug, Debugf and Debugln are called.
	*/
	Debugger *log.Logger

	/*
	   The log stream. Used by Log, Logf and Logln.
	*/
	Logger *log.Logger

	/*
	   The error stream. Used by Error, Errorf and Errorln.
	*/
	Errorer *log.Logger

	/*
	   Which Log Level this Logger is using. See the descrption of the
	   log level constants to understend what this means.
	*/
	LogLevel int
}

/*
   Log emits this message to the log stream.
*/
func (l *Logger) Log(out string) {
	if l.LogLevel > LEVEL_LOG {
		return
	}
	l.Logger.Print(out)
}

/*
   Logf emits this message to the log stream. It supports formatting.
   See fmt.Printf for details on the formatting options.
*/
func (l *Logger) Logf(format string, args ...interface{}) {
	if l.LogLevel > LEVEL_LOG {
		return
	}
	l.Logger.Printf(format, args...)
}

/*
   Logln emits this message to the log stream and adds a new line at the end of
   the message. Similar to fmt.Println.
*/
func (l *Logger) Logln(out string) {
	if l.LogLevel > LEVEL_LOG {
		return
	}
	l.Logger.Println(out)
}

/*
   Debug emits this message to the debug stream.
*/
func (l *Logger) Debug(out string) {
	if l.LogLevel > LEVEL_DEBUG {
		return
	}
	l.Debugger.Print(out)
}

/*
   Debugf emits this message to the debug stream. Supports fmt.Printf formatting.
*/
func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.LogLevel > LEVEL_DEBUG {
		return
	}
	l.Debugger.Printf(format, args...)
}

/*
   Debugln emits this message to the debug streamand adds a new line at the end of
   the message. Similar to fmt.Println.
*/
func (l *Logger) Debugln(out string) {
	if l.LogLevel > LEVEL_DEBUG {
		return
	}
	l.Debugger.Println(out)
}

/*
   Error emits this message to the error stream.
*/
func (l *Logger) Error(out string) {
	if l.LogLevel > LEVEL_ERROR {
		return
	}
	l.Errorer.Print(out)
}

/*
   Errorf emits this message to the error stream. Supports fmt.Printf formatting.
*/
func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.LogLevel > LEVEL_ERROR {
		return
	}
	l.Errorer.Printf(format, args...)
}

/*
   Error emits this message to the error stream and adds a new line at the end of the
   message. Similar to fmt.Println.
*/
func (l *Logger) Errorln(out string) {
	if l.LogLevel > LEVEL_ERROR {
		return
	}
	l.Errorer.Println(out)
}

/*
   SetErrorOutput changes the error output stream. It preserves the flags
   and prefix of the old stream.
*/
func (l *Logger) SetErrorOutput(w io.Writer) {
	l.Errorer = log.New(w, l.Errorer.Prefix(), l.Errorer.Flags())
}

/*
   SetDebugOutput changes the debug output stream. It preserves the flags
   and prefix of the old stream.
*/
func (l *Logger) SetDebugOutput(w io.Writer) {
	l.Debugger = log.New(w, l.Debugger.Prefix(), l.Debugger.Flags())
}

/*
   SetLogOutput changes the log output stream. It preserves the flags
   and prefix of the old stream.
*/
func (l *Logger) SetLogOutput(w io.Writer) {
	l.Logger = log.New(w, l.Logger.Prefix(), l.Logger.Flags())
}

// New creates and returns a new Logger. Its debug and log streams
// are stdout and its error stream is stderr.
// The returned logger will have log level set to LEVEL_LOG.
func New() *Logger {
	l := &Logger{}
	l.Debugger = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags)
	l.Logger = log.New(os.Stdout, "[LOG] ", log.LstdFlags)
	l.Errorer = log.New(os.Stderr, "[ERROR] ", log.LstdFlags)
	l.LogLevel = LEVEL_LOG
	return l
}

// Default returns a pointer to the default Logger.
func Default() *Logger {
	return defaultLogger
}
