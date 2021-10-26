package jenkins

import (
	"context"
	"github.com/bndr/gojenkins"
	"github.com/jonasgao/jenkins-cli/conf"
	"sync"
)

var (
	ctx        context.Context
	initClient sync.Once
	client     *gojenkins.Jenkins
)

func GetClient() (*gojenkins.Jenkins, context.Context, error) {
	config, err := conf.ReadConfig()
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
