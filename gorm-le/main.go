package main

import (
	"gorm-le/route"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	gin.DisableConsoleColor()
	//gin.ForceConsoleColor()
	logPathRoot := "./logs/"
	if _, err := os.Stat(logPathRoot); os.IsNotExist(err) {
		err := os.MkdirAll(logPathRoot, 0777)
		if err != nil {
			log.Println("创建日志文件失败:", err)
		}
	}
	f, err := os.Create(filepath.Join(logPathRoot, "gin_http.log"))
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.Use(gin.Recovery())

	router.MaxMultipartMemory = 8 << 32
	//初始化路由
	route.Init(router)
	log.Println("启动服务器成功:8081")
	err = router.Run(":8081")
	if err != nil {
		log.Println("启动失败:", err)
	}
}
