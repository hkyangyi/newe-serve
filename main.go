package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"newe-serve/common/db"
	"newe-serve/common/redis"
	"newe-serve/common/router"
	"newe-serve/common/setting"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	appPath              string
	Server               *http.Server
	flagServiceName      = flag.String("service-name", "NeweServer", "Set service name")
	flagServiceDesc      = flag.String("service-desc", "NeweServer", "Set service description")
	flagServiceInstall   = flag.Bool("install", false, "Install service")
	flagServiceUninstall = flag.Bool("remove", false, "Remove service")
	flagServiceStart     = flag.Bool("start", false, "Start service")
	flagServiceStop      = flag.Bool("stop", false, "Stop service")
)

func main() {
	setting.SetUp()
	redis.Setup()
	db.Setup()
	webserve()
	fmt.Println("1111111111111")
}

func webserve() {
	gin.SetMode(setting.SYS.RunMode)
	routersInit := router.InitRouter()
	//routersInit.Run(":1212")
	readTimeout := time.Duration(int64(setting.SYS.ReadTimeout)) * time.Second
	writeTimeout := time.Duration(int64(setting.SYS.WriteTimeout)) * time.Second
	endPoint := fmt.Sprintf(":%d", setting.SYS.HttpPort)
	maxHeaderBytes := 1 << 20
	Server = &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	//go Gootherexe()
	Server.ListenAndServe()
}
