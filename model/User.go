/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  User.go
 * @Version: 1.0.0
 * @Date: 2021/4/1 7:46
 */
package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20)" json:"username"`
	Password string `gorm:"type:varchar(20)" json:"password"`
	Role     int    `gorm:"type:int" json:"role"` 			// 0= 管理员
}

//新增用户
func CreateUser(data *User) (code int) {
	// 插入用户
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询用户是否存在
func CheckUser(username string) int {
	var users User
	db.Select("id").Where("username=?",username).First(&users)
	// 用户名已经存在
	if users.ID > 0 {
		return errmsg.ERR_USERNAME_USER
	}
	return  errmsg.SUCCESS
}
// pageNum 当前页数
// pageSize 页的条数
func GetUsers(pageSize int, pageNum int )[]User{
	var users []User
	err := db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&users).Error
	if err != nil && gorm.ErrRecordNotFound != nil {
		return nil
	}
	return users
}
//
//删除用户
