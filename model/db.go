/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  db.go
 * @Version: 1.0.0
 * @Date: 2021/4/1 7:47
 */
package model

import (
	"fmt"
	"ginblog/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//
var db *gorm.DB
var err error

func InitDb (){
	// 链接数据库
	// "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	sql := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", utils.DbUser,utils.DbPassWord,utils.DbName)
	db, err = gorm.Open(utils.Db, sql)
	if err != nil {
		panic("failed to connect database")
	}




	// Migrate the schema 初始化数据库model
	db.AutoMigrate(&Article{},&Category{},&User{})

	sqlDB := db.DB()

	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)


}
