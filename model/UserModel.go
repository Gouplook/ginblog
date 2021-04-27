/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  UserModel.go
 * @Version: 1.0.0
 * @Date: 2021/4/1 7:46
 */
package model

import (
	"encoding/base64"
	"ginblog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20)" json:"username"`
	Password string `gorm:"type:varchar(64)" json:"password"`
	Role     int    `gorm:"type:int" json:"role"` 			// 0= 管理员
}


//查询用户是否存在
func CheckUser(username string) int {
	var users User
	db.Select("id").Where("username=?",username).First(&users)
	// 用户名已经存在
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return  errmsg.SUCCSE
}

//新增用户
func CreateUser(data *User) (code int) {
	// 函数密文存储
	data.Password = ScryptPassWord(data.Password)
	// 插入用户
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 获取用户列表
// pageNum 当前页数
// pageSize 页的条数
func GetUsers(pageSize int, pageNum int )[]User{
	var users []User
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&users).Error
	if err != nil && gorm.ErrRecordNotFound != nil {
		return nil
	}
	return users
}
//编辑用户
func EditUsers( id int, data *User) int {
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role

	err = db.Where("id=?", id).Model(&User{}).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return  errmsg.SUCCSE
}

//删除用户
func DeleteUser(id int) int {
	err = db.Where("id=?",id).Delete(&User{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return  errmsg.SUCCSE
}


// 钩子函数 密文存储密码
//func (u *User)BeforeSave(){
//	u.Password = ScryptPassWord(u.Password)
//}
//
func ScryptPassWord(passWord string) string{
	// 加密的盐
	salt := []byte{1,2,3,4,5,6,7,8}
	dk,err := scrypt.Key([]byte(passWord),salt,64,8,1,32)
	if err != nil {
		log.Fatal(err)
	}
	passWord = base64.StdEncoding.EncodeToString(dk)
	return passWord
}

