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
	Title        string `gorm:"type:varchar(100)" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Decs         string `gorm:"type:varchar(200)" json:"decs"`
	Content      string `gorm:"type:longtext" json:"content"`
	ReadNum      int    `gorm:"type:int;not null; default:0" json:"read_num"`      // 阅读量
	CommentCount int    `gorm:"type:int;not null; default:0" json:"comment_count"` //评论量
	Img          string `gorm:"type:varchar(100)" json:"img"`
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

// 获取文章详情
func GetArtInfo(id int) (Article, int) {
	var art Article
	// 先进行查询，存在更新阅读量
	err = db.Where("id = ?", id).Preload("Category").First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}

	// 获取详情是，更新阅读量
	db.Model(&Article{}).Where("id = ?", id).UpdateColumn("read_num", gorm.Expr("read_num + ?", 1))

	return art, errmsg.SUCCSE
}

// 获取文章列表
// pageNum 当前页数
// pageSize 页的条数
func GetArtList(pageSize int, pageNum int) ([]Article, int64) {
	var artLists []Article
	var total int64
	err = db.Select("article.id, title, img, created_at, updated_at, decs , comment_count, read_num, category.name").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").
		Joins("Category").Find(&artLists).Error
	if err != nil && gorm.ErrRecordNotFound != nil {
		return nil, 0
	}
	// 单独进行计数
	db.Table("Article").Count(&total)

	return artLists, total
}

// 根据文章分类类型查找文章 GetArtByCidLists
func GetArtByCidLists(cid int, pageSize int, pageNum int) ([]Article, int, int64) {
	var artlists []Article
	var total int64
	//artlists := make([]Article, 0)
	// 偏移量计算
	err = db.Select("*").Where("cid = ? ", cid).
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Joins("Category").
		Find(&artlists).Error
	if err != nil && gorm.ErrRecordNotFound != nil {
		return nil, cid, 0
	}
	// 统计数量
	db.Model(&artlists).Where("cid = ?", cid).Count(&total)

	return artlists, errmsg.SUCCSE, total
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
