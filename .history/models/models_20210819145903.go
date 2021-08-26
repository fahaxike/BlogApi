/*
 * @Author: dsz
 * @Date: 2021-08-19 14:14:44
 * @LastEditTime: 2021-08-19 14:59:03
 * @Description:
 * @FilePath: \BlogApi\models\models.go
 */
package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func init() {
	var err error
	gorm.Open()

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", host, port, user, password, dbname)
	common.Db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln(err.Error())
	}

	ctx := context.Background()
	err = common.Db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Connected!")
}
