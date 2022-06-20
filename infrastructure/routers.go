package infrastructure

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func RouterConfig(register func(engine *gin.Engine)) {
	engine := gin.Default()

	register(engine)
	addHealthCheck(engine)

	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err.Error())
	}

}

func addHealthCheck(router *gin.Engine) {
	router.GET("/health", func(context *gin.Context) {
		context.String(http.StatusOK, "I`m ok thanks!!!")
	})
}
