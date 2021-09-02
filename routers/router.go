package routers

import (
	"github.com/gin-gonic/gin"
	"jixiang/routers/api"
	"jixiang/routers/api/v1"
	"jixiang/routers/middleware/jwt"
	"jixiang/routers/middleware/cors"
)

func InitRouter() *gin.Engine {
	r := gin.New()
//中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Cors())
//生成jwt
	r.GET("/api/v1/auth", api.GetAuth)
	//修改密码
	r.POST("/api/v1/auth/passedit", v1.EditPass)


	apiv1 := r.Group("/api/v1")
//需要分组url，header中需要有生成的Authorization：
	apiv1.Use(jwt.JWT())
	{


		//apiv1.GET("/downtasklog/:id", v1.DownTaskLog)
		// masscan资产任务列表
		apiv1.GET("/masstasks", v1.GetIplist)
		// ip资产列表搜索
		apiv1.GET("/masstask", v1.GetIplist)




		apiv1.POST("/addcron",v1.AddTask)


	}

	return r

}

