package router

import (
	"img-server/core"
)

func Register(r *core.Router) {
	r.POST("/login", Login)
	r.POST("/register", UserRegister)
	r.GET("/menus", GetMenus)
	r.GET("/file",GetFilePath)
}
