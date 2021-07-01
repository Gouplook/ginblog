/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  category.go
 * @Version: 1.0.0
 * @Date: 2021/4/2 7:02
 */
package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)



// 添加分类
func AddCate(c *gin.Context) {
	var data model.Category
	_ = c.ShouldBindJSON(&data) // json 格式绑定

	code := model.CreateCate(&data)

	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个分类
func GetInfo(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	var data model.Category
	var code int
	_ = c.ShouldBindJSON(&data)
	data, code = model.FindInfo(id)
	c.JSON(http.StatusOK,gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(errmsg.SUCCSE),
	})
}

// 查询分类列表
func GetList(c *gin.Context) {
	// 列表涉及到分页
	// pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	// pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	pageSize,_ := strconv.Atoi(c.PostForm("pageSize"))
	pageNum,_ := strconv.Atoi(c.PostForm("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data := model.GetCateLists(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCSE,
		"data":    data,
		"message": errmsg.GetErrMsg(errmsg.SUCCSE),
	})
}

// 编辑分类
func EditCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Category
	_ = c.ShouldBindJSON(&data)
	//检查输入的分类是否存在
	code := model.EditCate(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除分类
func DeleteCate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Category
	data.IsDel,_ = strconv.Atoi(c.PostForm("isDel"))
	code := model.DeleteCate(id,&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
