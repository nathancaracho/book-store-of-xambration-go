package infrastructure

import (
	"book-store-of-xambration-go/api/routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouterSetup() {
	router := gin.Default()

	apiV1 := router.Group("/v1/api")

	routers.BooksRegister(apiV1)
	addHealthCheck(router)

	router.Run(":8080")
}

func addHealthCheck(router *gin.Engine) {
	router.GET("/health", func(context *gin.Context) {
		context.String(http.StatusOK, "I`m ok thanks!!!")
	})
}
