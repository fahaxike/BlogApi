/*
 * @Author: dsz
 * @Date: 2021-08-15 21:41:36
 * @LastEditTime: 2021-08-15 22:02:35
 * @Description:分页处理
 * @FilePath: \BlogApi\pkg\util\pagination.go
 */
package util

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"

	"/pkgpkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result

}
