// package handlers

// import (
// 	"fmt"
// 	"log"
// 	"main/database"
// 	"main/models"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// )

// func LevelEducationCreate(ctx *gin.Context) {
// 	var requset models.LevelEducation

// 	if err := ctx.ShouldBindJSON(&requset); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
// 		return
// 	}

// 	tx := database.DB.Begin()

// 	leveleducation := models.LevelEducation{
// 		ID_LevelEducation: uuid.New(),
// 		Education:         requset.Education,
// 	}

// 	if err := tx.Create(&leveleducation).Error; err != nil {
// 		tx.Rollback()
// 		log.Fatal(err)
// 	}

// 	if err := tx.Commit().Error; err != nil {
// 		log.Fatal(err)
// 	}

// 	ctx.JSON(http.StatusCreated, nil)
// }

// func LevelEducationUpdate(ctx *gin.Context) {

// 	idParam := ctx.Param("id")
// 	id, err := uuid.Parse(idParam)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error:": err})
// 		return
// 	}

// 	var request models.LevelEducation

// 	if err := database.DB.First(&request, "id_leveleducation = ?", id).Error; err != nil {
// 		ctx.JSON(http.StatusNotFound, gin.H{"error:": "Запись не найдена"})
// 		return
// 	}

// 	if err := ctx.ShouldBindJSON(&request); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error:": err})
// 		return
// 	}

// 	result := database.DB.Model(&models.LevelEducation{}).Where("id_leveleducation = ?", id).Update("education", request.Education)
// 	if result.Error != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error:": err})
// 		return
// 	}

// }

// func LevelEducationDelete(ctx *gin.Context) {
// 	idParam := ctx.Param("id")
// 	id, err := uuid.Parse(idParam)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error:": err})
// 		return
// 	}

// 	if err != database.DB.Delete()

// }

package handlers

import (
	"fmt"
)

func nocommit() {
	fmt.Print("Читайте описание")
}
