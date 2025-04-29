package enrollmentHandlers

import (
	"main/database"
	"main/logging"
	"main/models"
	"main/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecordListenerOnProgram(ctx *gin.Context) {
	_, err := tools.CheckParamID(ctx)
	if err != nil {
		return
	}
	var request models.EnrollmentsRequests

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка сервера!"})
		return
	}

	enrollmetns := models.ListenerProgramEducation{
		ID_Listener:         request.ID_Listener,
		ID_ProgramEducation: request.ID_Program,
		StartDate:           request.StartDate,
		EndDate:             request.EndDate,
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		logging.WriteLog(logging.ERROR, "Транзакция не создана")
		return
	}

	if err := tx.Create(&enrollmetns).Error; err != nil {
		logging.WriteLog(logging.ERROR, "Слушатель не записан на курс", enrollmetns.ID_Listener, "-", enrollmetns.ID_ProgramEducation)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка при записи слушателя на курс!"})
		return
	}

	dataListener, err := tools.GetAllListenerData(ctx, enrollmetns.ID_Listener)

	if err != nil {

		logging.WriteLog(logging.ERROR, err, "Не удалось получить данные пользователя")

	}

	dataEducation, err := tools.FindEducationData(ctx, request.ID_Program)

	if err != nil {

		logging.WriteLog(logging.ERROR, err, "Не удалось получить данные об образовании")

	}

	err = tools.CreatePersonalCard(dataListener, dataEducation)
	if err != nil {
		tx.Rollback()
		logging.WriteLog(logging.ERROR, err, "Слушатель не записан на курс")
		return
	}

	tx.Commit()
	logging.WriteLog(logging.DEBUG, "Слушатель", enrollmetns.ID_Listener, "записан на курс")

	ctx.JSON(http.StatusCreated, nil)

}
