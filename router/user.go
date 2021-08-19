package router

import (
	"img-server/dao"

	"github.com/gin-gonic/gin"
)

func GetUserByUsername(ctx *gin.Context) {
	var username = ctx.Query("username")
	user := dao.GetUserByUsername(username)
	if user == nil {
		// helper.Err("MEIY")
		return
	}
	// helper.Ok(user)
}
