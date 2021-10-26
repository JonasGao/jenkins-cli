package commands

import (
	"errors"
	"github.com/jonasgao/jenkins-cli/jenkins"
	"github.com/urfave/cli/v2"
)

func latest() *cli.Command {
	return &cli.Command{
		Name:  "latest",
		Usage: "Get the job latest build",
		Action: func(c *cli.Context) error {
			client, ctx, err := jenkins.GetClient()
			if err != nil {
				return err
			}
			nArg := c.NArg()
			if nArg <= 0 {
				return errors.New("please provide a JOB name")
			}
			jobName := c.Args().Get(0)
			ids, err := client.GetAllBuildIds(ctx, jobName)
			if len(ids) > 0 {
				err = printBuild(client, ctx, jobName, ids[0].Number)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
}
