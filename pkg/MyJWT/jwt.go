package MyJWT

import (
	"github.com/dgrijalva/jwt-go"
	"goweb/global"
	"goweb/pkg/util"
	"time"
)

type Claims struct{
	AppUser string `json:"app_user"`
	AppKey string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims	//jwt签证信息
}

var jwtkey []byte

//拿到jwt密钥，必须是byte切片不然报错
func GetJWTSecret() []byte{
	return jwtkey
}

//jwt注册
func GenerateToken(appKey, appSecret ,userphone string) (string, error) {
	jwtkey = []byte(appSecret)	//为密钥赋值
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire*time.Second) //配置文件定义的过期时间
	claims := &Claims{
		AppUser:   userphone,
		AppKey:    util.EncodeMD5(appKey), //用户发送的密码密钥1
		AppSecret: util.EncodeMD5(appSecret), //用户发送的密钥2
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),     //token过期时间
			IssuedAt:  time.Now().Unix(),     //token发放时间
			Issuer:    global.JWTSetting.Issuer,	//签发者
			Subject:   global.JWTSetting.Topic,     //主题
		},
	}

	//生成token编码
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())

	return token, err
}

//jwt解码
func ParseToken(token string) (*Claims,error){
	tokenClaims,err := jwt.ParseWithClaims(token,&Claims{},func(token *jwt.Token)(interface{},error){
		return GetJWTSecret(),nil
	})
	if err !=nil{
		return nil,err
	}
	if tokenClaims != nil{
		if claims,ok := tokenClaims.Claims.(*Claims);ok && tokenClaims.Valid{
			return claims,nil	//解码之后的数据
		}
	}
	return nil,err
}

