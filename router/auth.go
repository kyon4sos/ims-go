package router

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"img-server/auth"
	"img-server/core"
	. "img-server/core"
	"img-server/dao"
	"img-server/dto"
	"img-server/model"
	"img-server/service"
)

// 登录 godoc
// @Tags 登录
// @Summary 登录
// @Description 登录
// @Param login body dto.UserDto true "用户登录"
// @Accept  json
// @Produce  json
// @Success 200 {object} core.Resp{msg=string,data=string,code=int} "desc"
// @Router /login [post]
func Login(ctx *gin.Context) *Resp {
	var user dto.UserDto
	err := ctx.ShouldBindJSON(&user)
	if err!=nil {
		return Err(4000,err.Error())
	}
	valid := core.Valid(user)
	if len(valid) >0{
		return Err(40000,valid)
	}
	findUser := dao.GetUserByUsername(user.Username)
	if findUser == nil {
		return Err(40001, "用户名或密码错误")
	}
	matched:= comparePsw(findUser.Password, user.Password)
	if !matched {
		return Err(40001, "用户名或密码错误")
	}
	token := service.GetTokenService().CreateToken(user.Username)
	return Ok(token)

}

// 注册 godoc
// @Tags 注册
// @Summary 注册
// @Description 注册
// @Param dto.UserRegDto body dto.UserRegDto true "用户注册"
// @Accept  json
// @Produce  json
// @Success 200 {object} core.Resp{msg=string,data=string,code=int} "desc"
// @Router /register [post]
func UserRegister(ctx *gin.Context) *Resp {
	var user dto.UserRegDto
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return Err(40001, err.Error())
	}
	findUser := dao.GetUserByUsername(user.Username)
	if findUser != nil {
		return Err(40000, "用户名已存在")
	}
	newUser := model.User{
		Password: string(hashPassword(user.Password)),
		Username: user.Username,
	}
	dao.GetDb().Create(&newUser)
	token := service.GetTokenService().CreateToken(user.Username)
	return Ok(token)
}

func hashPassword(password string) []byte {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil
	}
	return bytes
}
func comparePsw(hash string, plainTxt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(plainTxt))
	if err != nil {
		return false
	}
	return true
}
func GetFilePath(ctx *gin.Context) *Resp {
	path := auth.GetAuthInstance().Enforce.GetPolicy()
	return Ok(path)
}
