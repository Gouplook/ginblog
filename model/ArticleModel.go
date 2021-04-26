/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  ArticleModel.go
 * @Version: 1.0.0
 * @Date: 2021/4/1 7:46
 */
package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"` // 添加外键
	gorm.Model
	Title   string `gorm:"type:varchar(100)"json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Decs    string `gorm:"type:varchar(200)" json:"decs"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

//新增文章
func CreateArt(data *Article) (code int) {
	// 插入文章
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 获取文章列表
// pageNum 当前页数
// pageSize 页的条数
func GetArt(pageSize int, pageNum int) []Article {
	var atrs []Article
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&atrs).Error
	if err != nil && gorm.ErrRecordNotFound != nil {
		return nil
	}
	return atrs
}

//编辑文章
func EditArt(id int, data *Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["decs"] = data.Decs
	maps["content"] = data.Content
	maps["img"] = data.Img

	err = db.Where("id=?", id).Model(&Article{}).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

//删除文章
func DeleteArt(id int) int {
	err = db.Where("id=?", id).Delete(&Article{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
