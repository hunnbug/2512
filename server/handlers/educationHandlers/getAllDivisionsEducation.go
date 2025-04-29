package educationHandlers

import (
	"main/database"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllDivisions(ctx *gin.Context) {
	var fullDivisions []models.DivisionsEducation
	querryDivisionsEducation := database.DB.Find(&fullDivisions)

	if querryDivisionsEducation.Error != nil {

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: querryDivisionsEducation.Error, Message: "Подразделения не найдены"})

		return
	}

	ctx.JSON(200, gin.H{
		"LevelsEducation": fullDivisions,
	})
}
