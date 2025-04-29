package programHandlers

import (
	"main/database"
	"main/logging"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateProgram(ctx *gin.Context) {
	var request models.ProgramEducationRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка сервера!"})
		return
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		logging.WriteLog(logging.ERROR, "Транзакция не создана")
		return
	}

	program := models.ProgramEducation{
		ID_ProgramEducation:   uuid.New(),
		NameProfEducation:     request.NameProfEducation,
		TimeEducation:         request.TimeEducation,
		IndividualPrice:       request.IndividualPrice,
		GroupPrice:            request.GroupPrice,
		CampusPrice:           request.CampusPrice,
		ID_EducationType:      request.ID_EducationType,
		ID_DivisionsEducation: request.ID_DivisionsEducation,
	}

	if err := tx.Create(&program).Error; err != nil {
		tx.Rollback()
		logging.WriteLog(logging.ERROR, "Программа - ", program.NameProfEducation, "не создана")
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка при создании программы обучения!"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		logging.TxDenied(ctx, program.ID_ProgramEducation)
		return
	}
	logging.WriteLog(logging.DEBUG, "Создана программа - ", program.NameProfEducation)
}
