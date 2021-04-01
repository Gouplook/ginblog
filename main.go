/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  main.go
 * @Version: 1.0.0
 * @Date: 2021/4/1 19:38
 */
package main

import (
	"ginblog/model"
	"ginblog/routes"
)

func main(){

	// 连接数据库
	model.InitDb()

	// 启动路由
	routes.InitRouter()
}
