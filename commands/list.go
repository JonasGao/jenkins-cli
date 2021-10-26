package commands

import (
	"fmt"
	"github.com/jonasgao/jenkins-cli/jenkins"
	"github.com/urfave/cli/v2"
)

func list() *cli.Command {
	return &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "List all jobs",
		Action: func(c *cli.Context) error {
			client, ctx, err := jenkins.GetClient()
			if err != nil {
				panic(err)
			}
			jobs, err := client.GetAllJobs(ctx)
			for i := range jobs {
				fmt.Printf("%s\n", jobs[i].GetName())
			}
			return nil
		},
	}
}
