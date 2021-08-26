/*
 * @Author: dsz
 * @Date: 2021-08-15 16:36:28
 * @LastEditTime: 2021-08-15 16:57:30
 * @Description:获取初始化ini配置
 * @FilePath: \BlogApi\pkg\setting.go
 */
package setting

import (
	"time"

	"gopkg.in/ini.v1"
)

var (
	Cfg *ini.File

	RunModel string

	HttpPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration

	Page 
)

func init() {

}
