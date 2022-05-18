package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-programming/blog-service/docs"
	"github.com/go-programming/blog-service/global"
	"github.com/go-programming/blog-service/internal/model"
	"github.com/go-programming/blog-service/internal/routers"
	"github.com/go-programming/blog-service/pkg/logger"
	setting "github.com/go-programming/blog-service/pkg/setting"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

func init() {
	// init setting in configs
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
	// init database
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setDBEngine err: %v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}
}
func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

//use github.com/gin-gonic/gin@v1.6.3
// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService
func main() {
	//log.Printf("%v", global.ServerSetting)
	//log.Printf("%v", global.AppSetting)
	global.Logger.Infof("%s: go-programming/%s", "gjw", "blog-service")

	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		return
	}

}
