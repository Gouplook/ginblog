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
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);comment:用户名称" json:"username"`
	Password string `gorm:"type:varchar(64);comment:用户密码" json:"password"`
	Role     int    `gorm:"type:int;comment:用户角色" json:"role"`                 // 0= 管理员
	IsDel    int    `gorm:"type:tinyint;default:0;comment:是否删除" json:"is_del"` // 是否删除
}

// 查询用户是否存在
func CheckUser(username string) int {
	var users User
	db.Select("id").Where("username= ?", username).First(&users)
	// 用户名已经存在
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCSE
}

// 新增用户
func CreateUser(data *User) (code int) {
	// 函数密文存储
	data.Password = ScryptPassWord(data.Password)
	// 插入用户
	err = db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 获取用户列表
// pageNum  当前第几页数
// pageSize 每页显示的条数
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && gorm.ErrRecordNotFound != nil {
		return nil
	}
	return users
}

// 编辑用户
func EditUsers(id int, data *User) int {
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role

	err = db.Where("id= ?", id).Model(&User{}).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除用户
func DeleteUser(id int) int {
	err = db.Where("id=?", id).Delete(&User{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 钩子函数 密文存储密码
// func (u *User)BeforeSave(){
//	u.Password = ScryptPassWord(u.Password)
// }
//
func ScryptPassWord(passWord string) string {
	// 加密的盐
	salt := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	dk, err := scrypt.Key([]byte(passWord), salt, 64, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	passWord = base64.StdEncoding.EncodeToString(dk)
	return passWord
}

//  CheckLogin 后台登录验证
func CheckLogin(username string, password string) (User, int) {
	var user User
	var PasswordErr error

	db.Where("username= ?", username).First(&user)
	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return user, errmsg.ERROR_USER_NO_RIGHT
	}

	return user, errmsg.SUCCSE
}

// CheckLoginFront 前台登录
func CheckLoginFront(username string, password string) (User, int) {
	var user User
	var PasswordErr error
	db.Where("username= ?",username).First(&user)
	PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if PasswordErr != nil {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	return user, errmsg.SUCCSE
}
