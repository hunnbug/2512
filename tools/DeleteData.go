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
		CheckDeleteError(ctx, err, tx, idString)
	}
	if querry.RowsAffected == 0 {
		CheckDeleteError(ctx, querry.Error, tx, idString)
		return tx.Error
	}
	return nil
}

func CheckDeleteError(ctx *gin.Context, err error, tx *gorm.DB, id string) {
	tx.Rollback()
	ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: id})
	logging.WriteLog(id, "не найден")
	logging.TxDenied(err)
	return
}
