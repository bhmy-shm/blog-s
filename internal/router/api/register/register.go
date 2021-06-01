package register

import (
	"github.com/gin-gonic/gin"
	"goweb/global"
	"goweb/internal/service"
	"goweb/pkg/app"
	"goweb/pkg/errcode"
	"log"
)

type register struct{}

func Newregister() register{
	return register{}
}

//创建用户表
func (this register) Post(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c.Request.Context())

	err := svc.CreatUser()
	if err != nil {
		log.Println(err)
	}else{
		response.ToResponse("创建用户表成功")
	}
}

//@Summary 注册新的用户
//@Produce json
//@Param user_phone body string true "用户名" maxlength(11)
//@Param user_pass  body string true "密码" minlength(6)
//@Param user_email body string true "邮箱号"
//@Parm state body int true "状态码"
// @Success 200 {object} UserModel.User "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /register/auth [POST]
func (this register) Register(c *gin.Context){
	param := service.InsertRequest{}
	response := app.NewResponse(c)

	valid,err := app.BindAndValid(c,&param)
	if valid == true{
		global.Logger.ErrorF(c,"app.BindAndValid errs :%v",err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	//调用service
	svc := service.New(c.Request.Context())
	errs := svc.InsertUser(&param)
	if errs == errcode.ErrorUserRegisterRepetition{
		global.Logger.InfoF(c,"svc.InsertUser errs: %v",errs)
		response.ToErrorResponse(errcode.ErrorUserRegisterRepetition.WithDetails("重新注册"))
		return

	}else if errs == errcode.ErrorUserRedirect {
		global.Logger.InfoF(c,"Svc InsertUser errs: %v",errs)
		response.ToErrorResponse(errcode.ErrorUserRedirect.WithDetails("您已注册成功，请输入验证码"))
		return

	}else{
		response.ToRedirect(c)
		return
	}
}

//@Summary 确定手机验证码
//@Produce json
//@Param auth body string true "手机验证码" maxlength(6)
// @Success 200 {object} service.AuthRequest "验证码输入正确"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /register/authentication [POST]
func (this register) Authentication(c *gin.Context){
	param := service.AuthRequest{}
	response := app.NewResponse(c)

	valid,err := app.BindAndValid(c,&param)
	if valid == true{
		global.Logger.ErrorF(c,"app.BindAndValid errs :%v",err)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	errs := svc.Authentication(&param)
	if errs != nil {
		global.Logger.ErrorF(c,"Svc Authentication errs: %v",errs)
		response.ToErrorResponse(errcode.ErrorUserRegisterAuth.WithDetails(errs.Error()))
		return
	}else{
		response.ToResponse("用户验证码输入正确，注册成功")
		return
	}
}
