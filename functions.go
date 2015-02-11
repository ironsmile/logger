package logger

import (
	"io"
)

// Log calls the default logger's Log function.
// Arguments are handled in the manner of fmt.Print.
func Log(out string) {
	defaultLogger.Log(out)
}

// Logln calls the default logger's Logln function.
// Arguments are handled in the manner of fmt.Println.
func Logln(out string) {
	defaultLogger.Logln(out)
}

// Logf calls the default logger's Logf function.
// Arguments are handled in the manner of fmt.Printf.
func Logf(format string, args ...interface{}) {
	defaultLogger.Logf(format, args...)
}

// Debug calls the default logger's Debug function.
// Arguments are handled in the manner of fmt.Print.
func Debug(out string) {
	defaultLogger.Debug(out)
}

// Debugln calls the default logger's Debugln function.
// Arguments are handled in the manner of fmt.Println.
func Debugln(out string) {
	defaultLogger.Debugln(out)
}

// Debugf calls the default logger's Debugf function.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// Error calls the default logger's Error function.
// Arguments are handled in the manner of fmt.Print.
func Error(out string) {
	defaultLogger.Error(out)
}

// Errorln calls the default logger's Errorln function.
// Arguments are handled in the manner of fmt.Println.
func Errorln(out string) {
	defaultLogger.Errorln(out)
}

// Errorf calls the default logger's Errorf function.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// SetDebugOutput replaces the debug stream output of the default logger.
func SetDebugOutput(w io.Writer) {
	defaultLogger.SetDebugOutput(w)
}

// SetErrorOutput replaces the error stream output of the default logger.
func SetErrorOutput(w io.Writer) {
	defaultLogger.SetErrorOutput(w)
}

// SetLogOutput replaces the log stream output of the default logger.
func SetLogOutput(w io.Writer) {
	defaultLogger.SetLogOutput(w)
}

// Sets the log level of the default logger.
func SetLevel(level int) {
	defaultLogger.LogLevel = level
}
