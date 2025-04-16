package tools

import (
	"main/logging"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CheckParamID(ctx *gin.Context) (uuid.UUID, error) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка Parse id"})
		logging.WriteLog("Ошибка Parse id")
		return uuid.Nil, err
	}
	return id, nil
}
