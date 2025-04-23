package tools

import (
	"main/database"
	"main/logging"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func FindEducationData(ctx *gin.Context, id uuid.UUID) (*models.EducationData, error) {
	var programeducation models.ProgramEducation
	if err := database.DB.Find(&programeducation, "id_programeducation = ?", id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "программа обучения не найдена"})
		logging.WriteLog(logging.ERROR, "программа обучения не найдена")
		return nil, err
	}

	programeducationDTO := models.ProgramEducationDTO{
		NameProfEducation: programeducation.NameProfEducation,
		TimeEducation:     programeducation.TimeEducation,
		IndividualPrice:   programeducation.IndividualPrice,
		GroupPrice:        programeducation.GroupPrice,
		CampusPrice:       programeducation.CampusPrice,
	}

	var division models.DivisionsEducation
	if err := database.DB.Find(&division, "id_divisionseducation = ?", programeducation.ID_DivisionsEducation).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "подразделение не найдено"})
		logging.WriteLog(logging.ERROR, "подразделение не найдено")
		return nil, err
	}

	var typeEducation models.EducationTypes
	if err := database.DB.Find(&typeEducation, "id_educationtype = ?", programeducation.ID_EducationType).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "программа обучения не найдена"})
		logging.WriteLog(logging.ERROR, "программа обучения не найдена")
		return nil, err
	}

	fullEducationData := models.EducationData{
		ProgramEducation: programeducationDTO,
		Division:         division,
		TypeEducation:    typeEducation,
	}

	return &fullEducationData, nil
}
