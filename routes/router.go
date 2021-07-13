/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  router.go
 * @Version: 1.0.0
 * @Date: 2021/4/1 7:39
 */
package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 如何剥离出func
func InitRouter() {
	gin.SetMode(utils.AppMode)
	// r := gin.Default() // 默认加两个中间件

	// 创建一个不包含中间件的路由器
	r := gin.New()

	// 跨域问题
	r.Use(middleware.Cors())
	// 日志
	r.Use(middleware.Log())


	// 后台管理路由接口
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{

		// 用户模块 v1 是API中v1文件的包名
		auth.POST("user/add", v1.AddUser)
		auth.POST("uses", v1.GetUsers)
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("dele/:id", v1.DeleteUser)

		// 文章模块 atricle
		auth.POST("art/add", v1.AddArt)
		auth.POST("art/getList", v1.GetArtslist)            // 获取文章列表
		auth.POST("art/getInfo", v1.GetArtInfo)             // 获取文章详情 Query 请求
		auth.PUT("art/edit/:id", v1.EditArt)                // 编辑文章
		auth.POST("art/delete", v1.DeleteArt)               // 删除文章
		auth.GET("art/getArtCidLists", v1.GetArtByCidLists) // 根据分类Cid 获取文章详情

		// 文章分类模块
		auth.POST("cate/add", v1.AddCate)
		auth.POST("cate", v1.GetList)
		auth.GET("cateinfo/:id", v1.GetInfo) // 获取分类详情
		auth.PUT("cate/edit/:id", v1.EditCate)
		auth.POST("cate/dele/:id", v1.DeleteCate)

		// 评论模块
		auth.POST("comment/add",v1.AddComment)
		auth.POST("comment/getComment",v1.GetComment)
		auth.POST("comment/getCommentLists",v1.GetCommentList)
		auth.POST("comment/deleteComment",v1.DeleteComment)

		// 文件上传
		auth.POST("upload",v1.UpLoad)

		//修改密码
		auth.POST("admin/changepw/:id", v1.ChangeUserPassword)

		// 更新个人设置
		auth.GET("admim/profile/:id",v1.Getprofile)
		auth.POST("admin/profile/:id", v1.UpdateProfile)
	}

	// 前端展示页面接口
	router := r.Group("api/v2")
	{
		// 路由组(测试）
		router.GET("Hello", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "测试路由.....",
			})
		})

		// 登录控制模块
		router.POST("login", v1.Login)           // 后台
		router.POST("loginfront", v1.LoginFront) // 前台

		// 用户模块 v1 是API中v1文件的包名
		router.POST("user/add", v1.AddUser)
		router.POST("uses", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("dele/:id", v1.DeleteUser)

		// 文章模块 atricle
		router.POST("art/add", v1.AddArt)
		router.POST("art/getList", v1.GetArtslist)            // 获取文章列表
		router.POST("art/getInfo", v1.GetArtInfo)             // 获取文章详情 Query 请求
		router.PUT("art/edit/:id", v1.EditArt)                // 编辑文章
		router.POST("art/delete", v1.DeleteArt)               // 删除文章
		router.GET("art/getArtCidLists", v1.GetArtByCidLists) // 根据分类Cid 获取文章详情

		// 文章分类模块
		router.POST("cate/add", v1.AddCate)
		router.POST("cate", v1.GetList)
		router.GET("cateinfo/:id", v1.GetInfo) // 获取分类详情
		router.PUT("cate/edit/:id", v1.EditCate)
		router.POST("cate/dele/:id", v1.DeleteCate)

		// 评论模块
		router.POST("comment/add",v1.AddComment)
		router.POST("comment/getComment",v1.GetComment)
		router.POST("comment/getCommentLists",v1.GetCommentList)
		router.POST("comment/deleteComment",v1.DeleteComment)
		// 文件上传
		router.POST("upload",v1.UpLoad)
		//修改密码
		router.POST("admin/changepw/:id", v1.ChangeUserPassword)
	}

	_ = r.Run(utils.HttpPort)
}
