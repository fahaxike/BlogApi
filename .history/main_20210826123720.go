/*
 * @Author: dsz
 * @Date: 2021-08-19 14:30:00
 * @LastEditTime: 2021-08-26 12:37:20
 * @Description:
 * @FilePath: \BlogApi\main.go
 */
package main

import (
	"strconv"

	"github.com/fahaxike/BlogApi/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "你好"})
	})
	router.
		router.Run(":" + strconv.Itoa(pkg.HttpPort))
}
