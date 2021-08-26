package routers

import "github.com/gin-gonic/gin"

func init() *gin.eEngine {
	r := gin.Default()
	gin.Mode(gin.DebugMode)
	return r
}
