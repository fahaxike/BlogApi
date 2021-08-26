/*
 * @Author: dsz
 * @Date: 2021-08-26 13:54:25
 * @LastEditTime: 2021-08-26 13:56:52
 * @Description:
 * @FilePath: \BlogApi\routers\router.go
 */
package routers

import "github.com/gin-gonic/gin"

func init() *gin.Engine {
	r := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "你好"})
	})
	gin.Mode(gin.DebugMode)
	return r
}
