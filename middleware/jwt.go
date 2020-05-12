package middleware

import (
	"dzero/godb/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWTAuth() gin.HandlerFunc {

	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = util.SUCCESS
		token := c.Request.Header.Get("Token")
		if token == "" {
			code = util.INVALID_PARAMS
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = util.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = util.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != util.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  util.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}

}
