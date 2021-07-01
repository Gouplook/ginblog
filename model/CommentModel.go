/**
 * @Author: yinjinlin
 * @File:  CommentModel
 * @Description:
 * @Date: 2021/7/1 下午3:57
 */

package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserId       uint   `json:"user_id"`
	ArticleId    uint   `json:"article_id"`
	ArticleTitle string `json:"article_title"`
	Username     string `json:"username"`
	Content      string `gorm:"type:varchar(500);not null;" json:"content"`
	Status       int8   `gorm:"type:tinyint;default:2" json:"status"`
	IsDel        int   `gorm:"type:tinyint;default:0;comment:0-未删除，1-删除" json:"is_del`
}

// AddComment 新增评论
func AddComment(data *Comment) int {
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// GetComment 查询单个评论
func GetComment(id int) (Comment, int) {
	var comment Comment
	err = db.Where("id= ?", id).First(&comment).Error
	if err != nil {
		return comment, errmsg.ERROR
	}
	return comment, errmsg.SUCCSE
}

// GetCommentList 后台获取评论列表
func GetCommentList(pageSize, pageNum int) ([]Comment, int64, int) {
	var commentList []Comment
	var total int64
	db.Find(&commentList).Count(&total)
	err = db.Model(&commentList).Limit(pageSize).Offset((pageNum - 1) * pageSize).Select("*").Error
	if err != nil {
		return commentList, 0, errmsg.ERROR
	}
	return commentList, total, errmsg.SUCCSE
}

// GetCommentCount 获取评论数量
func GetCommentCount(id int) int64 {
	var comment Comment
	var total int64
	db.Find(&comment).Where("id = ?", id).Where("status = ?", 1).Count(&total)
	return total
}

// DeleteComment 删除评论
func DeleteComment(id uint,data *Comment) int {
	var maps = make(map[string]interface{})
	maps["is_del"] = data.IsDel
	err = db.Where("id = ?", id).Model(&Comment{}).Updates(maps).Error
	if err != nil {
		return  errmsg.ERROR
	}
	return errmsg.SUCCSE
}
