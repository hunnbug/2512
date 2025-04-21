package handlers

import (
	"fmt"
	"main/database"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllLevelsEducation(ctx *gin.Context) {
	var fullLevelEducation []models.LevelEducation
	querryLevelEducation := database.DB.Find(&fullLevelEducation)

	if querryLevelEducation.Error != nil {

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: querryLevelEducation.Error, Message: "Уровни обучения не найдены"})

		return
	}

	fmt.Println(fullLevelEducation)

	ctx.JSON(200, gin.H{
		"LevelsEducation": fullLevelEducation,
	})
}
