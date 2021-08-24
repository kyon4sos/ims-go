package router

import (
	"github.com/gin-gonic/gin"
	. "img-server/core"
	"img-server/dao"
	"img-server/service"
)

// 菜单 godoc
// @Summary 菜单
// @Description 获取菜单
// @Accept  json
// @Produce  json
// @Success 200 {object} core.Resp{data=model.Menu}
// @Router /menus [get]
func GetMenus(ctx *gin.Context) *Resp {
	menus := dao.GetAllMenus()
	if menus ==nil{
	}
	for i, m := range menus {
		menus[i].Icon = service.GenFilePath(m.Icon)
	}
	//time.Sleep(time.Millisecond * 3000)
	return Ok(menus)
}
