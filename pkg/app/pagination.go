package app

import (
	"github.com/gin-gonic/gin"
	"goweb/global"
	"goweb/pkg/util"
)

type Pager struct{
	Page int `json:"page"`
	PageSize int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

//查看指定页数
func GetPage(c *gin.Context) int {
	page := util.StrTo(c.Query("page")).MustInt()
	if page <=0 { return 1 }
	return page
}

//返回每一页的数据个数，默认10个，上限100个
func GetPageSize(c *gin.Context) int {
	pageSize := util.StrTo(c.Query("page_size")).MustInt()
	//查看每页默认的10个
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	//最大页数
	if pageSize > global.AppSetting.MaxPageSize  {
		return global.AppSetting.MaxPageSize
	}
	return pageSize
}
//分页
func GetPageOffset(page,pageSize int) int{
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}