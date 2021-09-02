package main

import (
	"jixiang/routers"
	"log"
)
import "jixiang/configs"


func init() {
	err := configs.SetupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}


}

func main(){



	routers.InitRouter().Run(":8089")
	println("吉祥安全扫描器")
}