/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  User.go
 * @Version: 1.0.0
 * @Date: 2021/4/1 7:46
 */
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20)" json:"username"`
	Password string `gorm:"type:varchar(20)" json:"password"`
	Role     int    `gorm:"type:int" json:"role"` 			// 0= 管理员
}
