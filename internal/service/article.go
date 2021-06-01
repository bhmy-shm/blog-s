package service

import (
	"fmt"
	"goweb/internal/dao"
	"goweb/internal/model"
	"goweb/pkg/app"
	"goweb/pkg/util"
)

//返回指定的文章
type ArticleRequest struct {
	ID    uint8 `json:"id" form:"id" binding:"required,gte=1"`
	State uint8  `json:"state" form:"state,default=0" binding:"oneof=0 1"`
}

//返回文章列表
type ArticleListRequest struct {
	TagID uint32 `json:"tag_id" form:"tag_id" binding:"gte=1"`
	Name []string `json:"tag_name" form:"name" binding:"required"`
	State uint8  `json:"state" form:"state,default=0" binding:"oneof=0 1"`
}

//创建文章
type CreateArticleRequest struct {
	TagID         []uint32 `json:"tag_id" form:"tag_id" binding:"required,gte=1"`
	Title         string `json:"title" form:"title" binding:"required,min=2,max=100"`
	Desc          string `json:"desc" form:"desc" binding:"required,min=2,max=255"`
	Content       string `json:"content" form:"content" binding:"required,min=2,max=4294967295"`
	CoverImageUrl string `json:"cover_image_url" form:"cover_image_url" binding:"required,url"`
	State         uint8  `json:"state" form:"state,default=0" binding:"oneof=0 1"`
}

//更新文章
type UpdateArticleRequest struct {
	ID            uint8 `json:"id" form:"id" binding:"required,gte=1"`
	TagID         []uint32 `json:"tag_id" form:"tag_id" binding:"required,gte=1"`
	Title         string `json:"title" form:"title" binding:"min=2,max=100"`
	Desc          string `json:"desc" form:"desc" binding:"min=2,max=255"`
	Content       string `json:"content" form:"content" binding:"min=2,max=4294967295"`
	CoverImageUrl string `json:"cover_image_url" form:"cover_image_url" binding:"required,url"`
	State         uint8  `json:"state" form:"state,default=0" binding:"oneof=0 1"`
}
//文章列表
type Article struct {
	ID            uint8     `json:"id"`
	Title         string     `json:"title"`
	Desc          string     `json:"desc"`
	Content       string     `json:"content"`
	CoverImageUrl string     `json:"cover_image_url"`
	State         uint8      `json:"state"`
	Tag           []string  `json:"tag"`
}

//删除文章
type DeleteArticleRequest struct {
	ID uint8 `json:"id" form:"id" binding:"required,gte=1"`
}

func (svc Service) CreateArticleTable() error {
	return svc.dao.CreateArticleTable()
}

//创建文章表数据，同时创建文章标签关联表数据
func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	//先创建文章
	article, err := svc.dao.CreateAritcle(&dao.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		State:         param.State,
	})
	if err != nil {
		return err
	}

	//然后创建文章标签关联表
	err = svc.dao.CreateArticleTag(article.ID, param.TagID)
	if err != nil {
		return err
	}
	return nil
}

//更新文章表数据，同时更新文章标签关联表数据
func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	//先更新文章数据
	err := svc.dao.UpdateArticle(&dao.Article{
		ID:            param.ID,	//文章ID
		Title:         param.Title,	//文章标题
		Desc:          param.Desc,	//文章简述
		Content:       param.Content,	 //文章内容
		CoverImageUrl: param.CoverImageUrl,	//文章图片
		State:         param.State,	//文章状态
	})
	if err != nil {
		return err
	}
	//再更新关联表的标签
	if param.TagID != nil {
		err = svc.dao.UpdateArticleTag(param.ID, param.TagID)
		return err
	}
	return nil
}

//删除指定id的文章
func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	err := svc.dao.DeleteArticle(param.ID)
	if err != nil {
		fmt.Println("zz",err)
		return err
	}
	err = svc.dao.DeleteArticleTag(param.ID)
	if err != nil {
		fmt.Println("XX",err)
		return err
	}
	return nil
}

//查看指定id的文章
func (svc *Service) GetArticle(param *ArticleRequest) (*Article, error) {

	var tags []string

	//拿到指定id的文章结果
	article, err := svc.dao.GetByIDArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}

	//拿到指定id的文章的标签信息
	articleTag, err := svc.dao.GetArticleTagByID(article.ID)
	if err != nil {
		return nil, err
	}

	//再根据拿到的指定ID标签的信息
	for _,articletag := range articleTag{
		tag, err := svc.dao.GetTagByID(uint8(articletag.TagID), model.STATE_OPEN)
		if err != nil {
			return nil, err
		}
		tags = append(tags,tag)
	}

	return &Article{
		ID:            article.ID,
		Title:         article.Title,
		Desc:          article.Desc,
		Content:       article.Content,
		CoverImageUrl: article.CoverImageUrl,
		State:         article.State,
		Tag:           tags,
	}, nil
}

//获取文章列表
func (svc *Service) GetArticleList(param *ArticleListRequest,pager *app.Pager)([]model.Article,int,error) {
	var articles []model.Article

	//1.通过标签名字，获取对应得标签ID号
	tagIDs,err := svc.dao.CountArticleListByTagName(param.Name,param.State)
	if err != nil {
		fmt.Println("11",err)
		return nil,0,err
	}

	//2.根据标签ID找到使用该标签的文章ID
	articleIDs,err := svc.dao.GetArticleIDByTagID(tagIDs)
	if err != nil{
		fmt.Println("22",err)
		return nil,0,err
	}
	//2-2.这里面的ID有可能会重复，需要去重,然后按照降序排列
	NewIDS := util.DescIDArr(util.IDWeight(articleIDs))

	//3.根据文章ID获取指定的列表信息
	articles,err = svc.dao.GetArticleByArticleID(NewIDS,param.State,pager.Page, pager.PageSize)
	if err != nil {
		fmt.Println("33",err)
		return nil,0,err
	}
	return articles,len(articles),nil
}
