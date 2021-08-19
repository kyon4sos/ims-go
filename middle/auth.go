package middle

import (
	"github.com/gin-gonic/gin"
	"img-server/auth"
	"img-server/core"
)

type ValideJwtFunc func(token string) bool
func JwtAuth(appConfig core.AppConfig,config auth.Config) gin.HandlerFunc {
	newAuth :=auth.NewAuth(appConfig)
	return func(ctx *gin.Context) {
		if !config.IsWhiteUrl(ctx) {
			newAuth.CheckPremiss(ctx)
		}
	}
}



