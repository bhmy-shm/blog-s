package login

import (
	"github.com/gin-gonic/gin"
	"goweb/global"
	"goweb/internal/service"
	"goweb/pkg/MyJWT"
	"goweb/pkg/app"
	"goweb/pkg/errcode"
	"log"
)

type login struct{}

func Newlogin() login{
	return login{}
}

//创建密文表
func (this login) CreateJwt(c *gin.Context){
	response := app.NewResponse(c)

	svc := service.New(c.Request.Context())
	err := svc.CreatJwt()
	if err != nil {
		log.Println(err)
	}else{
		response.ToResponse("创建jwt密文表成功")
	}
}

//@Summary 申请获取token编码
//@Produce json
//@Param user_phone body string true "用户账户" maxlength(11)
//@Param user_pass  body string true "用户密码" minlength(6)
//@Param app_key  body string true "token密钥键"
//@Param app_secret body string true "token密钥值"
//@Success 200 {object} JWT.AuthJwt "申请token成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部错误"
//@Router /login/getToken [POST]
func (this login) GetUserJwt(c *gin.Context){

	response := app.NewResponse(c)
	request := service.LoginRequest{}

	err := c.Bind(&request)
	if err != nil{
		log.Println(err)
	}

	svc := service.New(c.Request.Context())
	err = svc.GetAuthJwt(&request)
	if err != nil {
		c.JSON(403,gin.H{"error":err})
		return
	}

	//获取token编码,并返回
	token,err := MyJWT.GenerateToken(request.AppKey,request.AppSecret,request.UserPhone)
	if err != nil {
		log.Println("bb",err)
	}
	//将token入redis
	err = svc.InputJwtRedis(token,request.UserPhone)
	if err != nil {
		log.Fatal("jwt写入Redis失败",err)
	}
	response.ToResponse(token)
}



//@Summary 用户携带token编码进行登录
//@Produce json
//@Param user_phone body string true "用户账户" maxlength(11)
//@Param user_pass  body string true "用户密码" minlength(6)
//@Param token header string true "token编码"
//@Success 200 {object} UserModel.User "登录成功"
//@Failure 400 {object} errcode.Error "请求错误"
//@Failure 500 {object} errcode.Error "内部错误"
//@Router /login/users [POST]
func (this login) UserLogin (c *gin.Context ) {
	response := app.NewResponse(c)

	UserPhone := c.PostForm("user_phone")
	UserPass := c.PostForm("user_pass")

	request := service.LoginRequest{UserPhone: UserPhone,UserPass: UserPass}


	svc := service.New(c.Request.Context())
	err := svc.LoginUser(&request)
	if err != nil {
		global.Logger.ErrorF(c,"svc LoginUser err:%s",err)
		response.ToErrorResponse(errcode.ErrorUserLogin.WithDetails(err.Error()))
	}else {
		//记录Cookie
		user := c.Request.Form.Get("user_phone") //phone
		pass := c.Request.Form.Get("user_pass") //token
		c.SetCookie(user,pass,7200,"/","localhost",false,true)

		response.ToResponse("登陆成功")
	}
}
