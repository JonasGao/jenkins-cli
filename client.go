package main

import (
	"context"
	"github.com/bndr/gojenkins"
	"sync"
)

var (
	ctx        context.Context
	initClient sync.Once
	client     *gojenkins.Jenkins
)

func getClient() (*gojenkins.Jenkins, context.Context, error) {
	config, err := readConfig()
	if err != nil {
		return nil, nil, err
	}
	initClient.Do(func() {
		ctx = context.Background()
		client = gojenkins.CreateJenkins(nil, config.Domain, config.Username, config.Password)
		_, err := client.Init(ctx)
		if err != nil {
			return
		}
	})
	return client, ctx, nil
}
