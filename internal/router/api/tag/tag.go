package tag

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goweb/global"
	"goweb/internal/service"
	"goweb/pkg/app"
	"goweb/pkg/errcode"
)

type tag struct {
}

func NewTag() tag {
	return tag{}
}

//创建标签表
func (t tag) CreateTable (c *gin.Context) {
	svc := service.New(c.Request.Context())
	err := svc.CreateTagTable()
	if err != nil {
		c.JSON(400,gin.H{"message":"创建标签表失败"})
	}else{
		c.JSON(200,gin.H{"message":"创建标签表成功"})
	}
}

//统计指定标签的个数
func (t tag) Count(c *gin.Context) {
	param := service.CountRequest{}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)
	fmt.Println("xx：",errs)

	//如果校验错误
	if valid == true {
		global.Logger.ErrorF(c,"app.BindAndValid errs :%v",errs)
		response.ToErrorResponse(errcode.ErrorCountTagField.WithDetails(errs.Errors()...))
	}

	svc := service.New(c.Request.Context())

	num,err := svc.CountTag(&param)
	fmt.Println("统计标签总数：",err)

	if err != nil{
		global.Logger.ErrorF(c,"svc.CountTag err:= %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagField)	//入库失败不需要写入日志
		return
	}
	response.ToResponse(num)
	return
}

//通过名字获取详细的标签信息
func (t tag) Get(c *gin.Context) {
	param := service.StateTagRequest{}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)

	//从路由中拿到
	parmName := c.Param("name")

	//如果校验错误
	if valid == true {
		global.Logger.ErrorF(c,"app.BindAndValid errs :%v",errs)
		response.ToErrorResponse(errcode.ErrorGetTagField.WithDetails(errs.Errors()...))
	}

	//如果参数校验没错，则开始入库
	svc := service.New(c.Request.Context())
	fmt.Println("Get param:",param)
	tag,err := svc.GetTag(&param,parmName)
	if err != nil{
		global.Logger.ErrorF(c,"svc.GetTagField err:= %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagField)	//入库失败不需要写入日志
		return
	}
	response.ToResponse(tag)
	return
}

func (t tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)

	//如果校验错误
	if valid == true {
		global.Logger.ErrorF(c,"app.BindAndValid errs :%v",errs)
		response.ToErrorResponse(errcode.ErrorCreateTagField.WithDetails(errs.Errors()...))
	}
	//获取分页
	svc := service.New(c)
	fmt.Println("page and size",app.GetPage(c),app.GetPageSize(c))
	pager := app.Pager{Page: app.GetPage(c),PageSize: app.GetPageSize(c)}

	tags,err := svc.GetTagList(&param,&pager)
	if err != nil{
		global.Logger.ErrorF(c,"svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListField)
		return
	}
	response.ToResponseList(tags,len(tags))
	return
}

//创建标签
func (t tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)

	//如果校验错误
	if valid == true {
		global.Logger.ErrorF(c,"app.BindAndValid errs :%v",errs)
		response.ToErrorResponse(errcode.ErrorCreateTagField.WithDetails(errs.Errors()...))
	}

	//如果参数校验没错，则开始入库
	svc := service.New(c.Request.Context())
	fmt.Println("create param:",&param)
	err := svc.CreateTag(&param)
	if err != nil{
		global.Logger.ErrorF(c,"svc.CreateTag err:= %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagField)	//入库失败不需要写入日志
		return
	}
	response.ToResponse("创建标签成功")
	return
}

//修改标签
func (t tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)

	//如果校验错误
	if valid == true {
		global.Logger.ErrorF(c,"UpdateTag,app.BindAndValid errs:%v",errs)
		response.ToErrorResponse(errcode.ErrorUpdateTagField.WithDetails(errs.Errors()...))
	}

	//入库
	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil{
		global.Logger.ErrorF(c,"svc.UpdateTag err:= %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagField)	//入库失败不需要写入日志
		return
	}
	response.ToResponse("修改标签成功")
	return
}

//删除标签
func (t tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)

	//如果校验错误
	if valid == true {
		global.Logger.ErrorF(c,"DeleteTag,app.BindAndValid errs:%v",errs)
		response.ToErrorResponse(errcode.ErrorDeleteTagField.WithDetails(errs.Errors()...))
	}

	//入库
	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil{
		global.Logger.ErrorF(c,"svc.deleteTag err:= %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagField)	//入库失败不需要写入日志
		return
	}
	response.ToResponse("删除标签成功")
	return
}


