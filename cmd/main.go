package main

import (
	"fmt"
	"healthchecker/config"
	"healthchecker/http"
	"log"
	"os"

	"github.com/pkg/errors"
)

func Run(args []string) {
	defer func() {
		if rvr := recover(); rvr != nil {
			panic(fmt.Sprintf("%s -> app panicked", rvr))
		}
	}()
	command, err := app.Parse(args[1:])
	if err != nil {
		log.Fatal(errors.Wrap(err, "Failed to parse cli command"))
	}
	switch command {
	case runCmd.FullCommand():
		cfg := config.Config{
			StopOnFail: *stopOnFail,
			Port:       *port,
			Timeout:    *maxTimeout,
		}
		http.Run(cfg)
	default:
		panic("failed to parse command string")
	}
}

func main() {
	args := os.Args
	Run(args[1:])
}