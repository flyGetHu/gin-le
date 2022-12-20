package main

import (
	_ "gin-le/route"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"path/filepath"
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
}
