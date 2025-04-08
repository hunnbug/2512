package handlers

import (
	"main/database"
	"main/logging"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//
//	UUID каждой изменяемоей сущности приходит в теле запроса.
// 	Для оптимизации работы сделаны методы для переиспользования.
//  Каждый update сделан отдельно для избежания передачи одной гигаструктуры (не всегда же нужно менять всё).
//

// Обновление паспорта
func UpdateListenersPassport(ctx *gin.Context) {

	var request models.Passport
	// проверка на существовании записи в отдельной переменной чтобы избежать перезаписывания модели
	var existsPassport models.Passport

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка обработки запроса!"})
		logging.WriteLog("Ошибка привязки данных к структуре")
		return
	}

	if err := database.DB.First(&existsPassport, "id_passport = ?", request.ID_Passport).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Паспорт не найден"})
		logging.WriteLog("Паспорт не найден")
		return
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		err := logging.WriteLog("Транзакция не создана")
		logging.CheckLogError(err)
	}

	query := tx.Model(&models.Passport{}).Where("id_passport = ?", request.ID_Passport).Updates(map[string]interface{}{
		"placebirth":    request.PlaceBirth,
		"citizenship":   request.Citizenship,
		"gender":        request.Gender,
		"seria":         request.Seria,
		"number":        request.Number,
		"passportgiven": request.PassportGiven,
		"dategiven":     request.DateGiven,
		"code":          request.Code,
	})
	if query.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: query.Error, Message: "Ошибка обновления записи"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		logging.CheckLogError(err)
	}

	ctx.JSON(http.StatusOK, nil)
}

// Обновление места работы
func UpdateListenersPlaceWork(ctx *gin.Context) {

	var request models.PlaceWork
	var existsPlaceWork models.PlaceWork

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка обработки запроса!"})
		logging.WriteLog("Ошибка привязки данных к структуре")
		return
	}

	if err := database.DB.First(&existsPlaceWork, "id_placework = ?", request.ID_PlaceWork).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Место работы не найдено"})
		logging.WriteLog("Место работы не найдено")
		return
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		err := logging.WriteLog("Транзакция не создана")
		logging.CheckLogError(err)
	}

	query := tx.Model(&models.PlaceWork{}).Where("id_placework = ?", request.ID_PlaceWork).Updates(map[string]interface{}{
		"namecompany":        request.NameCompany,
		"jobtitle":           request.JobTitle,
		"allexperience":      request.AllExperience,
		"jobtitleexperience": request.JobTitleExpirience,
	})
	if query.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: query.Error, Message: "Ошибка обновления записи"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		logging.CheckLogError(err)
	}

	ctx.JSON(http.StatusOK, nil)
}

// Обновление адреса регистрации
func UpdateListenersRegAddress(ctx *gin.Context) {

	var request models.RegistrationAddress
	var existsRegAddress models.RegistrationAddress

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка обработки запроса!"})
		logging.WriteLog("Ошибка привязки данных к структуре")
		return
	}

	if err := database.DB.First(&existsRegAddress, "id_regaddress = ?", request.ID_regAddress).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Место работы не найдено"})
		logging.WriteLog("Место работы не найдено")
		return
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		err := logging.WriteLog("Транзакция не создана")
		logging.CheckLogError(err)
	}

	query := tx.Model(&models.RegistrationAddress{}).Where("id_regaddress = ?", request.ID_regAddress).Updates(map[string]interface{}{
		"mailindex": request.MailIndex,
		"region":    request.Region,
		"city":      request.City,
		"street":    request.Street,
		"house":     request.House,
		"building":  request.Building,
		"apartment": request.Apartment,
	})
	if query.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: query.Error, Message: "Ошибка обновления записи"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		logging.CheckLogError(err)
	}

	ctx.JSON(http.StatusOK, nil)
}

// Обновление образования слушателя
func UpdateListenersEducation(ctx *gin.Context) {

	var request models.EducationListenerDTO
	var existsEducationListener models.EducationListener

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка обработки запроса!"})
		logging.WriteLog("Ошибка привязки данных к структуре")
		return
	}

	if err := database.DB.First(&existsEducationListener, "id_educationlistener = ?", request.ID_EducationListener).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Образование не найдено"})
		logging.WriteLog("Образование работы не найдено")
		return
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		err := logging.WriteLog("Транзакция не создана")
		logging.CheckLogError(err)
	}

	var levelEducation models.LevelEducation

	if err := database.DB.First(&levelEducation, "education = ?", request.LevelEducation).Error; err != nil {
		err = logging.WriteLog("Уровень образования не найден", request.LevelEducation)
		logging.CheckLogError(err)
		return
	}

	query := tx.Model(&models.EducationListener{}).Where("id_educationlistener = ?", request.ID_EducationListener).Updates(map[string]interface{}{
		"diplomseria":            request.DiplomSeria,
		"diplomnumber":           request.DiplomNumber,
		"dategiven":              request.DateGiven,
		"city":                   request.City,
		"region":                 request.Region,
		"educationalinstitution": request.EducationalInstitution,
		"speciality":             request.Speciality,
		"id_leveleducation":      levelEducation.ID_LevelEducation,
	})
	if query.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: query.Error, Message: "Ошибка обновления записи"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		logging.CheckLogError(err)
	}

	ctx.JSON(http.StatusOK, nil)
}
