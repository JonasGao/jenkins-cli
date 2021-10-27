package commands

import (
	"errors"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jonasgao/jenkins-cli/jenkins"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
	"strings"
)

func GetCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Get the job",
			Action: func(c *cli.Context) error {
				if c.NArg() <= 0 {
					return errors.New("please provide a JOB name")
				}
				jobName := c.Args().Get(0)
				client, ctx, err := jenkins.GetClient()
				if err != nil {
					panic(err)
				}
				job, err := client.GetJob(ctx, jobName)
				if err != nil {
					panic(err)
				}
				_, err = job.Poll(ctx)
				if err != nil {
					panic(err)
				}
				t := table.NewWriter()
				t.SetOutputMirror(os.Stdout)
				t.AppendHeader(table.Row{"#", "Name", "Type", "Default", "Desc / Choices"})
				for i, s := range job.Raw.Property {
					for i2, definition := range s.ParameterDefinitions {
						var desc string
						if len(definition.Choices) == 0 {
							desc = definition.Description
						} else {
							desc = strings.Join(definition.Choices, ", ")
						}
						t.AppendRow([]interface{}{
							i + i2,
							definition.Name,
							definition.Type,
							definition.DefaultParameterValue.Value,
							desc,
						})
					}
				}
				t.Render()
				return err
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
				client, ctx, err := jenkins.GetClient()
				if err != nil {
					panic(err)
				}
				job, err := client.BuildJob(ctx, jobName, nil)
				fmt.Printf("Build number: %d", job)
				return err
			},
		},
		{
			Name:    "history",
			Aliases: []string{"h"},
			Usage:   "Get the job history builds",
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
		},
		{
			Name:    "latest",
			Usage:   "Get the job latest build",
			Aliases: []string{"lt"},
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
		},
		{
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
		},
	}
}
