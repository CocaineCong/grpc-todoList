package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/CocaineCong/grpc-todolist/consts"
	"github.com/CocaineCong/grpc-todolist/pkg/e"
	"github.com/CocaineCong/grpc-todolist/pkg/util/jwt"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = 200
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 404
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(uint(code)),
				"data":   data,
			})
			c.Abort()
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			code = e.ErrorAuthCheckTokenFail
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = e.ErrorAuthCheckTokenTimeout
		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(uint(code)),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Set(consts.UserIdKey, claims.UserID)
		c.Next()
	}
}
