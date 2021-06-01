package model

import (
	"gorm.io/gorm"
	"time"
)

type Tag struct {
	*Model
	State uint8	`json:"state" gorm:"type:tinyint(3);unsigned"`
	Name string `json:"name" gorm:"type:varchar(100)"`
}

//标签名
func (t Tag) tableName() string {
	return "blog_tag"
}
//创建标签表
func (t Tag) CreateTable(db *gorm.DB) error {
	return db.Table(t.tableName()).AutoMigrate(&t)
}

//统计标签总数
func(t Tag) Count(db *gorm.DB) (int,error){
	var count int64
	err := db.Table(t.tableName()).Select("name").Where("name=? and state=?",t.Name,t.State).Count(&count).Error

	return int(count), err
}

//根据标签名获取指定标签
func(t Tag) Get(db *gorm.DB,name string,state uint8) (Tag,error) {
	var tag Tag
	err := db.Table(t.tableName()).Select("*").Where("name = ? and state=?",name,state).Scan(&tag).Error
	return tag,err
}
//根据标签ID获取指定标签
func(t Tag) GetByID(db *gorm.DB) (string,error) {
	var tag string
	err := db.Table(t.tableName()).Select("name").Where("id = ? AND is_del = ? AND state = ?", t.Model.ID, 0, t.State).Scan(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}
	return tag, nil
}



//查看指定的标签列表
func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		err = db.Table(t.tableName()).Offset(pageOffset).Where("state=?",t.State).Limit(pageSize).Find(&tags).Error
	}
	if err != nil {
		return nil, err
	}
	return tags, nil
}


func(t Tag) Create(db *gorm.DB)error{
	return db.Table(t.tableName()).Create(&t).Error
}

func(t Tag) Update(db *gorm.DB,values interface{})error{
	err := db.Table(t.tableName()).Where("id = ? And is_del = ?",t.ID,0).Updates(&values).Error
	err = db.Exec("update blog_tag set updated_at=? where id=?",time.Now().Unix(),t.ID).Error
	return err
}

func(t Tag) Delete(db *gorm.DB)error{
	return db.Table(t.tableName()).Where("id = ? And is_del = ?",t.ID,0).Unscoped().Delete(&t).Error
}


//List查看指定标签名字得ID号
func (t Tag) CountByTagID(db *gorm.DB,tagName string)(int,error){
	var id int
	err := db.Table(t.tableName()).Select("id").Where("name= ?",tagName).Scan(&id).Error
	if err != nil {
		return 0,err
	}
	return id,nil
}