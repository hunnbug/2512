package handlers

import (
	"main/database"
	"main/logging"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
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

	// program := models.ProgramEducation{
	// 	ID_ProgramEducation: uuid.New(),
	// 	NameProfEducation     :
	// 	TimeEducation         :
	// 	ID_DivisionsEducation :
	// 	IndividualPrice       :
	// 	GroupPrice            :
	// 	CampusPrice           :
	// 	ID_EducationType      :
	// }

	// if err := tx.Create(&program).Error; err != nil {

	// }
}
