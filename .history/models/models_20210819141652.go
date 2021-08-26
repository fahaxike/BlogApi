/*
 * @Author: dsz
 * @Date: 2021-08-19 14:14:44
 * @LastEditTime: 2021-08-19 14:16:51
 * @Description:
 * @FilePath: \BlogApi\models\models.go
 */
package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
