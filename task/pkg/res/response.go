package res

import (
	"net/http"
	"task/pkg/e"

	"github.com/gin-gonic/gin"
)

// Response 基础序列化器
type Response struct {
	Status uint        `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

//DataList 带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

//TokenData 带有token的Data结构
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

// 返回200 自定义code data
func Ok(ctx *gin.Context, msgCode int, data interface{}) {
	ctx.JSON(http.StatusOK, ginH(msgCode, data))
}

// 无权限err
func Unauthorized(ctx *gin.Context, msgCode int) {
	ctx.JSON(http.StatusUnauthorized, ginH(msgCode, nil))
}

// 内部err
func InternalError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, ginH(e.ERROR, nil))
}

// 禁止访问err
func ForbiddenError(ctx *gin.Context, msgCode int) {
	ctx.JSON(http.StatusForbidden, ginH(msgCode, nil))
}

// 自定义 err
func Error(ctx *gin.Context, httpCode, msgCode int) {
	ctx.JSON(httpCode, ginH(msgCode, nil))
}

func ginH(msgCode int, data interface{}) gin.H {
	return gin.H{
		"code": msgCode,
		"msg":  e.GetMsg(uint(msgCode)),
		"data": data,
	}
}
