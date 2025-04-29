package enrollmentHandlers

import (
	"main/database"
	"main/logging"
	"main/models"
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
