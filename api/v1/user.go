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
	"strconv"
)

// 查询用户是否存在
func UserExist(c *gin.Context) {

}

// 添加用户
func AddUser(c *gin.Context){
	// todo
	var data model.User
	_ = c.ShouldBindJSON(&data)  // json 格式绑定
	// 添加用户之前，需要查找用户是否存在
	code :=  model.CheckUser(data.Username)
	// 用户已经存在
	if code == errmsg.SUCCESS {
		// 插入数据
		model.CreateUser(&data)
	}
	if code == errmsg.ERR_USERNAME_USER{
		code = errmsg.ERR_USERNAME_USER
	}

	c.JSON(http.StatusOK,gin.H{
		"status" :code,
		"data" :data,
		"message":errmsg.GetErrMsg(code),
	})
}
// 查询单个用户

//type Page struct {
//	pageSize int `json:"page_size"`
//	pageNum  int `json:"page_num"`
//}


// 查询用户列表
func GetUsers(c *gin.Context){
	// 列表涉及到分页
	pageSize,_ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize= -1
	}
	if pageNum == 0{
		pageNum =-1
	}

	data := model.GetUsers(pageSize,pageNum)

	// 此类型不行
	//var pag Page
	////_ = c.ShouldBindJSON(&pag)
	//data := model.GetUsers(pag.pageSize,pag.pageNum)

	c.JSON(http.StatusOK,gin.H{
		"status": errmsg.SUCCESS,
		"data":data,
		"message":errmsg.GetErrMsg(errmsg.SUCCESS),
	})

}
// 编辑用户
func EditUser(c *gin.Context){

}
// 删除用户
func DeleteUser(c *gin.Context){
	id ,_ := strconv.Atoi(c.Param("id"))
	id = model.DeleteUser(id)
	if id < 0{
		return
	}
}


