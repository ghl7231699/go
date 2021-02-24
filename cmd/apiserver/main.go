package main

import (
	"fmt"
	"os"

	"com.ghl.go/api/rest"
	_ "com.ghl.go/docs"
	_ "go.uber.org/automaxprocs"

	"gitlab.idc.xiaozhu.com/xz-go/common/client"
	"gitlab.idc.xiaozhu.com/xz-go/common/config"
	"gitlab.idc.xiaozhu.com/xz-go/common/log"
	"gitlab.idc.xiaozhu.com/xz-go/common/orm"
	"gitlab.idc.xiaozhu.com/xz-go/common/plugins/wrapper/k8straffic"
	"gitlab.idc.xiaozhu.com/xz-go/common/plugins/wrapper/trace/zipkin"
	"gitlab.idc.xiaozhu.com/xz-go/common/server"
	"gitlab.idc.xiaozhu.com/xz-go/common/util/trace"
)

func main() {
	// init common component.
	// there default comment orm and redis component,
	// if your project dependency these component,
	// you should open the comment.
	filePath := "config/app.yaml"
	if os.Getenv("APP_ENV") == "dev" {
		filePath = "config/app_dev.yaml"
	}
	if err := config.InitWithPath(filePath); err != nil {
		panic(fmt.Sprintf("init config failed, err: %v", err))
	}
	log.InitWithPath("log", "prod")
	orm.Setup()
	//redis.Setup()

	// register route and init server
	router := rest.InitRouter()
	srv := server.NewServer(router)

	// Setting trace's service name
	trace.SetServiceName("com.ghl.go")

	// Add client wrappers
	// Out of the box inject xve and trace wrapper
	client.AddDefaultWrappers(k8straffic.WrapperChain)
	client.AddDefaultWrappers(zipkin.WrapperChain)

	srv.Run()
}
