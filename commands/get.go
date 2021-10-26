package commands

import (
	"errors"
	"fmt"
	"github.com/jonasgao/jenkins-cli/jenkins"
	"github.com/urfave/cli/v2"
	"strconv"
)

func get() *cli.Command {
	return &cli.Command{
		Name:  "get",
		Usage: "Get the job builds",
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
			if nArg == 1 {
				ids, err := client.GetAllBuildIds(ctx, jobName)
				if err != nil {
					return err
				}
				for _, id := range ids {
					fmt.Printf("%d: %s\n", id.Number, id.URL)
				}
			} else {
				buildNumber := c.Args().Get(1)
				num, err := strconv.ParseInt(buildNumber, 10, 64)
				if err != nil {
					return err
				}
				err = printBuild(client, ctx, jobName, num)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
}
