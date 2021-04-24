/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  CategoryModel.go
 * @Version: 1.0.0
 * @Date: 2021/4/1 7:47
 */
package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID uint `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type: varchar(20);not null" json:"name"`
}

//查询分类是否存在

// 查询单个分类下的文章

//新增分类
func CreateCate(data *Category) (code int) {
	// 插入分类
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 获取分类详情
func FindInfo(id int ) (Category, int ) {
	var cate Category
	err := db.Limit(1).Where("ID=?",id).Find(&cate).Error
	if err != nil {
		return cate,errmsg.ERROR
	}
	return cate,errmsg.SUCCESS
}

// 获取分类列表
// pageNum 当前页数
// pageSize 页的条数
func GetCateLists(pageSize int, pageNum int) []Category {
	var cateLists []Category
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cateLists).Error
	if err != nil && gorm.ErrRecordNotFound != nil {
		return nil
	}
	return cateLists
}

//编辑分类
func EditCate(id int, data *Category) int {
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err = db.Where("id=?", id).Model(&Category{}).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除分类
func DeleteCate(id int) int {
	err = db.Where("id=?", id).Delete(&Category{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

