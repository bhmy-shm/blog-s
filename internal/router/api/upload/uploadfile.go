package upload

import (
	"github.com/gin-gonic/gin"
	"goweb/global"
	"goweb/internal/Check"
	"goweb/internal/service"
	"goweb/pkg/app"
	"goweb/pkg/errcode"
	"goweb/pkg/util"
)

func UploadFile(c *gin.Context){

	response := app.NewResponse(c)

	//1.接收文件信息
	file,fileHeader,err := c.Request.FormFile("file")
	fileType := util.StrTo(c.PostForm("type")).MustInt()

	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	//2.开始保存文件
	svc := service.New(c)
	fileInfo,err := svc.UploadFile(Check.FileType(fileType),file,fileHeader)
	if err != nil{
		global.Logger.ErrorF(c,"svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(fileInfo.AccessUrl)
}
