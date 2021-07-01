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
	// 两种不同的请求，对应的路由也不一样
	// a GET request to /user/5
	//id, _ := strconv.Atoi(c.Param("id"))
	// id, _ := strconv.Atoi(c.Query("id"))
	id, _ := strconv.Atoi(c.PostForm("id"))
	data, code = model.GetArtInfo(id)
	if code == errmsg.SUCCSE {
		code = errmsg.SUCCSE
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
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

	data, total := model.GetArtList(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"total":   total,
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

// 根据文章分类类型查找文章 GetArtByCidLists 暂时不用jion查询查询
func GetArtByCidLists(c *gin.Context) {
	// postman 传Json
	//cid, _ := strconv.Atoi(c.Param("cid"))
	//pageSize, _ := strconv.Atoi(c.Param("pageSize"))
	//pageNum, _ := strconv.Atoi(c.Param("pageNum"))

	//
	cid, _ := strconv.Atoi(c.Query("cid"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	data, code, total := model.GetArtByCidLists(cid, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"total":   total,
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除文章
func DeleteArt(c *gin.Context) {

	var data model.Article
	id, _ := strconv.Atoi(c.PostForm("id"))
	data.IsDel,_ = strconv.Atoi(c.PostForm("isDel"))

	code := model.DeleteArt(id,&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
