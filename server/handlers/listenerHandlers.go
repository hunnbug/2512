package handlers

import (
	"main/database"
	"main/logging"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateListener(ctx *gin.Context) {

	var request models.CreateListenerRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка привязки данных к структуре"})
		return
	}

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
		err = logging.WriteLog("Паспорт не создан", passport)
		logging.CheckLogError(err)
	}

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
		err = logging.WriteLog("Адрес не создан", registrationAddress)
		logging.CheckLogError(err)
	}

	var levelEducation models.LevelEducation
	if err := database.DB.First(&levelEducation, "Education = ?", request.EducationListener.LevelEducation).Error; err != nil {
		err = logging.WriteLog("Уровень образования не найден", request.EducationListener.LevelEducation)
		logging.CheckLogError(err)
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
		err = logging.WriteLog("Образование слушателя не создано", educationListener)
		logging.CheckLogError(err)
	}

	placeWork := models.PlaceWork{
		ID_PlaceWork:       uuid.New(),
		NameCompany:        request.PlaceWork.NameCompany,
		JobTitle:           request.PlaceWork.JobTitle,
		AllExperience:      request.PlaceWork.AllExperience,
		JobTitleExpirience: request.PlaceWork.JobTitleExpirience,
	}

	if err := tx.Create(&placeWork).Error; err != nil {
		tx.Rollback()
		err = logging.WriteLog("Место работы не создано", placeWork)
		logging.CheckLogError(err)
	}

	programEducation := models.ProgramEducation{
		ID_ProgramEducation: uuid.New(),
		NameProfEducation:   request.ProgramEducation.NameProfEducation,
		TypeOfEducation:     request.ProgramEducation.TypeOfEducation,
		TimeEducation:       request.ProgramEducation.TimeEducation,
	}

	if err := tx.Create(&programEducation).Error; err != nil {
		tx.Rollback()
		err = logging.WriteLog("Программа обучения не найдена", programEducation)
		logging.CheckLogError(err)
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
		err = logging.WriteLog("Пользователь не создан", listener)
		logging.CheckLogError(err)
	}

	if err := tx.Commit().Error; err != nil {
		err = logging.WriteLog("Транзакция отменена", listener)
		logging.CheckLogError(err)
	}

	ctx.JSON(http.StatusCreated, nil)
}

func UpdateListener(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка Parse id"})
		return
	}

	var request models.Listener
	if err := database.DB.First(&request, "id_listener = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Err: err, Message: "Пользователь не найден"})
		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка привязки данных к структуре"})
		return
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		err := logging.WriteLog("Транзакция не создана")
		logging.CheckLogError(err)
	}

	result := tx.Model(&models.Listener{}).Where("id_listener = ?", id).Updates(map[string]interface{}{
		"firstname":    request.FirstName,
		"secondname":   request.SecondName,
		"middlename":   request.MiddleName,
		"dateofbirth":  request.DateOfBirth,
		"snils":        request.SNILS,
		"contactphone": request.ContactPhone,
		"email":        request.Email,
	})
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка обновления записи"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		err = logging.WriteLog("Транзакция отменена", result)
		logging.CheckLogError(err)
	}

	ctx.JSON(http.StatusOK, nil)

}

func DeleteListener(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Ошибка Parse id"})
		return
	}

	var listener models.Listener
	if err := database.DB.First(&listener, "id_listener = ?", id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Err: err, Message: "Пользователь не найден"})
		return
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		err := logging.WriteLog("Транзакция не создана")
		logging.CheckLogError(err)
	}

	if err := tx.Delete(&models.Listener{}, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Err: err, Message: "Пользователь не удалён"})
		return
	}

	if err := tx.Where("id_passport = ?", listener.ID_passport).Delete(&models.Passport{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Err: err, Message: "Пользователь не найден"})
		return
	}

	if err := tx.Where("id_regaddress = ?", listener.ID_regAddress).Delete(&models.RegistrationAddress{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Err: err, Message: "Адрес не найден"})
		return
	}

	if err := tx.Where("id_educationlistener = ?", listener.ID_EducationListener).Delete(&models.EducationListener{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Err: err, Message: "Образование не найдено"})
		return
	}

	if err := tx.Where("id_placework = ?", listener.ID_PlaceWork).Delete(&models.PlaceWork{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Err: err, Message: "Место работы не найдено"})
		return
	}

	if err := tx.Where("id_programeducation = ?", listener.ID_ProgramEducation).Delete(&models.ProgramEducation{}).Error; err != nil {
		tx.Rollback()
		ctx.JSON(http.StatusNotFound, models.ErrorResponse{Err: err, Message: "Программа обучения не найдена"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		err = logging.WriteLog("Транзакция отменена")
		logging.CheckLogError(err)
	}

	ctx.JSON(http.StatusOK, nil)
}
