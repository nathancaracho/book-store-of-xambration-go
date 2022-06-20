package api

import (
	"book-store-of-xambration-go/api/controllers"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	registerV1(engine)
}
func registerV1(engine *gin.Engine) {
	apiV1 := engine.Group("/v1/api")
	controllers.BooksV1Register(apiV1)
}
