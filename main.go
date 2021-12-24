package main

import (
	"com.wangzhumo.distribute/conf"
	"com.wangzhumo.distribute/controller"
	"com.wangzhumo.distribute/database"
	"github.com/gin-gonic/gin"
)

func main() {
	// init database
	database.InitDB()
	database.InitDir()
	// database.GenerateModel()
	// router
	engine := gin.Default()
	e := controller.GinRouter(engine)

	// start gin
	gin.SetMode(gin.DebugMode)
	e.Run(conf.Addr())
}
