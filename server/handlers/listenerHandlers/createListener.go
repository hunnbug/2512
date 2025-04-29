package listenerHandlers

import (
	"fmt"
	"main/database"
	"main/logging"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// создание слушателя
func CreateListener(ctx *gin.Context) {

	logging.WriteLog(logging.DEBUG, "----------------------------------------------")
	logging.WriteLog(logging.DEBUG, "запрос на создание слушателя")

	var request models.CreateListenerRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка сервера!"})

		logging.WriteLog(logging.ERROR, "Ошибка привязки данных к структуре")
		fmt.Println(err)
		return
	}

	fmt.Println(request)

	tx := database.DB.Begin()
	if tx.Error != nil {
		logging.WriteLog(logging.ERROR, "Транзакция не создана")
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
		logging.WriteLog(logging.ERROR, "Паспорт не создан", passport.ID_Passport)

		logging.TxDenied(ctx, err)

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка при создании паспорта!"})
		return
	}
	logging.WriteLog(logging.DEBUG, "Создан паспорт", passport.ID_Passport)

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
		logging.WriteLog(logging.ERROR, "Адрес не создан", registrationAddress.ID_regAddress)

		logging.TxDenied(ctx, err)

		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка при добавлении адреса пользователя!"})
		return
	}
	logging.WriteLog(logging.DEBUG, "Создан адрес слушателя", registrationAddress.ID_regAddress)

	var ID_EducationListener *uuid.UUID = nil
	var educationListener models.EducationListener

	if request.EducationListener != (models.EducationListenerRequest{}) {
		newID := uuid.New()
		ID_EducationListener = &newID
		educationListener = models.EducationListener{
			ID_EducationListener:   *ID_EducationListener,
			DiplomSeria:            request.EducationListener.DiplomSeria,
			DiplomNumber:           request.EducationListener.DiplomNumber,
			DateGiven:              request.EducationListener.DateGiven,
			City:                   request.EducationListener.City,
			Region:                 request.EducationListener.Region,
			EducationalInstitution: request.EducationListener.EducationalInstitution,
			Speciality:             request.EducationListener.Speciality,
			LevelEducation:         request.EducationListener.LevelEducation,
		}

		if err := tx.Create(&educationListener).Error; err != nil {
			tx.Rollback()
			logging.WriteLog(logging.ERROR, "Образование слушателя не создано", educationListener.ID_EducationListener)
			logging.TxDenied(ctx, err)

			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка при добавлении информации об образовании слушателя!"})
			return
		}
		logging.WriteLog(logging.DEBUG, "Создано образования слушателя", educationListener.ID_EducationListener)
	}

	var ID_PlaceWork *uuid.UUID = nil

	var placeWork models.PlaceWork

	if request.PlaceWork != (models.PlaceWorkRequest{}) {
		newID := uuid.New()
		ID_PlaceWork = &newID
		placeWork = models.PlaceWork{
			ID_PlaceWork:       *ID_PlaceWork,
			NameCompany:        request.PlaceWork.NameCompany,
			JobTitle:           request.PlaceWork.JobTitle,
			AllExperience:      request.PlaceWork.AllExperience,
			JobTitleExpirience: request.PlaceWork.JobTitleExpirience,
		}

		if err := tx.Create(&placeWork).Error; err != nil {
			tx.Rollback()
			logging.WriteLog(logging.ERROR, "Место работы не создано", placeWork.ID_PlaceWork)
			logging.TxDenied(ctx, err)

			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка при добавлении места работы слушателя!"})
			return
		}
		logging.WriteLog(logging.DEBUG, "Создано место работы слушателя", placeWork.ID_PlaceWork)
	}

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
		ID_EducationListener: ID_EducationListener,
		ID_PlaceWork:         ID_PlaceWork,
	}

	if err := tx.Create(&listener).Error; err != nil {
		tx.Rollback()
		logging.WriteLog(logging.ERROR, "Слушатель не создан", listener.ID_Listener)

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
