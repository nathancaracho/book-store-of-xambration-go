package controllers

import (
	"book-store-of-xambration-go/api/dto"
	"book-store-of-xambration-go/api/error"
	"book-store-of-xambration-go/domain/models"
	"book-store-of-xambration-go/infrastructure"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getBooks(ctx *gin.Context) {
	var books []models.Book

	db, err := infrastructure.GetConnection()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, error.ApiError{Message: "Some error happen when try connect in database"})
	}
	db.Find(&books)
	ctx.JSON(http.StatusOK, books)

}

func createBook(ctx *gin.Context) {
	var book dto.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, error.ParseFieldError(err))
		return
	}
	db, err := infrastructure.GetConnection()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, error.ApiError{Message: "Some error happen when try connect in database"})
	}
	modelBook := models.Book{
		Name:             book.Name,
		Description:      book.Description,
		BriefDescription: book.BriefDescription,
		Value:            book.Value,
	}
	result := db.Create(&modelBook)
	if result.Error != nil && result.RowsAffected != 1 {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, error.ApiError{Message: result.Error.Error()})

	}
	ctx.JSON(http.StatusCreated, modelBook)
}

func BooksV1Register(router *gin.RouterGroup) {
	router.GET("/books", getBooks)
	router.POST("/book", createBook)
}
