package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

//将validator校验时发生的错误，封装到结构体中
type ValidError struct{
	Key string
	Message string
}

type ValidaErrors []*ValidError

//返回错误信息
func(v *ValidError) Error() string{
	return v.Message
}

//返回多个错误信息的组合
func(v ValidaErrors) Error() string{
	return strings.Join(v.Errors(),",")
}

//以切片的方式返回多个错误信息的组合
func(v ValidaErrors) Errors() []string{
	var errs []string
	for _,err := range v {
		errs = append(errs,err.Error())
	}
	return errs
}


//校验错误判断函数
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidaErrors) {
	var errs ValidaErrors

	err := c.ShouldBind(v)

	//如果传递的参数有校验错误，就记录到validator封装的错误结构体中
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		verrs, ok := err.(validator.ValidationErrors)

		//如果ok == false 代表没有错误
		//如果ok == true 代表有错误
		if ok == true {
			for key, value := range verrs.Translate(trans) {
				errs = append(errs, &ValidError{
					Key:     key,
					Message: value,
				})
			}
			return ok,errs
		}else{
			return false,nil
		}
	}
	return false, nil
}
