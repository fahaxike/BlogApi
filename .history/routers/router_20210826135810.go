/*
 * @Author: dsz
 * @Date: 2021-08-26 13:54:25
 * @LastEditTime: 2021-08-26 13:58:10
 * @Description:
 * @FilePath: \BlogApi\routers\router.go
 */
package routers

import "github.com/gin-gonic/gin"

func Initrouter() *gin.Engine {
	r := gin.New()
	r.Use()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "你好"})
	})
	gin.SetMode(gin.DebugMode)
	return r
}
