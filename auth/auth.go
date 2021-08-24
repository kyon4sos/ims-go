package auth

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"img-server/core"
	"img-server/service"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const (
	Header = "Authorization"
	Prefix = "Bearer "
)

var authInstance *Auth

type Config struct {
	WhiteUrl map[string]interface{}
	FilePath string
}

func (config *Config) IsWhiteUrl(ctx *gin.Context) bool {
	url := ctx.Request.URL.String()
	for k,_:=range config.WhiteUrl {
		match, err := regexp.Match(k,[]byte(url))
		if err != nil {
			return false
		}
		if match {
			return match
		}
	}
	return false
}

type Auth struct {
	Enforce    *casbin.Enforcer
	AuthConfig *Config
}

func NewAuth(c core.AppConfig) *Auth {
	if authInstance !=nil{
		return authInstance
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", c.Db.Username, c.Db.Password, c.Db.Host)
	a, _ := gormadapter.NewAdapter(c.Db.Driver, dsn, c.Db.DataBase) // Your driver and data source.
	enforce, err := casbin.NewEnforcer(c.Casbin, a)
	if err != nil {
		log.Panicln("casbin err", err.Error())
		return nil
	}
	auth := &Auth{
		Enforce: enforce,
	}
	err = auth.Enforce.LoadPolicy()
	if err != nil {
		log.Panicln("casbin err", err.Error())
		return nil
	}
	authInstance = auth
	return auth
}

func GetAuthInstance() *Auth {
	return authInstance
}
func (auth *Auth) CheckPremiss(ctx *gin.Context) bool {
	token := ctx.GetHeader(Header)
	if len(strings.TrimSpace(token)) == 0 {
		refuse(ctx, 40001, "请登录")
		return false
	}
	if !strings.HasPrefix(token, Prefix) {
		refuse(ctx, 40001, "请登录")
		return false
	}
	subject := service.GetTokenService().GetSubject(strings.TrimPrefix(token, Prefix))
	if len(subject) == 0 {
		refuse(ctx, 40001, "请登录")
		return false
	}
	log.Println(ctx.Request.RequestURI, ctx.Request.Method)
	enforce, err := auth.Enforce.Enforce(subject, ctx.Request.RequestURI, ctx.Request.Method)
	if err != nil {
		refuse(ctx, 40001, "请登录")
		return false
	}
	if !enforce {
		refuse(ctx, 40001, "权限不足")
		return false
	}
	return true
}

func refuse(ctx *gin.Context, code int, msg string) {
	ctx.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
		"code": code,
		"err":  msg,
	})
}
