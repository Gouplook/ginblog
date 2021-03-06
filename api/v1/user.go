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
	"ginblog/validator"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// 查询用户是否存在
func UserExist(c *gin.Context) {
	// 税号： 91340 10077 11142 5X3
	// 邮编：230051  电话：0551-62249972
	// 公司名称：安徽巨一科技股份有限公司
	// 公司地址：安徽省合肥市包河区繁华大道 5821号
	// 开户行：交通银行合肥北京路支行
	// 账户：34133 50000 18880 003426


}

// 添加用户
func AddUser(c *gin.Context){
	// todo
	var data model.User
	_ = c.ShouldBindJSON(&data)  // json 格式绑定

	// 用户参数校验
	msg, validCode :=  validator.Validate(&data)
	if validCode != errmsg.SUCCSE {
		c.JSON(http.StatusOK,gin.H{
				"status":validCode,
				"message":msg,
			})
		c.Abort()
		return
	}

	// 添加用户之前，需要查找用户是否存在
	code :=  model.CheckUser(data.Username)
	// 用户已经存在
	if code == errmsg.SUCCSE {
		// 插入数据
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED{
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK,gin.H{
		"status" :code,
		"data" :data,
		"message":errmsg.GetErrMsg(code),
	})

}

// 查询用户列表
func GetUsers(c *gin.Context){
	// 列表涉及到分页（URL方式请求 ：/path?id=1234&name=Manu&value= ）
	// pageSize,_ := strconv.Atoi(c.Query("pageSize"))
	// pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	pageSize,_  :=strconv.Atoi(c.PostForm("pageSize"))
	pageNum,_  :=strconv.Atoi(c.PostForm("pageNum"))

	if pageSize == 0 {
		pageSize= -1
	}
	if pageNum == 0{
		pageNum =-1
	}

	data := model.GetUsers(pageSize,pageNum)
	c.JSON(http.StatusOK,gin.H{
		"status": errmsg.SUCCSE,
		"data":data,
		"message":errmsg.GetErrMsg(errmsg.SUCCSE),
	})

}
// 编辑用户
func EditUser(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	var data model.User
	_ = c.ShouldBindJSON(&data)
	//检查输入的用户是否存在
	code := model.CheckUser(data.Username)
	if code == errmsg.ERROR_USERNAME_USED {
		log.Fatal("用户名已存在")
	}
	code = model.EditUsers(id, &data)

	c.JSON(http.StatusOK,gin.H{
		"status" :code,
		"data" :data,
		"message": errmsg.GetErrMsg(code),
	})
}
// 删除用户
func DeleteUser(c *gin.Context){
	id ,_ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)

	c.JSON(http.StatusOK,gin.H{
		"status" :code,
		"message" :errmsg.GetErrMsg(code),
	})
}

// 修改密码
func ChangeUserPassword(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.PostForm("id"))
	_ = c.ShouldBindJSON(&data)

	code := model.ChangePassword(id, &data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}



