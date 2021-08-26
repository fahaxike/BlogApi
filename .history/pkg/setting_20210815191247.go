/*
 * @Author: dsz
 * @Date: 2021-08-15 16:36:28
 * @LastEditTime: 2021-08-15 19:12:47
 * @Description:获取初始化ini配置
 * @FilePath: \BlogApi\pkg\setting.go
 */
package setting

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

var (
	Cfg *ini.File

	RunModel string

	HttpPort     int
	ReadTimeOut  time.Duration
	WriteTimeOut time.Duration

	PageSize  int
	JWTSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/api.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/api.ini':%v", err)
	}

}
func loadBase(){
	Cfg.GetSection().
}
