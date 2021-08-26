/*
 * @Author: dsz
 * @Date: 2021-08-19 14:14:44
 * @LastEditTime: 2021-08-19 15:13:35
 * @Description:
 * @FilePath: \BlogApi\models\models.go
 */
package models

import (
	"fmt"
	"log"

	"github.com/fahaxike/BlogApi/pkg"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func init() {
	var err error
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", pkg.DBSource.Host, pkg.DBSource.Port, pkg.DBSource.User, pkg.DBSource.Password, pkg.DBSource.Dbname)
	DB, err = gorm.Open(pkg.DBSource.Type, connStr)

	if err != nil {
		log.Println(err.Error())
	}

	// ctx := context.Background()
	// err = DB.PingContext(ctx)
	// if err != nil {
	// 	log.Fatalln(err.Error())
	// }
	// log.Println("Connected!")
}
