package handlers

import (
	"main/database"
	"main/logging"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AboutListener(ctx *gin.Context) {

	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка Parse id"})
		logging.WriteLog("Ошибка Parse id")
		return
	}

	// поиск слушателя чтобы вытащить id через структуру DTO

	//транзакции на get делать не стоит всё равно выход из метода при ненаходе одного из
	var listener models.Listener

	querry := database.DB.First(&listener, "id_listener = ?", id)
	if querry.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Пользователь не найден"})
		logging.WriteLog("Пользователь не найден")
		return
	}
	logging.WriteLog("Авторизация пользователя")

	logging.CheckLogError(err)
	responseUUID := models.ListenerIDDTO{
		ID_Passport:          listener.ID_passport,
		ID_RegAddress:        listener.ID_regAddress,
		ID_EducationListener: listener.ID_EducationListener,
		ID_PlaceWork:         listener.ID_PlaceWork,
	}

	// Все другие таблицы в которых поиск идёт по id из предыдущей структуры
	var passport models.Passport
	var regAddress models.RegistrationAddress
	var educationListener models.EducationListener
	var placework models.PlaceWork

	if err := database.DB.First(&passport, "id_passport = ?", responseUUID.ID_Passport).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Паспорт не найден"})
		logging.WriteLog("Паспорт не найден")
		return
	}

	if err := database.DB.First(&regAddress, "id_regaddress = ?", responseUUID.ID_RegAddress).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Адрес не найден"})
		logging.WriteLog("Адрес не найден")
		return
	}

	if err := database.DB.First(&educationListener, "id_educationlistener = ?", responseUUID.ID_EducationListener).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Образование не найдено"})
		logging.WriteLog("Образование не найдено")
		return
	}

	if err := database.DB.First(&placework, "id_placework = ?", responseUUID.ID_PlaceWork).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Место работы не найдено"})
		logging.WriteLog("Место работы не найдено")
		return
	}

	responseListener := models.ListenerDTO{
		FirstName:    listener.FirstName,
		SecondName:   listener.SecondName,
		MiddleName:   listener.MiddleName,
		DateOfBirth:  listener.DateOfBirth,
		SNILS:        listener.SNILS,
		ContactPhone: listener.ContactPhone,
		Email:        listener.Email,
	}
	responsePassport := models.PassportDTO{
		PlaceBirth:    passport.PlaceBirth,
		Citizenship:   passport.Citizenship,
		Gender:        passport.Gender,
		Seria:         passport.Seria,
		Number:        passport.Number,
		PassportGiven: passport.PassportGiven,
		DateGiven:     passport.DateGiven,
		Code:          passport.Code,
	}

	responseRegAddress := models.RegistrationAddressDTO{
		MailIndex: regAddress.MailIndex,
		Region:    regAddress.Region,
		City:      regAddress.City,
		Street:    regAddress.Street,
		House:     regAddress.House,
		Building:  regAddress.Building,
		Apartment: regAddress.Apartment,
	}

	// вложеная связь
	var leveleducation models.LevelEducation
	if err := database.DB.Find(&leveleducation, "id_leveleducation", educationListener.ID_LevelEducation).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Уровень образования не найден"})
		logging.WriteLog("Уровень образования не найден")
		return
	}
	responseEducationListener := models.EducationListenerDTO{
		DiplomSeria:            educationListener.DiplomSeria,
		DiplomNumber:           educationListener.DiplomNumber,
		DateGiven:              educationListener.DateGiven,
		City:                   educationListener.City,
		Region:                 educationListener.Region,
		EducationalInstitution: educationListener.EducationalInstitution,
		Speciality:             educationListener.Speciality,
		LevelEducation:         leveleducation.Education,
	}

	responsePlacework := models.PlaceWorkDTO{
		NameCompany:        placework.NameCompany,
		JobTitle:           placework.JobTitle,
		AllExperience:      placework.AllExperience,
		JobTitleExpirience: placework.JobTitleExpirience,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"listener":          responseListener,
		"passport":          responsePassport,
		"regAddress":        responseRegAddress,
		"educationListener": responseEducationListener,
		"placeWork":         responsePlacework,
		"responseuuid":      responseUUID,
	})

}

//
//	UUID каждой изменяемоей сущности приходит в теле запроса.
// 	Для оптимизации работы сделаны методы для переиспользования.
//  Каждый update сделан отдельно для избежания передачи одной гигаструктуры (не всегда же нужно менять всё).
//

func UpdateListenerData(ctx *gin.Context) {
	var request models.UpdateListenerRequest
	// проверка на существовании записи в отдельной переменной чтобы избежать перезаписывания модели
	var existsPassport models.Passport
	tx := database.DB.Begin()
	if tx.Error != nil {
		logging.WriteLog("Транзакция не создана")
		logging.TxDenied(ctx, tx.Error)
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка обработки запроса!"})
		logging.WriteLog("Ошибка привязки данных к структуре")
		return
	}

	if err := database.DB.First(&existsPassport, "id_passport = ?", request.Passport.ID_Passport).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Паспорт не найден"})
		logging.WriteLog("Паспорт не найден")
		return
	}

	querry := tx.Model(&models.Passport{}).Where("id_passport = ?", request.Passport.ID_Passport).Updates(map[string]interface{}{
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

	logging.WriteLog("Паспорт - ", request.Passport.ID_Passport, "- изменён")

	//Placework
	var existsPlaceWork models.PlaceWork

	if err := database.DB.First(&existsPlaceWork, "id_placework = ?", request.PlaceWork.ID_Placework).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Место работы не найдено"})
		logging.WriteLog("Место работы не найдено")
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

	logging.WriteLog("Место работы - ", request.PlaceWork.ID_Placework, "- изменёно")

	//RegAddres
	var existsRegAddress models.RegistrationAddress

	if err := database.DB.First(&existsRegAddress, "id_regaddress = ?", request.RegistrationAddress.ID_RegAddress).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Адрес регистрации не найден"})
		logging.WriteLog("Адрес регистрации не найден")
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

	logging.WriteLog("Адрес работы - ", request.RegistrationAddress.ID_RegAddress, "- изменён")

	//ListenerEducation
	var existsEducationListener models.EducationListener

	if err := database.DB.First(&existsEducationListener, "id_educationlistener = ?", request.EducationListener.ID_EducationListener).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Образование не найдено"})
		logging.WriteLog("Образование работы не найдено")
		return
	}

	var levelEducation models.LevelEducation

	if err := database.DB.First(&levelEducation, "education = ?", request.EducationListener.LevelEducation).Error; err != nil {
		logging.WriteLog("Уровень образования не найден", request.EducationListener.LevelEducation)
		logging.TxDenied(ctx, err)
		return
	}

	querry = tx.Model(&models.EducationListener{}).Where("id_educationlistener = ?", request.EducationListener.ID_EducationListener).Updates(map[string]interface{}{
		"diplomseria":            request.EducationListener.DiplomSeria,
		"diplomnumber":           request.EducationListener.DiplomNumber,
		"dategiven":              request.EducationListener.DateGiven,
		"city":                   request.EducationListener.City,
		"region":                 request.EducationListener.Region,
		"educationalinstitution": request.EducationListener.EducationalInstitution,
		"speciality":             request.EducationListener.Speciality,
		"id_leveleducation":      levelEducation.ID_LevelEducation,
	})
	if querry.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: querry.Error, Message: "Ошибка обновления записи"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		logging.TxDenied(ctx, err)
	}

	ctx.JSON(http.StatusOK, gin.H{"message:": "записи успешно изменены"})
	logging.WriteLog("Образование - ", request.EducationListener.ID_EducationListener, "- изменёно")
}
