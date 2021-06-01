package articles

import (
	"github.com/gin-gonic/gin"
	"goweb/global"
	"goweb/internal/service"
	"goweb/pkg/app"
	"goweb/pkg/errcode"
	"goweb/pkg/util"
)

type articles struct{}

func NewArticles() articles {
	return articles{}
}

//创建文章表
func (t articles) CreateTable(c *gin.Context) {
	svc := service.New(c.Request.Context())
	err := svc.CreateArticleTable()
	if err != nil {
		c.JSON(400,gin.H{"message":"创建文章表失败"})
	}else{
		c.JSON(200,gin.H{"message":"创建文章表成功"})
	}
}
//创建文章标签关联表
func (t articles) CreateAtTable(c *gin.Context) {
	svc := service.New(c.Request.Context())
	err := svc.CreateArticleTagTable()
	if err != nil {
		c.JSON(400,gin.H{"message":"创建文章标签关联表失败"})
	}else{
		c.JSON(200,gin.H{"message":"创建文章标签关联表成功"})
	}
}

//查看指定文章
func (t articles) Get(c *gin.Context) {

	param := service.ArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)

	if valid == true {
		global.Logger.ErrorF(c,"app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	article, err := svc.GetArticle(&param)

	if err != nil {
		global.Logger.ErrorF(c,"svc.GetArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticleFail)
		return
	}

	response.ToResponse(article)
	return
}

//查看文章列表
func (t articles) List(c *gin.Context) {
	param := service.ArticleListRequest{}
	valid, errs := app.BindAndValid(c, &param)

	response := app.NewResponse(c)
	if valid == true {
		global.Logger.ErrorF(c,"app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())

	//
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	articles,totalRows,err := svc.GetArticleList(&param,&pager)
	if err != nil {
		global.Logger.ErrorF(c,"svc.GetArticleList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetArticlesFail)
		return
	}

	response.ToResponseList(articles, totalRows)
	return

}

//创建文章
func (t articles) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)

	//如果参数校验出错
	if valid == true {
		global.Logger.ErrorF(c,"app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	//入库创建文章
	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.ErrorF(c,"svc.CreateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	response.ToResponse("创建文章成功")
	return
}

//更新文章
func (t articles) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{}
	response := app.NewResponse(c)
	valid,errs := app.BindAndValid(c,&param)

	//如果参数校验出错
	if valid == true {
		global.Logger.ErrorF(c,"app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	//入库创建文章
	svc := service.New(c.Request.Context())
	err := svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.ErrorF(c,"svc.UpdateArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateArticleFail)
		return
	}
	response.ToResponse("修改文章成功")
	return
}

//删除文章
func (t articles) Delete(c *gin.Context) {
	//从路由中拿到要删除的文章
	id := util.StrTo(c.Param("id")).MustUInt8()

	param := service.DeleteArticleRequest{ID: id}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if valid == true {
		global.Logger.ErrorF(c,"app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteArticle(&param)
	if err != nil {
		global.Logger.ErrorF(c,"svc.DeleteArticle err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteArticleFail)
		return
	}
	response.ToResponse("删除文章成功")
	return
}

