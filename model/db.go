/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  db.go
 * @Version: 1.0.0
 * @Date: 2021/4/1 7:47
 */
package model

import "gorm.io/gorm"

//
var db *gorm.DB
var err error

func InitDb (){
	// 链接数据库
	db,err := gorm.Open()
}
