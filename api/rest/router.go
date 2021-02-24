package rest

import (
	"com.ghl.go/api/rest/hello"
	"github.com/gin-gonic/gin"

	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gitlab.idc.xiaozhu.com/xz-go/common/log"
	"gitlab.idc.xiaozhu.com/xz-go/common/plugins/middleware/trace/zipkin"
)

var route gin.IRouter

func InitRouter() *gin.Engine {
    r := gin.New()

	r.Use(log.GinHandler(), gin.Recovery(), zipkin.Trace())

	r.GET("/", hello.Greeter)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	route = r.Group("/")

	registerRoute()

	return r
}
