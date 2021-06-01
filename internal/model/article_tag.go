package model

import (
	"gorm.io/gorm"
	"time"
)

type ArticleTag struct{
	*Model
	TagID uint32 `json:"tag_id" gorm:"type:int(10);unsigned"`	//标签ID
	ArticleID uint8 `json:"article_id" gorm:"type:int(11);not null"` //文章ID
}

//返回表名
func (a ArticleTag) tableName() string{
	return "blog_article_tag"
}

//创建文件标签关联数据表
func (a ArticleTag) CreateTable(db *gorm.DB) error {
	return db.Table(a.tableName()).AutoMigrate(&a)
}

//创建文章关联标签
func (a ArticleTag) Create(db *gorm.DB) error {
	return db.Table(a.tableName()).Create(&a).Error
}

//更新文章关联标签
func (a ArticleTag) Update(db *gorm.DB,articleId uint8,tagID uint32,i int) error {
	var ID int
	//先查到需要更改的关联表字段ID 值
	db.Table(a.tableName()).Select("id").Offset(i).Limit(1).Scan(&ID)

	//再根据字段ID去修改
	return db.Exec("UPDATE blog_article_tag SET tag_id=? WHERE id=?",tagID,ID).Error
}

//删除文章关联标签
func (a ArticleTag) Delete(db *gorm.DB) error{
	return db.Exec("Update blog_article_tag SET is_del=?,deleted_at=? WHERE article_id=?",1,time.Now().Unix(),a.ArticleID).Error
}

//查看指定ID的文章关联标签
func(a ArticleTag) GetByArticleID(db *gorm.DB) ([]ArticleTag,error){
	var articleTag []ArticleTag
	err := db.Table(a.tableName()).Where("article_id=? AND is_del = ?",a.ArticleID,0).Find(&articleTag).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return articleTag,err
	}
	return articleTag,nil
}

//通过标签ID查文章ID
func (a ArticleTag) GetArticleIDByTagID(db *gorm.DB,tagID []int) ([]int,error){
	var articleIDs []int
	err := db.Table(a.tableName()).Select("article_id").Where("tag_id IN ?",tagID).Scan(&articleIDs).Error
	if err != nil {
		return nil,err
	}
	return articleIDs,nil
}
