package configs

import (
	"jixiang/global"
	"jixiang/models"
	"strings"
	"time"
)

func SetupSetting() error {
	//支持从多个路径读取configs文件，配置初始化
	s, err := NewSetting(strings.Split("configs/", ",")...)
	if err != nil {
		return err
	}
	//此刻已经讲config文件中配置读到global中的结构体
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Masscan", &global.MasscanSetting)
	if err != nil {
		return err
	}
	global.AppSetting.DefaultContextTimeout *= time.Second  //C *= A 等于 C = C * A
	//global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
//数据库初始化
	models.Setup()

	return nil
}
