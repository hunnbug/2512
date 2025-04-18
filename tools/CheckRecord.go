package tools

import (
	"main/database"
	"main/logging"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CheckRecord(ctx *gin.Context, existsmodel interface{}, querryWhere string, id uuid.UUID) error {
	if err := database.DB.First(&existsmodel, querryWhere, id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: querryWhere + "не найден"})
		logging.WriteLog(logging.ERROR, querryWhere+"не найден")
		return err
	}
	return nil
}
