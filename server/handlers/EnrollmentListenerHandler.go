package handlers

import (
	"main/database"
	"main/logging"
	"main/models"
	"main/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SelectProgramEducation(ctx *gin.Context) {

	var divions models.DivisionsEducationRequests
	err := ctx.BindJSON(&divions)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка привязки подразделения"})
		logging.WriteLog(logging.ERROR, "Ошибка привязки подразделения: ", err)
		return
	}

	var programeducation []models.ProgramEducation
	querryProgramEducation := database.DB.Where("id_divisionseducation = ?", divions.ID_DivisionsEducation).Find(&programeducation)
	if querryProgramEducation.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: querryProgramEducation.Error, Message: "Програмы обучения не найдены"})
		logging.WriteLog(logging.ERROR, "Програмы обучения не найдены")
		return
	}

	var programDTO = make([]models.ProgramEducationDTO, 0)
	for _, item := range programeducation {
		dto := models.ProgramEducationDTO{
			ID_ProgramEducation:   item.ID_ProgramEducation,
			NameProfEducation:     item.NameProfEducation,
			TimeEducation:         item.TimeEducation,
			IndividualPrice:       item.IndividualPrice,
			GroupPrice:            item.GroupPrice,
			CampusPrice:           item.CampusPrice,
			ID_DivisionsEducation: item.ID_DivisionsEducation,
			ID_EducationType:      item.ID_EducationType,
		}
		programDTO = append(programDTO, dto)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"ProgramEducation": programDTO,
	})
}

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
