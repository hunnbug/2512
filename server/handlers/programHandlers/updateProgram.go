package programHandlers

import (
	"main/database"
	"main/logging"
	"main/models"
	"main/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateProgram(ctx *gin.Context) {
	id, err := tools.CheckParamID(ctx)
	if err != nil {
		return
	}

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

	var model models.ProgramEducation

	querry := tx.Model(&model).Where("id_programeducation = ?", id).Updates(map[string]interface{}{
		"nameprofeducation":     request.NameProfEducation,
		"timeeducation":         request.TimeEducation,
		"individualprice":       request.IndividualPrice,
		"groupprice":            request.GroupPrice,
		"campusprice":           request.CampusPrice,
		"id_divisionseducation": request.ID_DivisionsEducation,
		"id_educationtype":      request.ID_EducationType,
	})

	if querry.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: querry.Error, Message: "Ошибка обновления записи"})
		return
	}
	if err := tx.Commit().Error; err != nil {
		logging.TxDenied(ctx, id)
		return
	}
	ctx.JSON(http.StatusOK, nil)
	logging.WriteLog(logging.DEBUG, "Обновлена программа - ", id)
}
