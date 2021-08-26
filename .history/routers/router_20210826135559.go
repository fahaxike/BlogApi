/*
 * @Author: dsz
 * @Date: 2021-08-26 13:54:25
 * @LastEditTime: 2021-08-26 13:55:59
 * @Description:
 * @FilePath: \BlogApi\routers\router.go
 */
package routers

import "github.com/gin-gonic/gin"

func init() *gin.eEngine {
	r := gin.Default()
	gin.Mode(gin.DebugMode)
	return r
}
