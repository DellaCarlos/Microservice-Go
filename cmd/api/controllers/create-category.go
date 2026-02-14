package controllers

import (
	use_cases "microservice/internal/use-cases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(ctx *gin.Context) {
	var body createCategoryInput

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"sucess": false,
				"erro":   err.Error(),
			})
		return
	}

	useCase := use_cases.NewCreateCategoryUseCase()
	err := useCase.Execute(body.Name)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"sucess": false,
				"error":  err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"sucess": true})
}
