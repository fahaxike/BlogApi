/*
 * @Author: dsz
 * @Date: 2021-08-19 14:30:00
 * @LastEditTime: 2021-08-26 13:59:52
 * @Description:
 * @FilePath: \BlogApi\main.go
 */
package main

import (
	"strconv"

	"github.com/fahaxike/BlogApi/pkg"
	"github.com/fahaxike/BlogApi/routers"
)

func main() {
	router := routers.Initrouter()

	router.Run(":" + strconv.Itoa(pkg.HttpPort))
}
