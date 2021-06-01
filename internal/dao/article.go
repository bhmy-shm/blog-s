package dao

import (
	"gorm.io/gorm"
	"goweb/internal/model"
)

type Article struct{
	ID uint8 `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State uint8 `json:"state"`
}


//创建文章表
func (dao Dao) CreateArticleTable()error{
	var article model.Article
	return article.CreateTable(dao.db)
}

//创建文章
func(dao Dao) CreateAritcle(param *Article)(*model.Article,error){
	article := model.Article{
		Title:param.Title,
		Desc:param.Desc,
		Content:param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State: param.State,
	}
	return article.Create(dao.db)
}

//更新文章
func(dao Dao) UpdateArticle(param *Article) error {

	article := model.Article{Model:&model.Model{ID: param.ID}}

	values := map[string]interface{}{
		"state":param.State,
	}
	if param.Title != ""{
		values["title"] = param.Title
	}
	if param.CoverImageUrl != ""{
		values["cover_image_url"] = param.CoverImageUrl
	}
	if param.Desc != ""{
		values["desc"] = param.Content
	}
	if param.Content != ""{
		values["content"] = param.Content
	}

	return article.Update(dao.db,values)
}

//删除文章
func(dao Dao) DeleteArticle(id uint8) error {
	article := model.Article{Model:&model.Model{ID: id}}
	return article.Delete(dao.db)
}
//查看指定文章
func(dao Dao) GetByIDArticle(id uint8,state uint8) (model.Article,error) {
	article := model.Article{Model:&model.Model{ID: id},State: state}
	result,err := article.Get(dao.db)
	if err != nil && err != gorm.ErrRecordNotFound {
		return result,err
	}else{
		return result,nil
	}
}

//查看指定标签的文章，所使用的标签ID。
func(dao Dao) CountArticleListByTagName(tagName []string,state uint8) ([]int,error){
	var tagsId []int
	tag := model.Tag{State: state}
	for _,name := range tagName{
		id,err  := tag.CountByTagID(dao.db,name)
		if err != nil {
			return nil,err
		}
		tagsId = append(tagsId,id)
	}
	return tagsId,nil
}

