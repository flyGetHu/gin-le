package route

import (
	"github.com/gin-gonic/gin"
	"log"
)

// 检验账户
var accounts = gin.Accounts{
	"foo":    "bar",
	"austin": "1234",
	"lena":   "hello2",
	"manu":   "4321",
}

func init() {
	router := gin.Default()
	router.Use(gin.Recovery())

	router.MaxMultipartMemory = 8 << 32
	//初始化路由
	authorized := router.Group("/v1", gin.BasicAuth(accounts))
	userInitRouter(authorized)

	log.Println("启动服务器成功:8080")
	err := router.Run(":8080")
	if err != nil {
		log.Println("启动失败:", err)
	}

}
