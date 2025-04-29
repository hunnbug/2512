package listenerHandlers

import (
	"main/database"
	"main/logging"
	"main/models"
	"main/tools"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteListener(ctx *gin.Context) {
	id, err := tools.CheckParamID(ctx)
	if err != nil {
		return
	}

	var listener models.Listener
	err = tools.CheckRecord(ctx, &listener, "id_listener = ?", id)
	if err != nil {
		return
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		logging.WriteLog(logging.ERROR, "Транзакция не создана")
		return
	}

	tools.DeleteRows(ctx, tx, &models.Listener{}, "id_listener = ?", listener.ID_Listener)

	tools.DeleteRows(ctx, tx, &models.Passport{}, "id_passport = ?", listener.ID_passport)

	tools.DeleteRows(ctx, tx, &models.RegistrationAddress{}, "id_regaddress = ?", listener.ID_regAddress)

	if listener.ID_EducationListener != nil {
		tools.DeleteRows(ctx, tx, &models.EducationListener{}, "id_educationlistener = ?", *listener.ID_EducationListener)
	}

	if listener.ID_PlaceWork != nil {
		tools.DeleteRows(ctx, tx, &models.PlaceWork{}, "id_placework = ?", *listener.ID_PlaceWork)
	}

	if err := tx.Commit().Error; err != nil {
		logging.TxDenied("Удаление не произведено")
	}

	ctx.JSON(http.StatusOK, nil)

	if tx.Error == nil {
		logging.WriteLog(logging.DEBUG, "-----------------")
		logging.WriteLog(logging.DEBUG, "Удалён пользователь и все смежные данные", listener.ID_Listener)
		logging.WriteLog(logging.DEBUG, "-----------------")
	}
}
