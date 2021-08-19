package app

import (
	"github.com/gin-gonic/gin"
	"img-server/auth"
	"img-server/core"
	"img-server/dao"
	"img-server/middle"
	"img-server/router"
	"img-server/util"
	"log"
	"strconv"
)

func NewApi() {
	appConfig := util.ReadConfig()
	dao.NewDb(appConfig.Db)
	var whiteUrl = map[string]interface{}{
		"/v1/api/login": struct {
		}{},
	}
	jwtConfig:=auth.Config{WhiteUrl: whiteUrl,FilePath: appConfig.Casbin}
	app := core.NewApp().Engine
	app.Use(middle.JwtAuth(appConfig,jwtConfig))
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(middle.HandleError())
	app.Static("/assets", "./assets")
	app.StaticFile("/favicon.ico", "./resources/favicon.ico")
	group := app.Group("/v1/api")
	r := core.NewRouter(group)
	router.Register(r)
	err := app.Run(":"+strconv.Itoa(appConfig.Port))
	if err != nil {
		log.Fatal("server err {}", err.Error())
		return
	}
}
