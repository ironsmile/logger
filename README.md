# Logger [![GoDoc](https://godoc.org/github.com/ironsmile/logger?status.png)](http://godoc.org/github.com/ironsmile/logger)

This is a simple layer on top of the [standard log package](http://golang.org/pkg/log/). Its main purpose is to give the developer separate log streams and logging severity levels.

The goal of this library is to be as simple as possible to use. It only adds the feature mentioned in the previous paragraph and nothing else.

## Usage

Logging is as simple as 

```go
import (
    "github.com/ironsmile/logger"
)

func main() {
    logger.Errorf("I have an %d errors", 5)
    logger.SetLevel(logger.LevelDebug)
    logger.Debugln("A debug with a new line")
}
```

The above example uses the default logger. You can create as many loggers as you want:

```go
func main() {
    lgStdOut := logger.New()
    lgStdOut.Log("Log object which logs to the stdout")

    outFile, _ := os.Create("/tmp/file.log")
    lgFile := logger.New()
    lgFile.SetLogOutput(outFile)
    lgFile.Log("This will be written in the log file")

    lgFile.SetErrorOutput(os.Stdout)
    lgFile.Errorln("This will be in the standard output")
}
```

## Interface

Loggers support the `Debug(f|ln)?`, `Log(f|ln)?`, `Error(f|ln)?` and `Fatal(f|ln)?` functions which handle their arguments in the way `fmt.Print(f|ln)?` does. For more info see [the documentation](godoc.org/github.com/ironsmile/logger).

## Fatal functions

They do not have their own output stream but use the error stream. So there is no function `SetFatalOutput` and there is no property `Fataller` in the logger struct.

Underneath they actually use the [log.Fatal](http://golang.org/pkg/log/#Fatal) functions which means a call to `logger.Fatal(f|ln)?` will print the message and halt the program.
