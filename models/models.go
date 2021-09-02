package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"jixiang/global"
	"log"
)
//https://www.tizi365.com/archives/6.html gorm教程
var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
}


func Setup() {

	//todo可改用数据库连接池
	var err error
	db, err = gorm.Open(global.DatabaseSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		global.DatabaseSetting.UserName,
		global.DatabaseSetting.Password,
		global.DatabaseSetting.Host,
		global.DatabaseSetting.DBName))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return global.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
}
