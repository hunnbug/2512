package tools

import (
	"main/database"
	"main/logging"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllListenerData(ctx *gin.Context, id uuid.UUID) (*models.FullListenerDataDTO, error) {

	//listener
	var listener models.Listener

	querry := database.DB.First(&listener, "id_listener = ?", id)
	if querry.Error != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: querry.Error, Message: "слушатель не найден"})
		logging.WriteLog(logging.ERROR, "слушатель не найден")
		return nil, querry.Error
	}
	logging.WriteLog(logging.DEBUG, "Пользователь найден", id)

	//responseUUID
	responseUUID := models.ListenerIDDTO{
		ID_Passport:          listener.ID_passport,
		ID_RegAddress:        listener.ID_regAddress,
		ID_EducationListener: listener.ID_EducationListener,
		ID_PlaceWork:         listener.ID_PlaceWork,
	}
	//passport
	var passport models.Passport
	if err := database.DB.First(&passport, "id_passport = ?", responseUUID.ID_Passport).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Паспорт не найден"})
		logging.WriteLog(logging.ERROR, "Паспорт не найден")
		return nil, err
	}

	//regaddress
	var regAddress models.RegistrationAddress
	if err := database.DB.First(&regAddress, "id_regaddress = ?", responseUUID.ID_RegAddress).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Адрес не найден"})
		logging.WriteLog(logging.ERROR, "Адрес не найден")
		return nil, err
	}
	//education
	var educationListener models.EducationListener
	if responseUUID.ID_EducationListener != nil {
		if err := database.DB.First(&educationListener, "id_educationlistener = ?", responseUUID.ID_EducationListener).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Образование не найдено"})
			logging.WriteLog(logging.ERROR, "Образование не найдено")
			return nil, err
		}
	}
	//placework
	var placework models.PlaceWork
	if responseUUID.ID_PlaceWork != nil {
		if err := database.DB.First(&placework, "id_placework = ?", responseUUID.ID_PlaceWork).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Место работы не найдено"})
			logging.WriteLog(logging.ERROR, "Место работы не найдено")
			return nil, err
		}
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

	var leveleducation models.LevelEducation
	if err := database.DB.Find(&leveleducation, "id_leveleducation = ?", educationListener.ID_LevelEducation).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{Err: err, Message: "Уровень образования не найден"})
		logging.WriteLog(logging.ERROR, "Уровень образования не найден")
		return nil, err
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

	return &models.FullListenerDataDTO{
		Listener:            responseListener,
		Passport:            responsePassport,
		RegistrationAddress: responseRegAddress,
		EducationListener:   responseEducationListener,
		PlaceWork:           responsePlacework,
		ResponseUUID:        responseUUID,
	}, nil
}
