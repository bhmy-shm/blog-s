package middleware

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"goweb/global"
	"goweb/pkg/MyJWT"
	"net/url"
)


func CasBin() gin.HandlerFunc {
	//1.加载鉴权文件
	E,_ := casbin.NewEnforcer("internal/middleware/model.conf","internal/middleware/p.csv")
	return func(context *gin.Context) {
		urlStr := context.Request.URL.String()
		method := context.Request.Method
		token := context.GetHeader("token")

		claims,_ := MyJWT.ParseToken(token)
		user := claims.AppUser

		if newUser ,ok := CheckAccess(user,urlStr,method,token) ; ok {
			//开始鉴权
			access , err := E.Enforce(newUser,urlStr,method)
			fmt.Println(access,err)
			if err != nil {
				context.Abort()
			}
			context.Next()
		}
	}
}

//模拟权限列表
var allowAccess = map[string][]map[string]string{
	"shm":{
		map[string]string{"method":"POST","path":"/blog/tags"},
		map[string]string{"method":"POST","path":"/blog/articles"},
		map[string]string{"method":"POST","path":"/upload/file"},
		map[string]string{"method":"PUT","path":"/blog/tags"},
		map[string]string{"method":"PUT","path":"/blog/articles"},
		map[string]string{"method":"DELETE","path":"/blog/tags"},
		map[string]string{"method":"DELETE","path":"/blog/articles"},
	},
	"tourist":{
		map[string]string{"method":"POST","path":"/blog/tags/count"},
		map[string]string{"method":"POST","path":"/blog/tags/list"},
		map[string]string{"method":"POST","path":"/blog/tags"},
		map[string]string{"method":"POST","path":"/blog/articles/list"},
		map[string]string{"method":"POST","path":"/blog/articles"},
	},
}

func checkUser(user,token string) bool {

	err := global.RedisEngin.HGet(context.TODO(),user,"token").Err()

	if err == nil{
		return true
	}
	return false
}
//判断用户身份信息
func CheckAccess(user,urlStr,method,token string )(string,bool) {

	//如果什么都没有代表错误
	if urlStr == "" || user == "" {
		return "",false
	}

	//判断token是否存在并赋予权限
	if !checkUser(user,token){
		user = "tourist"
	}else{
		user = "shm"
	}

	//如果url解析不对也代表错误
	_,err := url.ParseRequestURI(urlStr)
	if err != nil {
		return "",false
	}

	//如果能够从权限列表中找到用户信息则代表正确
	if access,ok := allowAccess[user];ok{
		//如果用户访问的方式，符合其在权限列表中指定的权限，则代表正确
		for _,access := range access {
			if method == access["method"]  {
				return user,true
			}
		}
	}
	return "",false
}