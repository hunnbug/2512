package handlers

import (
	"main/database"
	"main/logging"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetProgramInfo(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка Parse id"})
		logging.WriteLog("Ошибка Parse id")
		return
	}

	var listener models.Listener
	if err := database.DB.First(&listener, "id_listener = ?", id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Пользователь не найден"})
		logging.WriteLog("Пользователь не найден")
		return
	}

	var divisionseducation []models.DivisionsEducation
	var educationtypes []models.EducationTypes

	querryDivisionEducation := database.DB.Find(&divisionseducation)
	if querryDivisionEducation.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: querryDivisionEducation.Error, Message: "Подразделения не найдены"})
		logging.WriteLog("Подразделения не найдены")
		return
	}

	querryEducationtype := database.DB.Find(&educationtypes)
	if querryEducationtype.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: querryEducationtype.Error, Message: "Типы обучения не найдены"})
		logging.WriteLog("Типы обучения не найдены")
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
		logging.WriteLog("Ошибка привязки подразделения: ", err)
		return
	}

	var programeducation []models.ProgramEducation
	querryProgramEducation := database.DB.Where("id_divisionseducation = ?", divions.ID_DivisionsEducation).Find(&programeducation)
	if querryProgramEducation.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: querryProgramEducation.Error, Message: "Програмы обучения не найдены"})
		logging.WriteLog("Програмы обучения не найдены")
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
