package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"jixiang/models"
	"jixiang/routers/middleware/jwt"
	"net/http"
	"strconv"
)

func GetIplist(c *gin.Context) {
	protocol := c.Query("protocol")
	ip := c.Query("ip")
	port := c.Query("port")
	title := c.Query("title")
	finger := c.Query("finger")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if protocol != "" {
		maps["protocol"] = protocol
	}

	if ip != "" {
		maps["ip"] = ip
	}

	if port != "" {
		maps["port"] = port
	}

	if title != "" {
		maps["title"] = title
	}

	if finger != "" {
		maps["finger"] = finger
	}
	code := 200

	data["lists"] = models.GetIplist(GetPage(c), 10, maps)
	data["total"] = models.GetIplistTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": jwt.GetMsg(code),
		"data": data,
	})
}


func GetPage(c *gin.Context) int {
	result := 0
	pagesizetmp := c.Query("pagesize")
	pagesize, _ := strconv.Atoi(pagesizetmp)

	page, _ := com.StrTo(c.Query("pagenum")).Int()
	if page > 0 {
		result = (page - 1) * pagesize
	}

	return result
}
