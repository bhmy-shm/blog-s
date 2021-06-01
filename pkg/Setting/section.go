package Setting

import "time"

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize       int
	MaxPageSize           int
	UploadImageMaxSize	  int
	UploadFileAllSize	  int
	DefaultContextTimeout time.Duration
	LogSavePath           string
	AccessLogSavePath	  string
	LogFileName           string
	LogFileExt            string
	UploadSavePath string
	UploadServerUrl string
	UploadImageAllowExts []string
}

type DatabaseSettingS struct {
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
}

type RabbitMQS struct{
	Port	uint
	Name	string
	Pass	string
	Host	string
	Exchange string
	Queue string
	Router string
}

//JWT映射结构体
type JWTSettingS struct{
	Secret string
	Issuer string
	Topic string
	Expire time.Duration
}

//Redis数据库
type RedisSettingS struct{
	Poolize int  //连接池数量
	MinIdleConns int  //最小连接数
	DB int
	DialTimeout time.Duration //连接建立超时时间
	Addr string
	Port string
	PassWord string
}

//Email邮箱
type EmailSettings struct{
	Host string
	Port int
	UserName string
	PassWord string
	IsSSL bool
	From string
	To []string
}

var sections = make(map[string]interface{})

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}
	return nil
}

func (s *Setting) ReloadAllSection() error {
	for k,v := range sections{
		err := s.ReadSection(k,v)
		if err != nil {
			return err
		}
	}
	return nil
}