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
	"ginblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func InitRouter() (engine *gin.Engine){
//	gin.SetMode(utils.AppMode)
//	r := gin.Default()  // 默认加两个中间件
//	router := r.Group("api/v1")
//	{
//		router.GET("Hello", func(context *gin.Context) {
//			context.JSON(http.StatusOK, gin.H{
//				"msg" :"ok",
//			})
//		})
//	}
//
//	return
//}

// 如何剥离出func
func InitRouter(){
	gin.SetMode(utils.AppMode)
	r := gin.Default()  // 默认加两个中间件

	router := r.Group("api/v1")
	{
		//路由组
		router.GET("Hello", func(context *gin.Context) {
			context.JSON(http.StatusOK,gin.H{
				"msg" :"ppppp",
			})
		})

		// user v1 是API中v1文件的包名
		router.POST("user/add", v1.AddUser)
		router.GET("uses", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("dele/:id",v1.DeleteUser)

		// atricle
		router.POST("art/add", v1.AddArt)
		router.GET("art/get", v1.GetArts)
		router.PUT("art/edit:id", v1.EditCate)
		router.DELETE("art/dele:id",v1.DeleteArt)

		// category
		router.POST("cate/add", v1.AddCate)
		router.GET("cate", v1.GetList)
		router.GET("cateinfo", v1.GetInfo)  // 获取分类详情
		router.PUT("cate/edit/:id", v1.EditCate)
		router.DELETE("cate/dele/:id",v1.DeleteCate)

	}

	r.Run(utils.HttpPort)
}


