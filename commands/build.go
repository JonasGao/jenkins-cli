package commands

import (
	"errors"
	"fmt"
	"github.com/jonasgao/jenkins-cli/jenkins"
	"github.com/urfave/cli/v2"
)

func build() *cli.Command {
	return &cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "Build the job",
		Action: func(c *cli.Context) error {
			if c.NArg() <= 0 {
				return errors.New("please provide a JOB name")
			}
			jobName := c.Args().Get(0)
			client, ctx, err := jenkins.GetClient()
			if err != nil {
				panic(err)
			}
			job, err := client.BuildJob(ctx, jobName, nil)
			fmt.Printf("Build number: %d", job)
			return err
		},
	}
}