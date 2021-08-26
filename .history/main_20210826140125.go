/*
 * @Author: dsz
 * @Date: 2021-08-19 14:30:00
 * @LastEditTime: 2021-08-26 14:01:25
 * @Description:
 * @FilePath: \BlogApi\main.go
 */
package main

import (
	"fmt"

	"github.com/fahaxike/BlogApi/pkg"
	"github.com/fahaxike/BlogApi/routers"
)

func main() {
	router := routers.Initrouter()
	router.Run(fmt.Sprintf(":%d", pkg.HttpPort))
}
