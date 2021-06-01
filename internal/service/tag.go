package service

import (
	"goweb/internal/model"
	"goweb/pkg/app"
)

//查看指定标签
type StateTagRequest struct {
	State uint8  `json:"state" form:"state" binding:"oneof=0 1"`
}

//统计行号
type CountRequest struct {
	State uint8  `json:"state" form:"state" binding:"oneof=0 1"`
	Name  string `json:"name" form:"name" binding:"max=100"`
}

//获取标签列表
type TagListRequest struct {
	State uint8  `json:"state" form:"state" binding:"oneof=0 1"`
	Name  string `json:"name" form:"name" binding:"max=100"`
}

//创建标签
type CreateTagRequest struct {
	State     uint8  `json:"state" form:"state" binding:"oneof=0 1"`
	Name      string `json:"name" form:"name" binding:"required,min=2,max=100"`
}

//更新标签
type UpdateTagRequest struct {
	ID         uint8 `json:"id" form:"id" binding:"required,gte=1"`
	State      uint8  `json:"state" form:"state" binding:"oneof=0 1"`
	Name       string `json:"name" form:"name" binding:"min=2,max=100"`
}

//删除标签
type DeleteTagRequest struct {
	ID uint8 `json:"id" form:"id" binding:"required,gte=1"`
}

func (svc Service) CreateTagTable() error {
	return svc.dao.CreateTagTable()
}

//统计标签总数
func (svc *Service) CountTag(param *CountRequest) (int, error) {
	return svc.dao.CountTag(param.State,param.Name)
}

//获取指定标签
func (svc *Service) GetTag(param *StateTagRequest,name string) (model.Tag, error) {
	return svc.dao.GetTag(name,param.State)
}

//获取标签列表
func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

//创建标签记录
func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State)
}

//修改标签记录
func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State)
}

//删除标签记录
func (svc *Service) DeleteTag(param *DeleteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
