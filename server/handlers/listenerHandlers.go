package handlers

import (
	"main/database"
	"main/logging"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// создание слушателя
func CreateListener(ctx *gin.Context) {

	var request models.CreateListenerRequest

	logging.WriteLog("запрос: ", request)

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка сервера!"})
		return
	}

	//начало блока логов

	tx := database.DB.Begin()
	if tx.Error != nil {
		err := logging.WriteLog("Транзакция не создана")
		logging.CheckLogError(err)
	}

	passport := models.Passport{
		ID_Passport:   uuid.New(),
		PlaceBirth:    request.Passport.PlaceBirth,
		Citizenship:   request.Passport.Citizenship,
		Gender:        request.Passport.Gender,
		Seria:         request.Passport.Seria,
		Number:        request.Passport.Number,
		PassportGiven: request.Passport.PassportGiven,
		DateGiven:     request.Passport.DateGiven,
		Code:          request.Passport.Code,
	}

	if err := tx.Create(&passport).Error; err != nil {
		tx.Rollback()
		e := logging.WriteLog("Паспорт не создан", passport.ID_Passport)
		logging.CheckLogError(e)
		//txDenied(ctx, models.ErrorResponse{Err: err, Message: "Ошибка при создании паспорта!"})

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка при создании паспорта!"})
		return
	}
	logging.WriteLog("Создан паспорт", passport.ID_Passport)

	registrationAddress := models.RegistrationAddress{
		ID_regAddress: uuid.New(),
		MailIndex:     request.RegistrationAddress.MailIndex,
		Region:        request.RegistrationAddress.Region,
		City:          request.RegistrationAddress.City,
		Street:        request.RegistrationAddress.Street,
		House:         request.RegistrationAddress.House,
		Building:      request.RegistrationAddress.Building,
		Apartment:     request.RegistrationAddress.Apartment,
	}

	if err := tx.Create(&registrationAddress).Error; err != nil {
		tx.Rollback()
		e := logging.WriteLog("Адрес не создан", registrationAddress.ID_regAddress)
		logging.CheckLogError(e)
		//txDenied(ctx, models.ErrorResponse{Err: err, Message: "Ошибка при создании адреса!"})

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка при добавлении адреса пользователя!"})
		return
	}
	logging.WriteLog("Создан адрес слушателя", registrationAddress.ID_regAddress)

	var levelEducation models.LevelEducation
	if err := database.DB.First(&levelEducation, "Education = ?", request.EducationListener.LevelEducation).Error; err != nil {
		e := logging.WriteLog("Уровень образования не найден", request.EducationListener.LevelEducation)
		logging.CheckLogError(e)
		//txDenied(ctx, models.ErrorResponse{Err: err, Message: "Ошибка при добавлении уровня образования!"})

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка при добавлении уровня образования слушателя!"})
		return
	}

	educationListener := models.EducationListener{
		ID_EducationListener:   uuid.New(),
		DiplomSeria:            request.EducationListener.DiplomSeria,
		DiplomNumber:           request.EducationListener.DiplomNumber,
		DateGiven:              request.EducationListener.DateGiven,
		City:                   request.EducationListener.City,
		Region:                 request.EducationListener.Region,
		EducationalInstitution: request.EducationListener.EducationalInstitution,
		Speciality:             request.EducationListener.Speciality,
		ID_LevelEducation:      levelEducation.ID_LevelEducation,
	}

	if err := tx.Create(&educationListener).Error; err != nil {
		tx.Rollback()
		e := logging.WriteLog("Образование слушателя не создано", educationListener.ID_EducationListener)
		logging.CheckLogError(e)
		//txDenied(ctx, models.ErrorResponse{Err: err, Message: "Ошибка при добавлении информации об образовании слушателя!"})

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка при добавлении информации об образовании слушателя!"})
		return
	}
	logging.WriteLog("Создано образования слушателя", educationListener.ID_EducationListener)

	placeWork := models.PlaceWork{
		ID_PlaceWork:       uuid.New(),
		NameCompany:        request.PlaceWork.NameCompany,
		JobTitle:           request.PlaceWork.JobTitle,
		AllExperience:      request.PlaceWork.AllExperience,
		JobTitleExpirience: request.PlaceWork.JobTitleExpirience,
	}

	if err := tx.Create(&placeWork).Error; err != nil {
		tx.Rollback()
		e := logging.WriteLog("Место работы не создано", placeWork.ID_PlaceWork)
		logging.CheckLogError(e)
		//txDenied(ctx, models.ErrorResponse{Err: err, Message: "Ошибка при добавлении информации о месте работы слушателя!"})

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка при добавлении места работы слушателя!"})
		return
	}
	logging.WriteLog("Создано место работы слушателя", placeWork.ID_PlaceWork)

	// var programEducation models.ProgramEducation
	// if err := database.DB.First(&programEducation, "nameprofeducation = ?", request.ProgramEducation.NameProfEducation).Error; err != nil {
	// 	e := logging.WriteLog("Уровень образования не найден", request.ProgramEducation.NameProfEducation)
	// 	logging.CheckLogError(err)
	// 	txDenied(ctx, models.ErrorResponse{Err: err, Message: ""})})
	// 	return
	// }

	listener := models.Listener{
		ID_Listener:          uuid.New(),
		FirstName:            request.FirstName,
		SecondName:           request.SecondName,
		MiddleName:           request.MiddleName,
		DateOfBirth:          request.DateOfBirth,
		SNILS:                request.SNILS,
		ContactPhone:         request.ContactPhone,
		Email:                request.Email,
		ID_passport:          passport.ID_Passport,
		ID_regAddress:        registrationAddress.ID_regAddress,
		ID_EducationListener: educationListener.ID_EducationListener,
		ID_PlaceWork:         placeWork.ID_PlaceWork,
	}

	if err := tx.Create(&listener).Error; err != nil {
		tx.Rollback()
		e := logging.WriteLog("Слушатель не создан", listener.ID_Listener)
		logging.CheckLogError(e)
		//txDenied(ctx, listener.ID_Listener)

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка при создании слушателя!"})
		return
	}
	logging.WriteLog("Создан слушатель", listener.ID_Listener)

	if err := tx.Commit().Error; err != nil {
		txDenied(ctx, listener.ID_Listener)
		logging.CheckLogError(err)
		return
	}
	logging.WriteLog("Записи успешно зарегистрированы, слушатель - ", listener.ID_Listener)

	//конец записи логов
	logging.WriteLog("----------------------------------------------")

	ctx.JSON(http.StatusCreated, nil)
}

func UpdateListener(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка Parse id"})
		logging.WriteLog("Ошибка Parse id")
		return
	}

	var request models.Listener
	var existsListener models.Listener

	if err := database.DB.First(&existsListener, "id_listener = ?", id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Пользователь не найден"})
		logging.WriteLog("Пользователь не найден")
		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка обработки запроса!"})
		logging.WriteLog("Ошибка привязки данных к структуре")
		return
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		err := logging.WriteLog("Транзакция не создана")
		logging.CheckLogError(err)
	}

	query := tx.Model(&models.Listener{}).Where("id_listener = ?", id).Updates(map[string]interface{}{
		"firstname":    request.FirstName,
		"secondname":   request.SecondName,
		"middlename":   request.MiddleName,
		"dateofbirth":  request.DateOfBirth,
		"snils":        request.SNILS,
		"contactphone": request.ContactPhone,
		"email":        request.Email,
	})
	if query.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: query.Error, Message: "Ошибка обновления записи"})
		txDenied(ctx, "Ошибка обновления записи")
		return
	}

	if err := tx.Commit().Error; err != nil {
		txDenied(ctx, query)
	}

	ctx.JSON(http.StatusOK, nil)

}

func DeleteListener(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка Parse id"})
		logging.WriteLog("Ошибка Parse id")
		return
	}

	var listener models.Listener
	if err := database.DB.First(&listener, "id_listener = ?", id).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Пользователь не найден"})
		logging.WriteLog("Пользователь не найден")
		return
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		err := logging.WriteLog("Транзакция не создана")
		logging.CheckLogError(err)
	}

	if err := tx.Delete(&models.Listener{}, id).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Пользователь не удалён"})
		logging.WriteLog("Пользователь не удалён")
		return
	}

	if err := tx.Where("id_passport = ?", listener.ID_passport).Delete(&models.Passport{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Паспорт не найден"})
		logging.WriteLog("Паспорт не найден")
		return
	}

	if err := tx.Where("id_regaddress = ?", listener.ID_regAddress).Delete(&models.RegistrationAddress{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Адрес не найден"})
		logging.WriteLog("Адрес не найден")
		return
	}

	if err := tx.Where("id_educationlistener = ?", listener.ID_EducationListener).Delete(&models.EducationListener{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Образование не найдено"})
		logging.WriteLog("Образование не найдено")
		return
	}

	if err := tx.Where("id_placework = ?", listener.ID_PlaceWork).Delete(&models.PlaceWork{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Место работы не найдено"})
		logging.WriteLog("Место работы не найдено")
		return
	}

	if err := tx.Commit().Error; err != nil {
		txDenied(ctx, "Удаление не произведено")
		logging.CheckLogError(err)
	}
	logging.WriteLog("-----------------")
	logging.WriteLog("Удалён пользователь и все смежные данные", listener.ID_Listener)
	logging.WriteLog("-----------------")

	ctx.JSON(http.StatusOK, nil)
}

// примерная версия (значения пагинации брать с url) желательно переделать под DTO
func ReadListener(ctx *gin.Context) {

	const LIMIT_COUNT = 10

	logging.WriteLog("получен запрос на получение слушателей")

	type request struct {
		CurrentPage  int
		FirstName    string
		SecondName   string
		MiddleName   string
		ContactPhone string
		Email        string
	}

	var _request request

	err := ctx.BindJSON(&_request)

	if err != nil {
		e := logging.WriteLog("не удалось получить ответ от страницы: ", err)

		logging.CheckLogError(e)
	}

	logging.WriteLog("был получен запрос: ", _request)

	var listeners []models.Listener

	query := database.DB.Limit(LIMIT_COUNT).Offset((_request.CurrentPage-1)*LIMIT_COUNT).Find(&listeners,
		`firstName LIKE ?
		OR secondName LIKE ?
		OR middleName LIKE ?
		OR email LIKE ?
		OR contactphone LIKE ?`,
		"%"+_request.FirstName+"%",
		"%"+_request.SecondName+"%",
		"%"+_request.MiddleName+"%",
		"%"+_request.Email+"%",
		"%"+_request.ContactPhone+"%")

	if query.Error != nil {

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: query.Error, Message: "Слушатели не найдены"})

		return
	}

	ctx.JSON(http.StatusOK, listeners)
}

func txDenied(ctx *gin.Context, v ...any) {
	logging.WriteLog("Транзакция отменена", v)
	logging.WriteLog("----------------------------------------------")
	ctx.JSON(http.StatusBadRequest, v)
}
