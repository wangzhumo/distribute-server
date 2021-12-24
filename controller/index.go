package controller

import (
	apk "com.wangzhumo.distribute/controller/v1/apk"
	info "com.wangzhumo.distribute/controller/v1/info"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GinRouter 加载所有路由
func GinRouter(engine *gin.Engine) *gin.Engine {
	// 静态文件
	engine.StaticFS("/static", http.Dir("./.data"))
	// Api 路由
	rr := engine.Group("/v1")
	info.Routers(rr)
	apk.Routers(rr)
	return engine
}
