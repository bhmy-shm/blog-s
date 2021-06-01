package dao

import (
	"goweb/internal/model"
	"goweb/pkg/app"
)

//创建标签表
func (d Dao) CreateTagTable()error{
	var tag model.Tag
	return tag.CreateTable(d.db)
}

//统计标签总数
func(d *Dao) CountTag(state uint8,name string)(int,error){
	tag := model.Tag{Name: name,State: state}
	return tag.Count(d.db)
}

//根据标签名字查看指定标签的全部信息
func(d Dao) GetTag(name string,state uint8)(model.Tag,error){
	tag := model.Tag{Name: name,State: state}
	return tag.Get(d.db,name,state)
}
//根据标签id查看指定标签的全部信息
func (d *Dao) GetTagByID(tagid uint8, state uint8) (string, error) {
	tag := model.Tag{Model:&model.Model{ID: tagid}}
	return tag.GetByID(d.db)
}

//查看标签列表
func(d *Dao) GetTagList(name string,state uint8,page,pageSize int) ([]*model.Tag,error){
	tag := model.Tag{Name:name,State: state}
	pageOffset := app.GetPageOffset(page,pageSize)
	return tag.List(d.db,pageOffset,pageSize)
}

//创建标签
func(d Dao) CreateTag(name string,state uint8) error {
	tag := model.Tag{Name: name,State: state}
	return  tag.Create(d.db)
}

//更新标签
func(d Dao) UpdateTag(id uint8,name string,state uint8)error {
	tag := model.Tag{Model:&model.Model{ID: id}}	//更新判断条件id号
	values:= map[string]interface{}{
		"name":name,
		"state":state,
	}
	return tag.Update(d.db,values)
}

//删除标签
func (d Dao) DeleteTag(id uint8) error {
	tag := model.Tag{Model:&model.Model{ID: id}}
	return tag.Delete(d.db)
}