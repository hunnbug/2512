package programHandlers

import (
	"main/database"
	"main/logging"
	"main/models"
	"main/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
