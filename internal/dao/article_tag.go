package dao

import (
	"goweb/internal/model"
	"goweb/pkg/app"
)

//创建文章标签关联表表
func (dao Dao) CreateArticleTagTable()error{
	var articleTag model.ArticleTag
	return articleTag.CreateTable(dao.db)
}

//文章关联表创建数据
func (d *Dao) CreateArticleTag(articleID uint8, tagID []uint32) error {
	var err error
	for _,tagid := range tagID{
		articleTag := model.ArticleTag{
			ArticleID: articleID,
			TagID:     tagid,
		}
		err = articleTag.Create(d.db)
	}
	return err
}

//更新文章关联表数据
func (d *Dao) UpdateArticleTag (articleId uint8,tagID []uint32) error {

	articleTag := model.ArticleTag{ArticleID: articleId}
	var err error
	for i,tagid := range tagID{
		err = articleTag.Update(d.db,articleId,tagid,i)
	}
	return err
}

//删除文章关联表数据
func(d *Dao) DeleteArticleTag (articleId uint8) error {
	article := model.ArticleTag{ArticleID: articleId}
	return article.Delete(d.db)
}

//查看指定文章ID的关联标签
func(d *Dao) GetArticleTagByID (articleID uint8) ([]model.ArticleTag,error) {
	articleTag := model.ArticleTag{ArticleID: articleID}
	return articleTag.GetByArticleID(d.db)
}

//查看指定标签ID的文章
func(d *Dao) GetArticleIDByTagID(tagIDs []int) ([]int,error){
	articleTag := model.ArticleTag{}
	return articleTag.GetArticleIDByTagID(d.db,tagIDs)
}

//根据文章ID查看文章
func(d *Dao) GetArticleByArticleID(articles []int,state uint8,page,pageSize int)([]model.Article,error){
	article := model.Article{State: state}
	pageOffset := app.GetPageOffset(page,pageSize)	//拿到偏移量

	return article.GetArticlesByID(d.db,articles,pageOffset,pageSize)
}