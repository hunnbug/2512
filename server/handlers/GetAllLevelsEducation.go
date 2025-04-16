package handlers

import (
	"main/database"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllLevelsEducation(ctx *gin.Context) {
	var fullLevelEducation []models.LevelEducation
	querryLevelEducation := database.DB.Find(&fullLevelEducation)

	if querryLevelEducation.Error != nil {

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: querryLevelEducation.Error, Message: "Слушатели не найдены"})

		return
	}

	ctx.JSON(200, gin.H{
		"LevelsEducation": fullLevelEducation,
	})
}
