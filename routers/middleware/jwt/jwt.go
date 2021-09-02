package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	. "jixiang/common"
	"net/http"
	"time"
)

var jwtSecret = []byte("jixiang11")

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}
//根据用户名密码生成jwt
func GenerateToken(username, password string) (string, error) {
	fmt.Println("GenerateToken")
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "jixiang",
		},
	}
	// 使用指定的签名方法创建签名对象
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

//jwt中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = SUCCESS
		token := c.GetHeader("Authorization")

		if token == "" {
			code = INVALID_PARAMS
		} else {
			claims, err := ParseToken(token)
			if claims == nil{
				c.JSON(http.StatusUnauthorized, gin.H{
					"code" : 401,
					"msg" : "cookie失效，请点击右上角退出重新登陆",
					"data" : data,
				})
				c.Abort()
				return
			}
			if err != nil {
				code = ERROR
			} else if time.Now().Unix() > claims.ExpiresAt {
				code =ERROR
			}
		}

		if code != SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : GetMsg(code),
				"data" : data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}

//验证jwt是否合法与过期
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
