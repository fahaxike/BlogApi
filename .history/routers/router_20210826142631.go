/*
 * @Author: dsz
 * @Date: 2021-08-26 13:54:25
 * @LastEditTime: 2021-08-26 14:26:31
 * @Description:
 * @FilePath: \BlogApi\routers\router.go
 */
package routers

import (
	v1 "github.com/fahaxike/BlogApi/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func Initrouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "你好"})
	})
	apiv1 := r.Group("/api/v1")
	apiv1.GET("/tags", v1.GetTags)
	gin.SetMode(gin.DebugMode)
	return r
}
