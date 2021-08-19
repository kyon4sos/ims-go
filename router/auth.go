package router

import (
	"github.com/gin-gonic/gin"
	"img-server/auth"
	. "img-server/core"
	"img-server/dto"
	"img-server/service"
)
func Login(ctx *gin.Context) *Resp {
	var user *dto.UserDto
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return nil
	}
	token := service.GetTokenService().CreateToken(user.Username)
	return Ok(token)
}

func GetFilePath(ctx *gin.Context) *Resp{
	path := auth.GetAuthInstance().Enforce.GetPolicy()
	return Ok(path)
}