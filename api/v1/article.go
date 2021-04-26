/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  article.go
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

// 查询单个文章
func GetArtInfo(c *gin.Context) {
	var data model.Article
	var code int
	id,_ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	data ,code = model.GetArtInfo(id)
	if code == errmsg.SUCCSE {
		code = errmsg.SUCCSE
	}
	c.JSON(http.StatusOK,gin.H{
		"status" :code,
		"data": data,
		"message":errmsg.GetErrMsg(code),
	})

}


// 添加文章
func AddArt(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data) // json 格式绑定
	
	// 插入数据
	code := model.CreateArt(&data)

	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 查询文章列表
func GetArtslist(c *gin.Context) {
	// 列表涉及到分页
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data,total := model.GetArtList(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"total" :total,
		"status":  errmsg.SUCCSE,
		"data":    data,
		"message": errmsg.GetErrMsg(errmsg.SUCCSE),
	})

}

// 编辑文章
func EditArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var data model.Article
	_ = c.ShouldBindJSON(&data)
	//检查输入的文章是否存在
	code := model.EditArt(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除文章
func DeleteArt(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteArt(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
