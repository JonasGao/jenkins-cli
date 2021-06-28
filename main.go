package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
)

func main() {
	app := &cli.App{
		Name:  "jk",
		Usage: "Make Jenkins like a nice JK",
		Commands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "List all jobs",
				Action: func(c *cli.Context) error {
					client, ctx, err := getClient()
					if err != nil {
						panic(err)
					}
					jobs, err := client.GetAllJobs(ctx)
					for i := range jobs {
						fmt.Printf("%s\n", jobs[i].GetName())
					}
					return nil
				},
			},
			{
				Name:    "build",
				Aliases: []string{"b"},
				Usage:   "Build the job",
				Action: func(c *cli.Context) error {
					if c.NArg() <= 0 {
						return errors.New("please provide a JOB name")
					}
					jobName := c.Args().Get(0)
					client, ctx, err := getClient()
					if err != nil {
						panic(err)
					}
					job, err := client.BuildJob(ctx, jobName, nil)
					fmt.Printf("Build number: %d", job)
					return err
				},
			},
			{
				Name:  "get",
				Usage: "Get the job builds",
				Action: func(c *cli.Context) error {
					client, ctx, err := getClient()
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
			},
			{
				Name:  "latest",
				Usage: "Get the job latest build",
				Action: func(c *cli.Context) error {
					client, ctx, err := getClient()
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
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
