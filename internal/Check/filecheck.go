package Check

import (
	"fmt"
	"goweb/global"
	"goweb/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

//对上传后的文件名进行格式化，用md5加密后在进行写入，避免直接暴露原始文件。
type  FileType int

//定义上传文件的类型，如果有其它的上传文件可以如下定义：
//const TypeImage FileType = iota+1
//TypeExcel
//TypeTxt
//TypeWord
const (
	TypeImage FileType = 1
	TypeAll FileType = 2
)


//获取文件名称
func GetFileName(name string) string{
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name,ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}
//获取文件后缀
func GetFileExt(name string) string{
	return path.Ext(name)
}
//获取文件的保存地址
func GetSaveImagePath() string{
	return global.AppSetting.UploadSavePath
}
//获取其它类型文件的保存地址
func GetSaveFilePath() string{
	return global.AppSetting.UploadSavePath + `file\`
}

/* -------------------------------- 检查文件的相关方法，确保文件在写入时已经达到了必备条件----------------------------------*/
//检验存储路径
func CheckSavePath(dst string) bool{
	fileInfo,err := os.Stat(dst)
	fmt.Println("save path:",fileInfo,err)
	if err == nil {
		if len(fileInfo.Name()) > 0{
			return true
		}
		return false
	}else{
		return false
	}
}

//检验文件后缀
func CheckContainExt(t FileType,name string) bool{
	ext := GetFileExt(name)
	switch t {
	case TypeImage:
		for _,allowExt := range global.AppSetting.UploadImageAllowExts{
			if strings.ToUpper(allowExt) == strings.ToUpper(ext){
				return true
			}
		}
	}
	return false
}
//检验文件大小
func CheckMaxSize(t FileType,f multipart.File) bool {
	content,_ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024{
			return true
		}
	case TypeAll:
		if size >= global.AppSetting.UploadFileAllSize*1024*1024{
			return true
		}
	}
	return false
}

//检验文件权限是否足够
func CheckPermission(dst string) bool{
	_,err := os.Stat(dst)
	return os.IsPermission(err)
}

/*-------------------------------------- 通过检验后写入或者创建图片 -------------------------------------*/
func CreateSavePath(dst string,perm os.FileMode) error{
	err := os.Mkdir(dst,perm)
	if err != nil {
		fmt.Println("create path err:",err)
		return err
	}
	return nil
}

func SaveFile(file *multipart.FileHeader,dst string) error{
	src,err := file.Open()
	if err != nil{
		return err
	}
	defer src.Close()

	out,err := os.Create(dst)
	if err != nil{
		return err
	}
	defer out.Close()

	_,err = io.Copy(out,src)
	return err
}