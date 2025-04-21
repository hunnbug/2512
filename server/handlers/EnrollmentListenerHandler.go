package handlers

import (
	"main/database"
	"main/logging"
	"main/models"
	"main/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProgramInfo(ctx *gin.Context) {
	id, err := tools.CheckParamID(ctx)
	if err != nil {
		return
	}

	var listener models.Listener
	if err := database.DB.First(&listener, "id_listener = ?", id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Пользователь не найден"})
		logging.WriteLog(logging.ERROR, "Пользователь не найден")
		return
	}

	var divisionseducation []models.DivisionsEducation
	var educationtypes []models.EducationTypes

	querryDivisionEducation := database.DB.Find(&divisionseducation)
	if querryDivisionEducation.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: querryDivisionEducation.Error, Message: "Подразделения не найдены"})
		logging.WriteLog(logging.ERROR, "Подразделения не найдены")
		return
	}

	querryEducationtype := database.DB.Find(&educationtypes)
	if querryEducationtype.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: querryEducationtype.Error, Message: "Типы обучения не найдены"})
		logging.WriteLog(logging.ERROR, "Типы обучения не найдены")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"DivisionEducation": divisionseducation,
		"EducationTypes":    educationtypes,
	})
}

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

	if err := database.DB.Create(&enrollmetns).Error; err != nil {
		logging.WriteLog(logging.ERROR, logging.ERROR, "Слушатель не записан на курс", enrollmetns.ID_Listener, "-", enrollmetns.ID_ProgramEducation)
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка при записи слушателя на курс!"})
		return
	}

	ctx.JSON(http.StatusCreated, nil)

}
