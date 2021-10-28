package commands

import (
	"context"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jonasgao/gojenkins"
	"os"
	"strings"
)

func printBuild(client *gojenkins.Jenkins, ctx context.Context, jobName string, num int64) error {
	build, err := client.GetBuild(ctx, jobName, num)
	if err != nil {
		return err
	}
	running := build.IsRunning(ctx)
	result := build.GetResult()
	fmt.Printf("Running: %t; Result: %s\n", running, result)
	for i, item := range build.Info().ChangeSet.Items {
		fmt.Printf("%d: %s, %s\n", i, item.Author.FullName, firstLine(item.Comment))
	}
	return nil
}

func firstLine(comment string) string {
	split := strings.Split(comment, "\n")
	return split[0]
}

func printParamTable(jobName string, job *gojenkins.Job) {
	t := table.NewWriter()
	t.SetTitle(fmt.Sprintf("Job %s parameters:", jobName))
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
}
