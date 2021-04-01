/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  user.go
 * @Version: 1.0.0
 * @Date: 2021/4/2 7:01
 */
package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 查询用户是否存在
func UserExist(c *gin.Context) {

}

// 添加用户
func AddUser(c *gin.Context){
	// todo
	var data model.User
	_ = c.ShouldBindJSON(&data)
	// 添加用户之前，需要查找用户是否存在
	code :=  model.CheckUser(data.Username)
	// 用户已经存在
	if code == errmsg.ERR_USERNAME_USER {
		code = errmsg.ERR_USERNAME_USER
	}
	c.JSON(http.StatusOK,gin.H{
		"status" :code,
		"data" :data,
		"message":errmsg.GetErrMsg(code),
	})
}
// 查询单个用户

// 查询用户列表
func GetUsers(c *gin.Context){

}
// 编辑用户
func EditUser(c *gin.Context){

}
// 删除用户
func DeleteUser(c *gin.Context){

}


