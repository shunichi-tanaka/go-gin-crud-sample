package main

import (
	"fmt"
	"go-gin-crud-sample/db"
	"go-gin-crud-sample/pkg/setting"
	"go-gin-crud-sample/routers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.SetUp()
	db.Init()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening ")
	server.ListenAndServe()
}
