package app

import (
	"github.com/gin-gonic/gin"
	"img-server/auth"
	"img-server/core"
	"img-server/dao"
	"img-server/doc"
	_ "img-server/docs"
	"img-server/middle"
	"img-server/router"
	"img-server/util"
	"log"
	"strconv"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9999
// @BasePath /v1/api
// @query.collection.format multi
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @x-extension-openapi {"example": "value on a json format"}

func NewApi() {
	appConfig := util.ReadConfig()
	dao.NewDb(appConfig.Db)
	core.UseValidate()
	app := core.NewApp().Engine
	jwtConfig := getJwtConfig(appConfig)
	app.Use(middle.JwtAuth(appConfig,jwtConfig))
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Use(middle.HandleError())
	app.Static("/assets", "./assets")
	app.StaticFile("/favicon.ico", "./resources/favicon.ico")
	group := app.Group("/v1/api")
	group.GET("/swagger/*any",doc.InitSwagger())
	r := core.NewRouter(group)
	router.Register(r)
	err := app.Run(":"+strconv.Itoa(appConfig.Port))
	if err != nil {
		log.Fatal("server err {}", err.Error())
		return
	}
}




func getJwtConfig(appConfig core.AppConfig) auth.Config {
	var whiteUrl = map[string]interface{}{
		"^/v1/api/register": struct {
		}{},
		"^/v1/api/swagger/": struct {
		}{},
		"^/v1/api/login": struct {
		}{},
	}
	return auth.Config{WhiteUrl: whiteUrl,FilePath: appConfig.Casbin}
}