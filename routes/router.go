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

		// category


	}

	r.Run(utils.HttpPort)
}


