package router

import (
	"img-server/core"

	"github.com/gin-gonic/gin"
)

func NewRouters(e *gin.RouterGroup) {
	// e.GET("/menus", GetMenus)
	// e.GET("/user", GetUserByUsername)
	// e.GET("/token", core.Do(GetToken))
	// e.GET("/login", core.Do(Login))
}

func Register(r *core.Router) {
	r.POST("/login", Login)
	r.GET("/menus", GetMenus)
	r.GET("/file",GetFilePath)
}
