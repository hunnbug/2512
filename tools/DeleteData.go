package tools

import (
	"main/logging"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DeleteRows(ctx *gin.Context, tx *gorm.DB, model interface{}, idString string, id uuid.UUID) error {
	querry := tx.Where(idString, id).Delete(model)
	if err := querry.Error; err != nil {
		ErrorDelete(ctx, err, tx, idString)
		return tx.Error
	}
	if querry.RowsAffected == 0 {
		ErrorDelete(ctx, querry.Error, tx, idString)
		return tx.Error
	}
	return nil
}

func ErrorDelete(ctx *gin.Context, err error, tx *gorm.DB, idString string) {
	tx.Rollback()
	ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: idString + "не найден"})
	logging.WriteLog(idString, "не найден")
	logging.TxDenied(err)
	return
}
