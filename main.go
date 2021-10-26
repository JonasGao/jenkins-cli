package main

import (
	"github.com/jonasgao/jenkins-cli/commands"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:     "jk",
		Usage:    "Make Jenkins like a nice JK",
		Commands: commands.GetCommands(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
