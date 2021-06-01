package app

import (
	"github.com/gin-gonic/gin"
	"goweb/pkg/errcode"
	"net/http"
)

type Response struct {
	Ctx *gin.Context
	Engine *gin.Engine
}


func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

//200反馈页面
func (r *Response) ToResponse(data interface{}) {
	r.Ctx.JSON(http.StatusOK, gin.H{"successful":data})
}

//301跳转页面
func (r *Response) ToRedirect(ctx *gin.Context){
	r.Ctx.Request.URL.Path = "/authentication"
	r.Engine.HandleContext(ctx)

}

//403错误信息
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{"code":err.Code(),"msg":err.Msg()}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(),response)
}


//写入cookie参数
func (r *Response) ToCookieResponse(ctx *gin.Context) {
	//第一个参数为 cookie 名（phone）；第二个参数为 cookie 值（token）；
	//第三个参数为 cookie 有效时长，单位秒，当 cookie 存在的时间超过设定时间时，cookie 就会失效，它就不再是我们有效的 cookie；
	//第四个参数为 cookie 所在的目录；
	//第五个为所在域，表示我们的 cookie 作用范围；
	//第六个表示是否只能通过 https 访问；第七个表示 cookie 是否可以通过 js代码进行操作
}


func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

