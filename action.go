package main

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
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
		fmt.Printf("%d: %s, %s\n", i, item.Author.FullName, item.Comment)
	}
	return nil
}
