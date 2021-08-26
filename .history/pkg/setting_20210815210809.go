/*
 * @Author: dsz
 * @Date: 2021-08-15 16:36:28
 * @LastEditTime: 2021-08-15 21:06:10
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
func loadBase() {

	RunModel = Cfg.Section("").Key("RUN_MODEL").In("debug", []string{"debug", "release"})
}
func loadApp() {

	PageSize = Cfg.Section("app").Key("PAGE_SIZE").MustInt(10)
	JWTSecret = Cfg.Section("app").Key("JWT_SECRET").MustString("23347$040412")

}
func loadServer() {

	HttpPort = Cfg.Section("server").Key("HTTP_PORT").MustInt(10)
	ReadTimeOut = time.Duration(Cfg.Section("server").Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeOut = time.Duration(Cfg.Section("server").Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

}
func loadDB() {

	HttpPort = Cfg.Section("database").Key("TYPE").MustInt(10)
	ReadTimeOut = time.Duration(Cfg.Section("database").Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeOut = time.Duration(Cfg.Section("database").Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

}
