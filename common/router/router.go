package router

import (
	"fmt"
	"net/http"
	"newe-serve/app/routers"
	"newe-serve/common/nelog"
	"newe-serve/common/setting"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// 你的自定义格式
		str := fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
		nelog.Trace(str)
		nelog.Trace("请求body:", param.Keys)
		nelog.Trace("-----------------------------------------------------------------------")
		return str
	}))
	r.Use(gin.Recovery())
	//r.LoadHTMLGlob("wwwroot/*.html")
	r.StaticFS("/upload/images", http.Dir(setting.Imgcfg.ImageSavePath)) //文件目录

	r.Use(Cors())
	routers.Router(r)

	//r.LoadHTMLGlob("template/**/*")
	// //-------------API文档----------------
	// r.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "NAME_OF_ENV_VARIABLE"))

	return r
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {

		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		headerKeys = append(headerKeys, "Token")
		headerKeys = append(headerKeys, "AccSn")
		headerKeys = append(headerKeys, "Content-Type")
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin,x-requested-with, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {

			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", headerStr)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			c.Set("content-type", "application/json")
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}
