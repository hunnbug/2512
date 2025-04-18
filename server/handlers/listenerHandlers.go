package handlers

import (
	"main/database"
	"main/logging"
	"main/models"
	"main/tools"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// создание слушателя
func CreateListener(ctx *gin.Context) {

	logging.WriteLog(logging.ERROR, logging.DEBUG, "----------------------------------------------")
	logging.WriteLog(logging.ERROR, logging.DEBUG, "запрос на создание слушателя")

	var request models.CreateListenerRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка сервера!"})
		return
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		logging.WriteLog(logging.ERROR, logging.ERROR, "Транзакция не создана")
		return
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
		logging.WriteLog(logging.ERROR, logging.ERROR, "Паспорт не создан", passport.ID_Passport)

		logging.TxDenied(ctx, err)

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка при создании паспорта!"})
		return
	}
	logging.WriteLog(logging.ERROR, logging.DEBUG, "Создан паспорт", passport.ID_Passport)

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
		logging.WriteLog(logging.ERROR, logging.ERROR, "Адрес не создан", registrationAddress.ID_regAddress)

		logging.TxDenied(ctx, err)

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка при добавлении адреса пользователя!"})
		return
	}
	logging.WriteLog(logging.ERROR, logging.DEBUG, "Создан адрес слушателя", registrationAddress.ID_regAddress)

	var levelEducation models.LevelEducation
	if err := database.DB.First(&levelEducation, "Education = ?", request.EducationListener.LevelEducation).Error; err != nil {
		logging.WriteLog(logging.ERROR, logging.ERROR, "Уровень образования не найден", request.EducationListener.LevelEducation)
		logging.TxDenied(ctx, err)

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
		logging.WriteLog(logging.ERROR, logging.ERROR, "Образование слушателя не создано", educationListener.ID_EducationListener)
		logging.TxDenied(ctx, err)

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка при добавлении информации об образовании слушателя!"})
		return
	}
	logging.WriteLog(logging.ERROR, logging.DEBUG, "Создано образования слушателя", educationListener.ID_EducationListener)

	placeWork := models.PlaceWork{
		ID_PlaceWork:       uuid.New(),
		NameCompany:        request.PlaceWork.NameCompany,
		JobTitle:           request.PlaceWork.JobTitle,
		AllExperience:      request.PlaceWork.AllExperience,
		JobTitleExpirience: request.PlaceWork.JobTitleExpirience,
	}

	if err := tx.Create(&placeWork).Error; err != nil {
		tx.Rollback()
		logging.WriteLog(logging.ERROR, logging.ERROR, "Место работы не создано", placeWork.ID_PlaceWork)
		logging.TxDenied(ctx, err)

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка при добавлении места работы слушателя!"})
		return
	}
	logging.WriteLog(logging.ERROR, logging.DEBUG, "Создано место работы слушателя", placeWork.ID_PlaceWork)

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
		logging.WriteLog(logging.ERROR, logging.ERROR, "Слушатель не создан", listener.ID_Listener)

		logging.TxDenied(ctx, listener.ID_Listener)

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка при создании слушателя!"})
		return
	}
	logging.WriteLog(logging.DEBUG, "Создан слушатель", listener.ID_Listener)

	if err := tx.Commit().Error; err != nil {
		logging.TxDenied(ctx, listener.ID_Listener)
		return
	}
	logging.WriteLog(logging.DEBUG, "Записи успешно зарегистрированы, слушатель - ", listener.ID_Listener)

	logging.WriteLog(logging.DEBUG, "----------------------------------------------")

	ctx.JSON(http.StatusCreated, nil)
}

func DeleteListener(ctx *gin.Context) {
	id, err := tools.CheckParamID(ctx)
	if err != nil {
		return
	}

	var listener models.Listener
	err = tools.CheckRecord(ctx, &listener, "id_listener = ?", id)
	if err != nil {
		return
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		logging.WriteLog(logging.ERROR, "Транзакция не создана")
		return
	}

	tools.DeleteRows(ctx, tx, &models.Listener{}, "id_listener = ?", listener.ID_Listener)

	tools.DeleteRows(ctx, tx, &models.Passport{}, "id_passport = ?", listener.ID_passport)

	tools.DeleteRows(ctx, tx, &models.RegistrationAddress{}, "id_regaddress = ?", listener.ID_regAddress)

	tools.DeleteRows(ctx, tx, &models.EducationListener{}, "id_educationlistener = ?", listener.ID_EducationListener)

	tools.DeleteRows(ctx, tx, &models.PlaceWork{}, "id_placework = ?", listener.ID_PlaceWork)

	if err := tx.Commit().Error; err != nil {
		logging.TxDenied("Удаление не произведено")
	}

	ctx.JSON(http.StatusOK, nil)

	if tx.Error == nil {
		logging.WriteLog(logging.DEBUG, "-----------------")
		logging.WriteLog(logging.DEBUG, "Удалён пользователь и все смежные данные", listener.ID_Listener)
		logging.WriteLog(logging.DEBUG, "-----------------")
	}
}

func ReadListener(ctx *gin.Context) {

	const LIMIT_COUNT = 10

	logging.WriteLog(logging.DEBUG, "получен запрос на получение слушателей")

	type request struct {
		CurrentPage int
		FirstField  string
		SecondField string
		ThirdField  string
		FourthField string
		FifthField  string
		EmptyForm   bool
	}

	var _request request

	err := ctx.BindJSON(&_request)

	if err != nil {
		logging.WriteLog(logging.ERROR, "не удалось получить ответ от страницы: ", err)

	}

	var listeners []models.Listener

	//
	//Если фильтры не введены - работает обычная пагинация
	//
	if _request.EmptyForm {

		query := database.DB.Limit(LIMIT_COUNT).Offset((_request.CurrentPage - 1) * LIMIT_COUNT).Find(&listeners)

		if query.Error != nil {

			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: query.Error, Message: "Слушатели не найдены"})

			return

		}

	} else {

		type fieldsDTO struct {
			firstField  string
			secondField string
			thirdField  string
			fourthField string
			fifthField  string
		}

		fields := fieldsDTO{

			firstField:  _request.FirstField,
			secondField: _request.SecondField,
			thirdField:  _request.ThirdField,
			fourthField: _request.FourthField,
			fifthField:  _request.FifthField,
		}

		value := reflect.ValueOf(fields)

		var notNullFields []string

		//
		//перебор всех полей из фильтрации и отсечение пустых
		//
		for i := range value.NumField() {

			if value.Field(i).String() != "" {

				notNullFields = append(notNullFields, value.Field(i).String())

			}

		}

		var requestString string

		//
		//добавление всех строк для фильтрации в запрос
		//
		for i := 0; i < len(notNullFields)-1; i++ {

			requestString += "(firstname LIKE '%" + notNullFields[i] + "%' OR "
			requestString += "secondname LIKE '%" + notNullFields[i] + "%' OR "
			requestString += "middlename LIKE '%" + notNullFields[i] + "%' OR "
			requestString += "contactphone LIKE '%" + notNullFields[i] + "%' OR "
			requestString += "email LIKE '%" + notNullFields[i] + "%') AND "

		}

		requestString += "(firstname LIKE '%" + notNullFields[len(notNullFields)-1] + "%' OR "
		requestString += "secondname LIKE '%" + notNullFields[len(notNullFields)-1] + "%' OR "
		requestString += "middlename LIKE '%" + notNullFields[len(notNullFields)-1] + "%' OR "
		requestString += "contactphone LIKE '%" + notNullFields[len(notNullFields)-1] + "%' OR "
		requestString += "email LIKE '%" + notNullFields[len(notNullFields)-1] + "%')"

		database.DB.Find(&listeners, requestString)

	}

	ctx.JSON(http.StatusOK, listeners)
}
