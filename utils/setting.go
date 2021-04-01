/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  setting.go
 * @Version: 1.0.0
 * @Date: 2021/4/1 7:51
 */
package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbUser     string
	DbHost     string
	DbPort     string
	DbName     string
	DbPassWord string
)

// 初始化配置文件
func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println(err)
	}

	LoadServer(file)
	LoadData(file)
}

// 读取配置文件
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":10086")

}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassword").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
	DbHost = file.Section("database").Key("DbHost").MustString("root")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
}
