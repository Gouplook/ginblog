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
	//defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Article{},&Category{},&User{})

	//// 创建
	//db.Create(&Product{Code: "L1212", Price: 1000})
	//
	//// 读取
	//var product Product
	//db.First(&product, 1) // 查询id为1的product
	//db.First(&product, "code = ?", "L1212") // 查询code为l1212的product
	//
	//// 更新 - 更新product的price为2000
	//db.Model(&product).Update("Price", 2000)
	//
	//// 删除 - 删除product
	//db.Delete(&product)



}
