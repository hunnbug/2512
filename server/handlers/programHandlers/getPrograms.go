package programHandlers

import (
	"main/database"
	"main/logging"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPrograms(ctx *gin.Context) {
	const LIMIT_COUNT = 10

	logging.WriteLog(logging.DEBUG, "получен запрос на получение програм обучения")

	type pagePrograms struct {
		CurrentPage int `json:"currentpage"`
	}

	var request pagePrograms

	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка обработки запроса!"})
		logging.WriteLog(logging.ERROR, "Ошибка привязки данных к структуре")
		return
	}

	var programs []models.ProgramEducation

	query := database.DB.Limit(LIMIT_COUNT).Offset((request.CurrentPage - 1) * LIMIT_COUNT).Find(&programs)
	if query.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: query.Error, Message: "Слушатели не найдены"})
		return
	}

	var programDTO = make([]models.ProgramEducationDTO, 0)
	for _, item := range programs {
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
	ctx.JSON(http.StatusOK, programDTO)
}
