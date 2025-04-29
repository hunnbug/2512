package listenerHandlers

import (
	"main/logging"
	"main/models"
	"main/tools"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetListenerByID(ctx *gin.Context) {

	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка Parse id"})
		logging.WriteLog(logging.ERROR, "Ошибка Parse id")
		return
	}

	// поиск слушателя чтобы вытащить id через структуру DTO
	data, err := tools.GetAllListenerData(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusForbidden, models.ErrorResponse{Err: err, Message: "Ошибка получения слушателя"})
		return
	}

	ctx.JSON(http.StatusOK, data)

	logging.WriteLog(logging.DEBUG, "Получен пользователь - ", id)

}
