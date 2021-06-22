package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
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
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
