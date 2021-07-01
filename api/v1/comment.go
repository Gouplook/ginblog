/**
 * @Author: yinjinlin
 * @File:  comment
 * @Description:
 * @Date: 2021/7/1 下午4:19
 */

package v1

import (
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加评论
func AddComment(c *gin.Context) {
	var data model.Comment
	_ = c.ShouldBindJSON(&data) // json 格式绑定
	// 插入数据
	code := model.AddComment(&data)

	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询评论
func GetComment(c *gin.Context) {
	var data model.Comment
	var code int

	id, _ := strconv.Atoi(c.PostForm("id"))
	data, code = model.GetComment(id)
	if code == errmsg.SUCCSE {
		code = errmsg.SUCCSE
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})

}

// 查询评论列表
func GetCommentList(c *gin.Context) {
	// 列表涉及到分页
	pageSize, _ := strconv.Atoi(c.PostForm("pageSize"))
	pageNum, _ := strconv.Atoi(c.PostForm("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, total, code := model.GetCommentList(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"total":   total,
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(errmsg.SUCCSE),
	})

}

// 删除评论
func DeleteComment(c *gin.Context) {
	var data model.Comment
	id, _ := strconv.Atoi(c.PostForm("id"))
	data.IsDel, _ = strconv.Atoi(c.PostForm("isDel"))

	code := model.DeleteComment(uint(id), &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
