package hello

import (
	"github.com/gin-gonic/gin"
	"gitlab.idc.xiaozhu.com/xz-go/common/util/app"
	"net/http"
)

// 问候
// @Router / [GET]
func Greeter(c *gin.Context)  {
	name, found := c.Get("name")
	if !found {
		name = "world"
	}

	data := "hello " + name.(string)

	app.Response(c, http.StatusOK, 200, data)
}


