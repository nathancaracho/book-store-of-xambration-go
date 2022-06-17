package routers

import (
	"book-store-of-xambration-go/api/Error"
	"book-store-of-xambration-go/api/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getBooks(ctx *gin.Context) {
	var books []responses.Book
	for i := 0; i < 10; i++ {
		books = append(books, responses.Book{})
	}
	ctx.JSON(http.StatusOK, books)
}

func registerBook(ctx *gin.Context) {
	var book responses.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, Error.ParseFieldError(err))
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

func BooksRegister(route *gin.RouterGroup) {
	route.GET("/books", getBooks)
	route.POST("/book", registerBook)
}
