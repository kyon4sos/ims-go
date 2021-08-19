package middle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError() gin.HandlerFunc {
	return gin.CustomRecovery(customHandle)
}
func customHandle(ctx *gin.Context, err interface{}) {
	ctx.AbortWithStatusJSON(http.StatusOK, map[string]interface{}{
		"code": 50000,
		"err":  "系统异常22233",
	})
}
