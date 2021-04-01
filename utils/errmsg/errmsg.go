/**
 * @Author: Yinjinlin
 * @Description:
 * @File:  errmsg.go
 * @Version: 1.0.0
 * @Date: 2021/4/2 6:51
 */
package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000 ...用户模块
	ERR_USERNAME_USER  = 1001
	ERR_PASSWORD       = 1002
	ERR_USER_NOT_EXIST = 1003
	ERR_TOKEN_EXIST    = 1004
	ERR_TOKEN_RUNTIME  = 1005
	ERR_TOKEN_WRONG    = 1006
	ERR_TOKEN_TYPE     = 1007
	// code = 2000 ...文章模块

)

// 字典
var codeMsg = map[int]string{
	SUCCESS:            "OK",
	ERROR:              "FAIL",
	ERR_USERNAME_USER:  "用户名已存在",
	ERR_PASSWORD:       "密码错误",
	ERR_USER_NOT_EXIST: "用户不存在",
	ERR_TOKEN_EXIST:    "TOKEN不存在",
	ERR_TOKEN_RUNTIME:  "TOKEN已过期",
	ERR_TOKEN_WRONG:    "TOKNE不正确",
	ERR_TOKEN_TYPE:     "TOKNE格式不正确",
}

func GetErrMsg(code int) string {

	return codeMsg[code]
}
