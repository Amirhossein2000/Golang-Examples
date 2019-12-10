package main

import (
	"github.com/apsdehal/go-logger"
	"os"
)

func main() {
	// Get the instance for logger class, "test" is the module name, 1 is used to
	// state if we want coloring
	// Third option is optional and is instance of type io.Writer, defaults to os.Stderr

	f, err := os.OpenFile("l.log", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		f, err = os.Create("l.log")

		if err != nil {
			panic(err)
		}
	}

	log, err := logger.New("test", 1, f)
	if err != nil {
		panic(err) // Check for error
	}

	log.SetLogLevel(logger.DebugLevel)

	log.Debug("dddddd")

	log.Error("errr")
}
