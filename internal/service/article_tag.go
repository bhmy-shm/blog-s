package service

//创建文章标签关联表
func (svc Service) CreateArticleTagTable() error {
	return svc.dao.CreateArticleTagTable()
}

