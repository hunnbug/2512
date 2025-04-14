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

func UpdateListenerData(ctx *gin.Context) {
	var requestPassport models.Passport
	// проверка на существовании записи в отдельной переменной чтобы избежать перезаписывания модели
	var existsPassport models.Passport

	if err := ctx.ShouldBindJSON(&requestPassport); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка обработки запроса!"})
		logging.WriteLog("Ошибка привязки данных к структуре")
		return
	}

	if err := database.DB.First(&existsPassport, "id_passport = ?", requestPassport.ID_Passport).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Паспорт не найден"})
		logging.WriteLog("Паспорт не найден")
		return
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		logging.WriteLog("Транзакция не создана")
		logging.TxDenied(ctx, tx.Error)
	}

	query := tx.Model(&models.Passport{}).Where("id_passport = ?", requestPassport.ID_Passport).Updates(map[string]interface{}{
		"placebirth":    requestPassport.PlaceBirth,
		"citizenship":   requestPassport.Citizenship,
		"gender":        requestPassport.Gender,
		"seria":         requestPassport.Seria,
		"number":        requestPassport.Number,
		"passportgiven": requestPassport.PassportGiven,
		"dategiven":     requestPassport.DateGiven,
		"code":          requestPassport.Code,
	})
	if query.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: query.Error, Message: "Ошибка обновления записи"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		logging.TxDenied(ctx, err)
	}

	ctx.JSON(http.StatusOK, nil)
	logging.WriteLog("Паспорт - ", requestPassport.ID_Passport, "- изменён")

	//Placework
	var requestPlacework models.PlaceWork
	var existsPlaceWork models.PlaceWork

	if err := ctx.ShouldBindJSON(&requestPlacework); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка обработки запроса!"})
		logging.WriteLog("Ошибка привязки данных к структуре")
		return
	}

	if err := database.DB.First(&existsPlaceWork, "id_placework = ?", requestPlacework.ID_PlaceWork).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Место работы не найдено"})
		logging.WriteLog("Место работы не найдено")
		return
	}

	tx = database.DB.Begin()
	if tx.Error != nil {
		logging.WriteLog("Транзакция не создана")
		logging.TxDenied(ctx, tx.Error)
	}

	query = tx.Model(&models.PlaceWork{}).Where("id_placework = ?", requestPlacework.ID_PlaceWork).Updates(map[string]interface{}{
		"namecompany":        requestPlacework.NameCompany,
		"jobtitle":           requestPlacework.JobTitle,
		"allexperience":      requestPlacework.AllExperience,
		"jobtitleexperience": requestPlacework.JobTitleExpirience,
	})
	if query.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: query.Error, Message: "Ошибка обновления записи"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		logging.TxDenied(ctx, err)
	}

	ctx.JSON(http.StatusOK, nil)
	logging.WriteLog("Место работы - ", requestPlacework.ID_PlaceWork, "- изменёно")

	//RegAddres
	var requestRegAddres models.RegistrationAddress
	var existsRegAddress models.RegistrationAddress

	if err := ctx.ShouldBindJSON(&requestRegAddres); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка обработки запроса!"})
		logging.WriteLog("Ошибка привязки данных к структуре")
		return
	}

	if err := database.DB.First(&existsRegAddress, "id_regaddress = ?", requestRegAddres.ID_regAddress).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Адрес регистрации не найден"})
		logging.WriteLog("Адрес регистрации не найден")
		return
	}

	tx = database.DB.Begin()
	if tx.Error != nil {
		logging.WriteLog("Транзакция не создана")
		logging.TxDenied(ctx, tx.Error)
	}

	query = tx.Model(&models.RegistrationAddress{}).Where("id_regaddress = ?", requestRegAddres.ID_regAddress).Updates(map[string]interface{}{
		"mailindex": requestRegAddres.MailIndex,
		"region":    requestRegAddres.Region,
		"city":      requestRegAddres.City,
		"street":    requestRegAddres.Street,
		"house":     requestRegAddres.House,
		"building":  requestRegAddres.Building,
		"apartment": requestRegAddres.Apartment,
	})
	if query.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: query.Error, Message: "Ошибка обновления записи"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		logging.CheckLogError(err)
	}

	ctx.JSON(http.StatusOK, nil)
	logging.WriteLog("Адрес работы - ", requestRegAddres.ID_regAddress, "- изменён")

	//ListenerEducation
	var requestEducationListener models.EducationListenerDTO
	var existsEducationListener models.EducationListener

	if err := ctx.ShouldBindJSON(&requestEducationListener); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка обработки запроса!"})
		logging.WriteLog("Ошибка привязки данных к структуре")
		return
	}

	if err := database.DB.First(&existsEducationListener, "id_educationlistener = ?", requestEducationListener.ID_EducationListener).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Образование не найдено"})
		logging.WriteLog("Образование работы не найдено")
		return
	}

	tx = database.DB.Begin()
	if tx.Error != nil {
		logging.WriteLog("Транзакция не создана")
		logging.TxDenied(ctx, tx.Error)
	}

	var levelEducation models.LevelEducation

	if err := database.DB.First(&levelEducation, "education = ?", requestEducationListener.LevelEducation).Error; err != nil {
		logging.WriteLog("Уровень образования не найден", requestEducationListener.LevelEducation)
		logging.TxDenied(ctx, err)
		return
	}

	query = tx.Model(&models.EducationListener{}).Where("id_educationlistener = ?", requestEducationListener.ID_EducationListener).Updates(map[string]interface{}{
		"diplomseria":            requestEducationListener.DiplomSeria,
		"diplomnumber":           requestEducationListener.DiplomNumber,
		"dategiven":              requestEducationListener.DateGiven,
		"city":                   requestEducationListener.City,
		"region":                 requestEducationListener.Region,
		"educationalinstitution": requestEducationListener.EducationalInstitution,
		"speciality":             requestEducationListener.Speciality,
		"id_leveleducation":      levelEducation.ID_LevelEducation,
	})
	if query.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: query.Error, Message: "Ошибка обновления записи"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		logging.TxDenied(ctx, err)
	}

	ctx.JSON(http.StatusOK, nil)
	logging.WriteLog("Образование - ", requestEducationListener.ID_EducationListener, "- изменёно")
}
