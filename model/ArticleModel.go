/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  Article.go
 * @Version: 1.0.0
 * @Date: 2021/4/1 7:46
 */
package model

import (
	"gorm.io/gorm"
)

type Article struct {
	Category Category
	gorm.Model
	Title    string `gorm:"type:varchar(100)"json:"title"`
	Cid      int `gorm:"type:int;not null" json:"cid"`
	Decs     string `gorm:"type:varchar(200)" json:"decs"`
	Content  string `gorm:"type:longtext" json:"content"`
	Img      string `gorm:"type:varchar(100)" json:"img"`
}
