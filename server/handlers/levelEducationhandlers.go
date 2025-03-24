package handlers

import (
	"log"
	"main/database"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LevelEducationCreate(ctx *gin.Context) {
	var requset models.LevelEducation

	if err := ctx.ShouldBindJSON(&requset); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	tx := database.DB.Begin()

	leveleducation := models.LevelEducation{
		ID_LevelEducation: uuid.New(),
		Education:         requset.Education,
	}

	if err := tx.Create(&leveleducation).Error; err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	if err := tx.Commit().Error; err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusCreated, nil)
}

func LevelEducationUpdate(ctx *gin.Context) {

	// idParam := ctx.Param("id")
	// id, err := uuid.Parse(idParam)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error:": err})
	// 	return
	// }

	var request models.LevelEducation

	updates := make(map[string]interface{})

	if request.Education != "" {
		updates["education"] = request.Education
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error:": err})
		return
	}

}
