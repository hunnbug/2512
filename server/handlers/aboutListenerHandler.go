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

	responseUUID := models.ListenerIDDTO{
		ID_Passport:          listener.ID_passport,
		ID_RegAddress:        listener.ID_regAddress,
		ID_EducationListener: listener.ID_EducationListener,
		ID_PlaceWork:         listener.ID_PlaceWork,
		ID_ProgramEducation:  listener.ID_ProgramEducation,
	}

	// Все другие таблицы в которых поиск идёт по id из предыдущей структуры
	var passport models.Passport
	var regAddress models.RegistrationAddress
	var educationListener models.EducationListener
	var placework models.PlaceWork
	var programEducation models.ProgramEducation

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

	if err := database.DB.First(&programEducation, "id_programeducation = ?", responseUUID.ID_ProgramEducation).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Программы обучения не найдены"})
		logging.WriteLog("Программы обучения не найдены")
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

	responseProgramEducation := models.ProgramEducationDTO{
		NameProfEducation: programEducation.NameProfEducation,
		TypeOfEducation:   programEducation.TypeOfEducation,
		TimeEducation:     programEducation.TimeEducation,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"listener":          responseListener,
		"passport":          responsePassport,
		"regAddress":        responseRegAddress,
		"educationListener": responseEducationListener,
		"placeWork":         responsePlacework,
		"programEducation":  responseProgramEducation,
	})

}
