/**
 * @Author: yinjinlin
 * @File:  validator
 * @Description:
 * @Date: 2021/7/6 上午9:14
 */

package validator

import (
	"fmt"
	"ginblog/utils/errmsg"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

// 参数校验
func Validate(data interface{})(string, int) {
	validate := validator.New()
	uni := unTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	// 对验证参数进行翻译
	err := zh.RegisterDefaultTranslations(validate, trans)

	if err != nil {
		// 输出到日志中...
		fmt.Println("err: ",err)
	}
	//
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})

	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}

	return "", errmsg.SUCCSE
}