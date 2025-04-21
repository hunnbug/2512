package handlers

import (
	"main/database"
	"main/logging"
	"main/models"
	"main/tools"
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

func DeleteProgram(ctx *gin.Context) {
	id, err := tools.CheckParamID(ctx)
	if err != nil {
		return
	}

	var model models.ProgramEducation

	err = tools.CheckRecord(ctx, &model, "id_programeducation = ?", id)
	if err != nil {
		return
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		logging.WriteLog(logging.ERROR, "Транзакция не создана")
		return
	}

	tools.DeleteRows(ctx, tx, &model, "id_programeducation = ?", id)

	if err := tx.Commit().Error; err != nil {
		logging.TxDenied(ctx, id)
		return
	}
	ctx.JSON(http.StatusOK, nil)
	logging.WriteLog(logging.DEBUG, "Удалена программа - ", id)
}

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

func ReadProgram(ctx *gin.Context) {
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
