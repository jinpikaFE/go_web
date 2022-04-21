package jwt

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"example/pkg/e"
	"example/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		author := c.GetHeader("Authorization")
		// 使用bear token
		token := ""
		if author != "" {
			token = strings.Split(author, " ")[1]
		}
		
		if token == "" {
			code = e.ERROR_NOT_EXIST_TOKEN
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
		        "code" : code,
		        "msg" : e.GetMsg(code),
		        "data" : data,
		    })

		    c.Abort()
		    return
		}

		c.Next()
	}
}