package app

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Result    interface{} `json:"result"`
	Success   bool        `json:"success"`
	Timestamp int64       `json:"timestamp"`
}

//登陆失败
func (g *Gin) LoginError(err error) {
	g.C.JSON(401, Response{
		Code:      10000,
		Message:   err.Error(),
		Result:    nil,
		Success:   false,
		Timestamp: time.Now().Unix(),
	})
	return
}

func (g *Gin) Success(data interface{}) {
	g.C.JSON(200, Response{
		Code:      200,
		Message:   "ok",
		Result:    data,
		Success:   true,
		Timestamp: time.Now().Unix(),
	})
	return
}

func (g *Gin) Error(err error) {
	g.C.JSON(200, Response{
		Code:      400,
		Message:   err.Error(),
		Result:    nil,
		Success:   false,
		Timestamp: time.Now().Unix(),
	})
	return
}

//公共日志写入
func AdminLogWirte(url string, formdata interface{}, result Response) {

}
