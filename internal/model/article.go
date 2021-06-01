package model

import (
	"gorm.io/gorm"
	"time"
)

type Article struct{
	*Model
	Title string `json:"title,omitempty" gorm:"type:varchar(100)"`
	Desc string	`json:"desc,omitempty" gorm:"type:varchar(255)"`
	Content string	`json:"content,omitempty" gorm:"type:longtext"`
	CoverImageUrl string `json:"cover_image_url,omitempty" gorm:"type:varchar(255)"`
	State uint8 `json:"state,omitempty" gorm:"type:tinyint(3)"`
}

//表名
func (a Article) tableName() string{
	return "blog_article"
}

//创建文章表
func (a Article) CreateTable(db *gorm.DB) error{
	return db.Table(a.tableName()).AutoMigrate(&a)
}

//创建文章
func (a Article) Create(db *gorm.DB) (*Article,error) {
	err := db.Table(a.tableName()).Create(&a).Error
	return &a,err
}

//更新文章
func (a Article) Update(db *gorm.DB,values interface{}) error{
	return db.Table(a.tableName()).Where("id =? And is_del=?",a.ID,0).Updates(&values).Error
}

//删除文章
func (a Article) Delete(db *gorm.DB) error{
	return db.Exec("Update blog_article SET is_del=?,deleted_at=? WHERE id=?",1,time.Now().Unix(),a.Model.ID).Error
}

//查看指定文章
func (a Article) Get(db *gorm.DB) (Article,error){
	var article Article
	err := db.Table(a.tableName()).Select("*").Where("id=? and is_del=? and state=?",a.ID,0,0).Scan(&article).Error
	if err != nil{
		return article,err
	}
	return article,nil
}

//通过文章ID查文章列表
func (a Article) GetArticlesByID(db *gorm.DB,articleIDs []int,pageOffset,pageSize int)([]Article,error){
	var articls []Article
	err := db.Table(a.tableName()).Offset(pageOffset).
		Select("id,created_at,title,blog_article.desc").
		Where("state=? AND id IN ?",a.State,articleIDs).
		Limit(pageSize).Scan(&articls).Error
	if err != nil{
		return nil,err
	}
	return articls,nil
}