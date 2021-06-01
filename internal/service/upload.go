package service

import (
	"errors"
	"fmt"
	"goweb/global"
	"goweb/internal/Check"
	"mime/multipart"
	"os"
)

type FileInfo struct{
	Name string
	AccessUrl string
}

//根据不同的文件类型进行上传
func (svc Service)UploadFile(fileType Check.FileType,file multipart.File,fileHeader *multipart.FileHeader) (*FileInfo,error){

	var fileName string
	var dst string
	var uploadSavePath string

	if fileType == Check.TypeImage {
		//图片的基本信息
		fileName = Check.GetFileName(fileHeader.Filename)
		uploadSavePath = Check.GetSaveImagePath()
		dst = uploadSavePath + fileName
	}else if fileType == Check.TypeAll{
		//其它文件的基本信息
		fileName = Check.GetFileName(fileHeader.Filename)
		uploadSavePath = Check.GetSaveFilePath()
		dst = uploadSavePath + fileName
	}

	//2.对文件进行检查，判断目录是否存在	,false代表没有，则需要创建；true代表有
	if  ok := Check.CheckSavePath(uploadSavePath) ; ok == false{
		fmt.Println("准备创建文件目录",ok)
		//2-1.在本地创建保存文件的路径
		if err := Check.CreateSavePath(uploadSavePath,os.ModePerm) ; err !=nil{
			return nil,err
		}
	}

	//3.测试文件大小是否符合要求
	if Check.CheckMaxSize(fileType,file){
		return nil,errors.New("exceeded maximum file limit.")
	}
	//4.监测文件权限是否足够
	if Check.CheckPermission(uploadSavePath) {
		return nil,errors.New("insufficient file permissions.")
	}
	//5.是否具备写入条件
	if err := Check.SaveFile(fileHeader,dst) ; err != nil{
		return nil,err
	}

	//真正的写入文件
	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName,AccessUrl: accessUrl},nil
}
