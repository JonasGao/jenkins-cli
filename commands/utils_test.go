package commands

import (
	"github.com/jonasgao/jenkins-cli/jenkins"
	"testing"
)

func TestPrintBuild(t *testing.T) {
	client, ctx, err := jenkins.GetClient()
	if err != nil {
		t.Errorf("获取客户端报错: %s", err)
		return
	}
	const job = "prod-voerp-account-web"
	ids, err := client.GetAllBuildIds(ctx, job)
	err = printBuild(client, ctx, job, ids[0].Number)
	if err != nil {
		t.Errorf("打印状态报错: %s", err)
	}
}
