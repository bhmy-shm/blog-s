# blog-s

Gin编写 - goweb博客后端，采用MVC结构，持续迭代中



## 项目设计

功能介绍：

- 提供用户注册，登录
- 提供token验证，Casbin验证
- 登录用户可以增删查改文章，标签，目录
- 非注册用户只能查看文章，标签，目录
- 接口限流
- 全部的后端API接口

待实现功能：

- 媒体库
- 评论区
- 前端页面

目录结构

- cmd：consumer消费者
- configs：全局配置
- docs：swagger接口文档
- global：全局模块
- internal：内部模块
  - model：数据库模型
  - dao：数据访问层
  - check：数据校验层
  - Service：数据逻辑层
  - middleware：中间件
  - router：路由
- pkg：第三方工具模块
- storage：存放运行日志，错误日志，上传的文件



## 公共组件

- 错误码标准化

- 全局配置管理

- 日志写入

- 第三方中间件连接（Redis，Mysql，RabbitMQ）

- 分页响应处理

- Swagger接口文档

  

## 路由接口

目前实现的路由接口，只能采用ApiPost的方式访问

```go
func NewRouter() *gin.Engine{

	register := register.Newregister()	//注册模块
	login := login.Newlogin()			//登录模块
	tag := tag.NewTag()					//标签模块
	articles := articles.NewArticles()	//文章模块


	r := gin.Default()

	r.Use(middleware.Translations())	//多语言类型中间件
	r.Use(middleware.AccessLog())		//访问日志记录
	r.Use(middleware.Recovery())		//panic捕获
	r.Use(middleware.RateLimit(methodLimiters))	//令牌桶
	r.Use(middleware.ContextTimeout(60*time.Second))	//统一超时时间控制
	r.Use(middleware.Tracing())	//Tracer链路追踪
    
    //swagger
	r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))

	//注册
	blog := r.Group("")
	{
		blog.POST("/register/item",register.Post)	//创建用户表
		blog.POST("/register/auth",register.Register)	//注册用户
		blog.POST("/register/authentication",register.Authentication)	//输入注册时的手机验证码
	}
	//登录
	{
		blog.POST("/login/creat",login.CreateJwt)	//创建jwt密文表
		blog.POST("/login/getToken",login.GetUserJwt) //用用申请token编码
		blog.Use(middleware.JWT(),middleware.CasBin())	//JWT验证中间件，在登录时使用
		blog.POST("/login/users",login.UserLogin)	//用户登录
	}

	//业务
	service := r.Group("/blog")
	{
		//查看
		service.Use(middleware.CasBin())
		service.POST("/tags/count",tag.Count) //获取标签总数
		service.POST("/tags/list",tag.List)	//获取标签列表
		service.POST("/tags/:name",tag.Get) //获取指定标签

		service.POST("/articles/list",articles.List)	//获取文章列表
		service.POST("/articles/:id",articles.Get)	 //获取指定文章

		//增删改
		service.Use(middleware.MustLogin())	//必须属于登录之后的用户才能，必须携带合格的JWT才属于登录

		service.POST("/tags",tag.Create)	//新增标签
		service.PUT("/tags/:id",tag.Update)	//更新标签
		service.DELETE("/tags/:id",tag.Delete)	//删除标签
		service.POST("/articles",articles.Create) //新增文章
		service.PUT("/articles",articles.Update)	//更新文章
		service.DELETE("/articles/:id",articles.Delete)	//删除文章

		//建表
		//casbin鉴权中间件，只有管理员用户才能创建对应的数据表
		service.POST("/tags/CreateTable",tag.CreateTable)	//创建标签表
		service.POST("/articles/CreateTable",articles.CreateTable)	//创建文章表
		service.POST("/at/CreateTable",articles.CreateAtTable) //创建文章标签关联表
	}

	//上传图片
	r.Use(middleware.MustLogin())
	r.POST("/upload/file",upload.UploadFile)
	r.StaticFS("/static",http.Dir(global.AppSetting.UploadSavePath))

	return r
}
```

