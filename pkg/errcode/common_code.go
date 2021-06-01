package errcode

var (
	//注册登录
	Success	= NewError(0,"成功")
	ServiceError = NewError(1000000,"服务器内部错误")
	InvalidParams = NewError(1000001,"入参错误")
	NotFound = NewError(1000002,"找不到")

	//用户注册、登录
	ErrorUserRegister = NewError(10011,"用户注册失败")
	ErrorUserRegisterRepetition = NewError(10013,"用户已存在")
	ErrorUserRegisterAuth = NewError(10012,"用户验证码填写错误")
	ErrorUserLogin = NewError(20011,"用户登录账户密码错误")
	ErrorUserRedirect = NewError(30010,"跳转到验证码界面，请输入验证码")

	//JWT鉴权
	UnauthorizedAuthNotExist = NewError(1000003,"鉴权失败,找不到对应的Appkey 和 AppSecret")
	UnauthorizedTokenError = NewError(1000004,"鉴权失败，Token错误")
	UnauthorizedTokenTimeOut = NewError(1000005,"鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(1000006,"鉴权失败,Token生成失败")
	TooManyRequests = NewError(1000007,"请求过多")

	//标签
	ErrorCountTagField = NewError(2001001,"统计标签总数失败")
	ErrorGetTagListField = NewError(2001002, "获取标签列表失败")
	ErrorCreateTagField = NewError(2001003,"创建标签失败")
	ErrorUpdateTagField = NewError(2001004,"更新标签失败")
	ErrorDeleteTagField = NewError(2001005,"删除标签失败")
	ErrorGetTagField = NewError(2001006, "获取指定标签失败")

	//文章
	ErrorGetArticleFail    = NewError(3001001, "获取单个文章失败")
	ErrorGetArticlesFail   = NewError(3001002, "获取多个文章失败")
	ErrorCreateArticleFail = NewError(3001003, "创建文章失败")
	ErrorUpdateArticleFail = NewError(3001004, "更新文章失败")
	ErrorDeleteArticleFail = NewError(3001005, "删除文章失败")

	//上传图片
	ErrorUploadFileFail = NewError(2002001,"上传文件失败")
)
