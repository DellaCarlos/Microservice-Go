package controllers

import (
	"microservice/internal/repositories"
	use_cases "microservice/internal/use-cases"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListCategories(ctx *gin.Context, repository repositories.ICategoryRepository) {
	useCase := use_cases.NewListCategoriesUseCase(repository)

	categories, err := useCase.Execute()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"sucess": false,
				"error":  err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"sucess": true, "categories": categories})
}
