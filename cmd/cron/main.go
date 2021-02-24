package main

import (
	"fmt"
	"gitlab.idc.xiaozhu.com/xz-go/common/config"
	"os"
	_ "go.uber.org/automaxprocs"
	"com.ghl.go/cmd/cron/cmd"
)

func main() {
	filePath := "config/app.yaml"
	if os.Getenv("APP_ENV") == "dev" {
		filePath = "config/app_dev.yaml"
	}
	if err := config.InitWithPath(filePath); err != nil {
		panic(fmt.Sprintf("init config failed, err: %v", err))
	}
	cmd.Execute()
}
