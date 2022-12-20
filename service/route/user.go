package route

import (
	"github.com/gin-gonic/gin"
	"log"
	"model"
	"net/http"
)

func userInitRouter(router *gin.RouterGroup) {
	router.GET("/ping", ping)
}

func ping(c *gin.Context) {
	var person model.Person
	err := c.ShouldBindQuery(&person)
	if err != nil {
		log.Println("失败:", err)
	}
	c.JSON(http.StatusOK, person)
}
