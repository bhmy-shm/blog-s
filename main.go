package main

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"goweb/cmd"
	"goweb/global"
	"goweb/internal/model"
	"goweb/internal/router"
	"goweb/pkg/MyLogger"
	"goweb/pkg/RabbitMQ"
	"goweb/pkg/Redis"
	"goweb/pkg/Setting"
	"goweb/pkg/tracer"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	config  string
	port string
	runMode string
)

func init(){
	//初始化配置文件
	err := setupSetting()
		if err != nil{ log.Fatalf("init.SetupSetting err:%v",err) }
	//初始化日志
	err = setupLogger()
		if err != nil { log.Fatalf("init.SetupLogger err%s",err) }
	//初始化Mysql数据库
	err = setupDBEngine()
		if err != nil { log.Fatalf("init.SetupEngine err:%v",err) }
	//初始化Redis
	err = setupRedis()
		if err != nil { log.Fatalf("init.SetupRedis err:%v",err) }
	err = setupTracer()
		if err != nil { log.Fatalf("init.SetupTracer err:%v",err) }
	err = setupFlag()
		if err != nil { log.Fatalf("init.SetupFlag err:%v",err) }

	//初始化RabbitMQ交换器
	err = RabbitMQ.MQExchangeInit()
	if err != nil {
		log.Println(err)
	}
	//初始化consumer，for循环开启多个消费者
	go func() {
		err = cmd.NewAuthConsumer("authConsumer")
		if err != nil {
			log.Println(err)
		}
	}()
}


//@title 博客系统
//@version 1.0
//@description Go-web 孙海铭的入门项目
//@termsOfService http://www.bhmy.top
//@contact.name github地址
//@contact.url https://github.com/
func main(){
	gin.SetMode(global.ServerSetting.RunMode)

	router := router.NewRouter()

	s := http.Server{
		Addr: ":"+global.ServerSetting.HttpPort,
		Handler: router,
		ReadTimeout: global.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout: global.ServerSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1<<20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServer err %v", err)
		}
	}()

	//等待信号中断
	quit := make(chan os.Signal)
	//接收 syscall.SIGINT  和 syscall.SIGTERM 信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shuting down server...")

	//最大控制时间，通知该服务端，它有5s时间处理原来的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func setupSetting() error {
	setting,err := Setting.NewSetting(strings.Split(config,",")...)
	if err != nil{ return err }
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil { return err }
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil { return err }
	err = setting.ReadSection("Database", &global.DataBaseSetting)
	if err != nil { return err }
	err = setting.ReadSection("RabbitMQ", &global.RabbitMQSetting)
	if err != nil { return err}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil { return err }
	err = setting.ReadSection("Redis", &global.RedisSetting)
	if err != nil { return err }
	err = setting.ReadSection("Email",&global.EmailSetting)
	if err != nil { return err }

	if port != ""{
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}
	return nil
}

//日志记录
func setupLogger() error {
	global.Logger = MyLogger.NewLogger(&lumberjack.Logger{
		Filename: global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,  //设置日志文件允许的最大占用空间 600MB
		MaxAge:    10,   //日志文件的最大声明周期 10天
		LocalTime: true, //日志文件名的时间格式为本地时间
	},"",log.LstdFlags).WithCaller(2)	//找两层栈调用进行记录

	return nil
}

//GROM -Mysql
func setupDBEngine() error {
	var err error
	global.DBEngine,err = model.NewDBEngine(global.DataBaseSetting)
	if err != nil { return err}
	return nil
}
//Redis-go
func setupRedis()error{
	var err error
	global.RedisEngin,err = Redis.Redis(global.RedisSetting)
	if err != nil {
		return err
	}
	return nil
}
//Tracer 链路追踪
func setupTracer()error{
	//这里的6381填写的是运行tracer服务的地址
	jaeger,_,err := tracer.NewJaegerTracer("blog-service","192.168.168.7:6831")
	if err != nil{
		return err
	}
	global.Tracer = jaeger
	return nil
}

//启动参数
func setupFlag() error {
	flag.StringVar(&port,"port","","启动端口")
	flag.StringVar(&runMode,"mode","","启动模式")
	flag.StringVar(&config,"config","configs/","指定要使用的配置文件路径")
	flag.Parse()

	return nil
}