/*
 * @Author: dsz
 * @Date: 2021-08-15 16:36:28
 * @LastEditTime: 2021-08-26 11:27:05
 * @Description:获取初始化ini配置
 * @FilePath: \BlogApi\pkg\setting.go
 */
package pkg

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)

type PQDb struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	Type     string
}

var (
	Cfg *ini.File

	RunModel     string
	DBSource     PQDb
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
	loadBase()
	loadApp()
	loadServer()
	loadDatabase()
}
func loadBase() {

	RunModel = Cfg.Section("").Key("RUN_MODEL").In("debug", []string{"debug", "release"})
}
func loadApp() {

	PageSize = Cfg.Section("app").Key("PAGE_SIZE").MustInt(10)
	JWTSecret = Cfg.Section("app").Key("JWT_SECRET").MustString("23347$040412")

}
func loadDatabase() {

	// 	TYPE=postgresql
	// USER=deng
	// PWD=123456
	// HOST=192.168.7.104:5432
	// NAME=mydb
	// TABLE_PREFIX=blog_

	DBSource.Host = Cfg.Section("database").Key("HOST").MustString("127.0.0.1")
	DBSource.Port = Cfg.Section("database").Key("PORT").MustInt(5432)
	DBSource.User = Cfg.Section("database").Key("USER").MustString("")
	DBSource.Password = Cfg.Section("database").Key("PWD").MustString("")
	DBSource.Type = Cfg.Section("database").Key("TYPE").MustString("")

}

func loadServer() {

	HttpPort = Cfg.Section("server").Key("HTTP_PORT").MustInt(10)
	ReadTimeOut = time.Duration(Cfg.Section("server").Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeOut = time.Duration(Cfg.Section("server").Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

}
