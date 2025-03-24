package handlers

import (
	"log"
	"main/database"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateListener(ctx *gin.Context) {

	var request models.CreateListenerRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error:": err})
		return
	}

	tx := database.DB.Begin()
	if tx.Error != nil {
		log.Fatal(tx.Error)
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
		log.Fatal(err)
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
		log.Fatal(err)
	}

	var levelEducation models.LevelEducation
	if err := database.DB.First(&levelEducation, "Education = ?", request.EducationListener.LevelEducation).Error; err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
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
		log.Fatal(err)
	}

	programEducation := models.ProgramEducation{
		ID_ProgramEducation: uuid.New(),
		NameProfEducation:   request.ProgramEducation.NameProfEducation,
		TypeOfEducation:     request.ProgramEducation.TypeOfEducation,
		TimeEducation:       request.ProgramEducation.TimeEducation,
	}

	if err := tx.Create(&programEducation).Error; err != nil {
		tx.Rollback()
		log.Fatal(err)
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
		log.Fatal(err)
	}

	if err := tx.Commit().Error; err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusCreated, nil)
}
