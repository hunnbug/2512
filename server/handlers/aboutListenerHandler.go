package handlers

import (
	"main/database"
	"main/logging"
	"main/models"
	"main/tools"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AboutListener(ctx *gin.Context) {

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

//
//	UUID каждой изменяемоей сущности приходит в теле запроса.
// 	Для оптимизации работы сделаны методы для переиспользования.
//  Каждый update сделан отдельно для избежания передачи одной гигаструктуры (не всегда же нужно менять всё).
//

func UpdateListenerData(ctx *gin.Context) {
	id, err := tools.CheckParamID(ctx)
	if err != nil {
		return
	}
	var request models.UpdateListenerRequest

	tx := database.DB.Begin()
	if tx.Error != nil {
		logging.WriteLog(logging.ERROR, "Транзакция не создана")
		logging.TxDenied(ctx, tx.Error)
	}

	if err := ctx.Bind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка обработки запроса!"})
		logging.WriteLog(logging.ERROR, "Ошибка привязки данных к структуре")
		return
	}

	var listener models.Listener

	if err := database.DB.First(&listener, "id_listener = ?", id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Слушатель не найден"})
		logging.WriteLog(logging.ERROR, "Слушатель не найден")
		return
	}

	querry := tx.Model(&models.Listener{}).Where("id_listener = ?", id).Updates(map[string]interface{}{
		"firstname":    request.FirstName,
		"secondname":   request.SecondName,
		"middlename":   request.MiddleName,
		"dateofbirth":  request.DateOfBirth,
		"snils":        request.SNILS,
		"contactphone": request.ContactPhone,
		"email":        request.Email,
	})

	if querry.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: querry.Error, Message: "Ошибка обновления записи"})
		return
	}
	logging.WriteLog(logging.DEBUG, "Слушатель - ", id, "- изменён")

	// проверка на существовании записи в отдельной переменной чтобы избежать перезаписывания модели
	var existsPassport models.Passport
	if err := database.DB.First(&existsPassport, "id_passport = ?", request.Passport.ID_Passport).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Паспорт не найден"})
		logging.WriteLog(logging.ERROR, "Паспорт не найден")
		return
	}

	querry = tx.Model(&models.Passport{}).Where("id_passport = ?", request.Passport.ID_Passport).Updates(map[string]interface{}{
		"placebirth":    request.Passport.PlaceBirth,
		"citizenship":   request.Passport.Citizenship,
		"gender":        request.Passport.Gender,
		"seria":         request.Passport.Seria,
		"number":        request.Passport.Number,
		"passportgiven": request.Passport.PassportGiven,
		"dategiven":     request.Passport.DateGiven,
		"code":          request.Passport.Code,
	})
	if querry.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: querry.Error, Message: "Ошибка обновления записи"})
		return
	}

	logging.WriteLog(logging.DEBUG, "Паспорт - ", request.Passport.ID_Passport, "- изменён")

	//Placework
	if request.PlaceWork != (models.PlaceWorkUpdateRequest{}) {
		var existsPlaceWork models.PlaceWork

		if err := database.DB.First(&existsPlaceWork, "id_placework = ?", request.PlaceWork.ID_Placework).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Место работы не найдено"})
			logging.WriteLog(logging.ERROR, "Место работы не найдено")
			return
		}

		querry = tx.Model(&models.PlaceWork{}).Where("id_placework = ?", request.PlaceWork.ID_Placework).Updates(map[string]interface{}{
			"namecompany":        request.PlaceWork.NameCompany,
			"jobtitle":           request.PlaceWork.JobTitle,
			"allexperience":      request.PlaceWork.AllExperience,
			"jobtitleexperience": request.PlaceWork.JobTitleExpirience,
		})
		if querry.Error != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: querry.Error, Message: "Ошибка обновления записи"})
			return
		}

		logging.WriteLog(logging.DEBUG, "Место работы - ", request.PlaceWork.ID_Placework, "- изменёно")
	}

	//RegAddres
	var existsRegAddress models.RegistrationAddress

	if err := database.DB.First(&existsRegAddress, "id_regaddress = ?", request.RegistrationAddress.ID_RegAddress).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Адрес регистрации не найден"})
		logging.WriteLog(logging.ERROR, "Адрес регистрации не найден")
		return
	}

	querry = tx.Model(&models.RegistrationAddress{}).Where("id_regaddress = ?", request.RegistrationAddress.ID_RegAddress).Updates(map[string]interface{}{
		"mailindex": request.RegistrationAddress.MailIndex,
		"region":    request.RegistrationAddress.Region,
		"city":      request.RegistrationAddress.City,
		"street":    request.RegistrationAddress.Street,
		"house":     request.RegistrationAddress.House,
		"building":  request.RegistrationAddress.Building,
		"apartment": request.RegistrationAddress.Apartment,
	})
	if querry.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: querry.Error, Message: "Ошибка обновления записи"})
		return
	}

	logging.WriteLog(logging.DEBUG, "Адрес работы - ", request.RegistrationAddress.ID_RegAddress, "- изменён")

	//ListenerEducation
	if request.EducationListener != (models.EducationListenerUpdateRequest{}) {
		var existsEducationListener models.EducationListener

		if err := database.DB.First(&existsEducationListener, "id_educationlistener = ?", request.EducationListener.ID_EducationListener).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Образование не найдено"})
			logging.WriteLog(logging.ERROR, "Образование работы не найдено")
			return
		}

		// var levelEducation models.LevelEducation

		// if err := database.DB.First(&levelEducation, "id_leveleducation = ?", request.EducationListener.LevelEducation).Error; err != nil {
		// 	logging.WriteLog(logging.ERROR, "Уровень образования не найден", request.EducationListener.LevelEducation)
		// 	logging.TxDenied(ctx, err)
		// 	return
		// }

		querry = tx.Model(&models.EducationListener{}).Where("id_educationlistener = ?", request.EducationListener.ID_EducationListener).Updates(map[string]interface{}{
			"diplomseria":            request.EducationListener.DiplomSeria,
			"diplomnumber":           request.EducationListener.DiplomNumber,
			"dategiven":              request.EducationListener.DateGiven,
			"city":                   request.EducationListener.City,
			"region":                 request.EducationListener.Region,
			"educationalinstitution": request.EducationListener.EducationalInstitution,
			"speciality":             request.EducationListener.Speciality,
			"id_leveleducation":      request.EducationListener.ID_LevelEducation,
		})
		if querry.Error != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: querry.Error, Message: "Ошибка обновления записи"})
			return
		}
		logging.WriteLog(logging.DEBUG, "Образование - ", request.EducationListener.ID_EducationListener, "- изменёно")
	}

	if err := tx.Commit().Error; err != nil {
		logging.TxDenied(ctx, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message:": "записи успешно изменены"})

}
