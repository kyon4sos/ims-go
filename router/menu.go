package router

import (
	"github.com/gin-gonic/gin"
	. "img-server/core"
	"img-server/dao"
	"img-server/service"
)

func GetMenus(ctx *gin.Context) *Resp {
	menus := dao.GetAllMenus()
	if menus ==nil{
	}
	for i, m := range menus {
		menus[i].Icon = service.GenFilePath(m.Icon)
	}
	//time.Sleep(time.Millisecond * 3000)
	return Err(5000,"粗我")
}
