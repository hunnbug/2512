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

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка привязки данных к структуре"})
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
		err = logging.WriteLog("Паспорт не создан", passport.ID_Passport)
		logging.CheckLogError(err)
		txDenied(ctx)
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
		err = logging.WriteLog("Адрес не создан", registrationAddress.ID_regAddress)
		logging.CheckLogError(err)
		txDenied(ctx)
		return
	}
	logging.WriteLog("Создан адрес слушателя", registrationAddress.ID_regAddress)

	var levelEducation models.LevelEducation
	if err := database.DB.First(&levelEducation, "Education = ?", request.EducationListener.LevelEducation).Error; err != nil {
		err = logging.WriteLog("Уровень образования не найден", request.EducationListener.LevelEducation)
		logging.CheckLogError(err)
		txDenied(ctx)
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
		err = logging.WriteLog("Образование слушателя не создано", educationListener.ID_EducationListener)
		logging.CheckLogError(err)
		txDenied(ctx)
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
		err = logging.WriteLog("Место работы не создано", placeWork.ID_PlaceWork)
		logging.CheckLogError(err)
		txDenied(ctx)
		return
	}
	logging.WriteLog("Создано место работы слушателя", placeWork.ID_PlaceWork)

	var programEducation models.ProgramEducation
	if err := database.DB.First(&programEducation, "nameprofeducation = ?", request.ProgramEducation.NameProfEducation).Error; err != nil {
		err = logging.WriteLog("Уровень образования не найден", request.ProgramEducation.NameProfEducation)
		logging.CheckLogError(err)
		txDenied(ctx)
		return
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
		ID_EducationListener: educationListener.ID_EducationListener,
		ID_PlaceWork:         placeWork.ID_PlaceWork,
		ID_ProgramEducation:  programEducation.ID_ProgramEducation,
	}

	if err := tx.Create(&listener).Error; err != nil {
		tx.Rollback()
		err = logging.WriteLog("Слушатель не создан", listener.ID_Listener)
		logging.CheckLogError(err)
		txDenied(ctx, listener.ID_Listener)
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
	if err := database.DB.First(&request, "id_listener = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Err: err, Message: "Пользователь не найден"})
		logging.WriteLog("Пользователь не найден")
		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка привязки данных к структуре"})
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
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка обновления записи"})
		txDenied(ctx, "Ошибка обновления записи")
		return
	}

	if err := tx.Commit().Error; err != nil {
		txDenied(ctx, query)
		logging.CheckLogError(err)
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
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Err: err, Message: "Пользователь не найден"})
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
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Err: err, Message: "Пользователь не удалён"})
		logging.WriteLog("Пользователь не удалён")
		return
	}

	if err := tx.Where("id_passport = ?", listener.ID_passport).Delete(&models.Passport{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Err: err, Message: "Паспорт не найден"})
		logging.WriteLog("Паспорт не найден")
		return
	}

	if err := tx.Where("id_regaddress = ?", listener.ID_regAddress).Delete(&models.RegistrationAddress{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Err: err, Message: "Адрес не найден"})
		logging.WriteLog("Адрес не найден")
		return
	}

	if err := tx.Where("id_educationlistener = ?", listener.ID_EducationListener).Delete(&models.EducationListener{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Err: err, Message: "Образование не найдено"})
		logging.WriteLog("Образование не найдено")
		return
	}

	if err := tx.Where("id_placework = ?", listener.ID_PlaceWork).Delete(&models.PlaceWork{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Err: err, Message: "Место работы не найдено"})
		logging.WriteLog("Место работы не найдено")
		return
	}

	if err := tx.Where("id_programeducation = ?", listener.ID_ProgramEducation).Delete(&models.ProgramEducation{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Err: err, Message: "Программа обучения не найдена"})
		logging.WriteLog("Программа обучения не найдена")
		return
	}

	if err := tx.Commit().Error; err != nil {
		txDenied(ctx, "Удаление не произведено")
		logging.CheckLogError(err)
	}

	ctx.JSON(http.StatusOK, nil)
}

// примерная версия (значения пагинации брать с url)
func ReadListener(ctx *gin.Context) {
	var listener []models.Listener

	query := database.DB.
		Preload("Passport").
		Preload("RegistrationAddress").
		Preload("EducationListener").
		Preload("PlaceWork").
		Preload("ProgramEducation").
		Limit(3).Offset(0).Find(&listener)
	if query.Error != nil {
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Err: query.Error, Message: "Слушатели не найдены"})
		return
	}

	ctx.JSON(http.StatusOK, listener)
}

func txDenied(ctx *gin.Context, v ...any) {
	logging.WriteLog("Транзакция отменена", v)
	logging.WriteLog("----------------------------------------------")
	ctx.JSON(http.StatusBadRequest, nil)
}
