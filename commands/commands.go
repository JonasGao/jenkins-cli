package commands

import (
	"errors"
	"fmt"
	"github.com/jonasgao/jenkins-cli/jenkins"
	"github.com/urfave/cli/v2"
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
				printParamTable(jobName, job)
				return err
			},
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "params",
					Aliases: []string{"p"},
				},
			},
			Usage: "Build the job",
			Action: func(c *cli.Context) error {
				if c.NArg() <= 0 {
					return errors.New("please provide a JOB name")
				}
				jobName := c.Args().Get(0)
				client, ctx, err := jenkins.GetClient()
				job, err := client.GetJob(ctx, jobName)
				if err != nil {
					panic(err)
				}
				parameters, err := job.GetParameters(ctx)
				// 如果 job 的定义里有参数，则先打印参数
				if len(parameters) != 0 {
					printParamTable(jobName, job)
				}
				// 然后检查是否提供了参数
				value := c.String("params")
				var p map[string]string
				if len(value) != 0 {
					fmt.Printf("\nWill use parameters:\n")
					p = make(map[string]string)
					entries := strings.Split(value, ";")
					for _, e := range entries {
						fmt.Println(e)
						parts := strings.Split(e, "=")
						p[parts[0]] = parts[1]
					}
				}
				buildNum, err := client.BuildJob(ctx, jobName, p)
				if err != nil {
					panic(err)
				}
				fmt.Printf("\nBuild number: %d", buildNum)
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
