package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	. "jixiang/common"
	"jixiang/models"
	"jixiang/routers/middleware/jwt"
	"net/http"
)

func EditPass(c *gin.Context) {

	username := c.PostForm("username")
	oldpass := c.PostForm("oldpass")
	newpass := c.PostForm("newpass")
	newpass2 := c.PostForm("newpass2")

	code := INVALID_DIFFPASS
	if newpass != newpass2{
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  "两次密码不一致",
			"data": make(map[string]string),
		})
		return
	}

	valid := validation.Validation{}

	code = INVALID_PASS
	if ! valid.HasErrors() {
		isExist := models.CheckAuth(username, oldpass)
		if isExist {
			data := make(map[string]interface{})
			data["password"] = newpass
			models.EditAuth(username,data)
			code = SUCCESS
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  jwt.GetMsg(code),
		"data": make(map[string]string),
	})
}
