package handlers

import (
	"main/database"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllEducationTypes(ctx *gin.Context) {
	var fullEducationType []models.EducationTypes
	querryLevelEducation := database.DB.Find(&fullEducationType)

	if querryLevelEducation.Error != nil {

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: querryLevelEducation.Error, Message: "Типы обучения не найдены"})

		return
	}

	ctx.JSON(200, gin.H{
		"LevelsEducation": fullEducationType,
	})
}
