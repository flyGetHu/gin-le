package route

import (
	"gorm-le/entity"
	"gorm-le/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.GET("/v1/warehouse/page", func(context *gin.Context) {
		page, err := service.Page()
		if err != nil {
			context.JSON(http.StatusOK, entity.Result{}.Error(err.Error()))
			return
		}
		context.JSON(http.StatusOK, entity.Result{}.Ok(page))
	})
}
