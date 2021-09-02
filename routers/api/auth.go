package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	. "jixiang/common"
	"jixiang/models"
	"jixiang/routers/middleware/jwt"
	"log"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := jwt.GenerateToken(username, password)
			if err != nil {
				code = ERROR
			} else {
				data["token"] = token
				code = SUCCESS
			}

		} else {
			code = ERROR
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : jwt.GetMsg(code),
		"data" : data,
	})
}