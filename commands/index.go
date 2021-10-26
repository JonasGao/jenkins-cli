package commands

import (
	"github.com/urfave/cli/v2"
)

func GetCommands() []*cli.Command {
	return []*cli.Command{
		list(),
		build(),
		get(),
		latest(),
	}
}
