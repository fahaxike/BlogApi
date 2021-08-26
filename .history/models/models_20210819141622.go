/*
 * @Author: dsz
 * @Date: 2021-08-19 14:14:44
 * @LastEditTime: 2021-08-19 14:15:42
 * @Description:
 * @FilePath: \BlogApi\models\models.go
 */
package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"gin-blog/pkg/setting"
)

var DB *gorm.DB
