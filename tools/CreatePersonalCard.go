package tools

import (
	"bytes"
	"fmt"
	"main/environment"
	"main/logging"
	"main/models"
	"main/storage"
	"strconv"
	"time"

	"github.com/nguyenthenguyen/docx"
)

func CreatePersonalCard(Listenerdata *models.FullListenerDataDTO, EducationData *models.EducationData) error {
	templatePath := "./documents/PersonalCard.docx"

	r, err := docx.ReadDocxFile(templatePath)
	if err != nil {
		panic(fmt.Sprintf("Не удалось открыть шаблон: %v", err))
	}
	defer r.Close()

	doc := r.Editable()

	replacePlaceholders(doc, Listenerdata, EducationData)

	var buffer bytes.Buffer
	err = doc.Write(&buffer)
	if err != nil {
		logging.WriteLog(logging.ERROR, err)
		return err
	}

	fio := Listenerdata.Listener.SecondName + Listenerdata.Listener.FirstName + Listenerdata.Listener.MiddleName
	nameFile := "Личное дело - " + fio + ".docx"

	s3client := storage.CreateS3Client()
	err = S3Load(s3client, environment.S3.Bucket, nameFile, buffer.Bytes())
	if err != nil {
		logging.WriteLog(logging.ERROR, err)
		return err
	}

	return nil

}

func replacePlaceholders(doc *docx.Docx, Listenerdata *models.FullListenerDataDTO, EducationData *models.EducationData) {
	doc.Replace("}}", "", -1)
	doc.Replace("{{", "", -1)
	doc.Replace("ProgramEducation", EducationData.ProgramEducation.NameProfEducation, -1)
	doc.Replace("TypeOfEducation", EducationData.TypeEducation.TypeName, -1)
	doc.Replace("Hour", strconv.Itoa(EducationData.ProgramEducation.TimeEducation), -1)

	fio := Listenerdata.Listener.SecondName + " " + Listenerdata.Listener.FirstName + " " + Listenerdata.Listener.MiddleName

	doc.Replace("FIO", fio, -1)

	dob, err := time.Parse(time.RFC3339, Listenerdata.Listener.DateOfBirth)
	if err == nil {
		doc.Replace("DateBirth", fmt.Sprintf("%02d", dob.Day()), -1)
		doc.Replace("MonthBirth", fmt.Sprintf("%02d", dob.Month()), -1)
		doc.Replace("YearBirth", fmt.Sprintf("%d", dob.Year()), -1)
	}

	doc.Replace("City", Listenerdata.Passport.PlaceBirth, -1)
	doc.Replace("Citizenship", Listenerdata.Passport.Citizenship, -1)

	if Listenerdata.Passport.Gender == "Мужской" {
		doc.Replace("мужской ☐", "мужской ☒", -1)
	} else if Listenerdata.Passport.Gender == "Женский" {
		doc.Replace("женский ☐", "женский ☒", -1)
	}

	doc.Replace("Seria", strconv.Itoa(Listenerdata.Passport.Seria), -1)
	doc.Replace("Number", strconv.Itoa(Listenerdata.Passport.Number), -1)
	doc.Replace("WhoGiven", Listenerdata.Passport.PassportGiven, -1)

	given, err := time.Parse(time.RFC3339, Listenerdata.Passport.DateGiven)
	if err == nil {
		doc.Replace("WhenGiven", given.Format("02.01.2006"), -1)
	}

	doc.Replace("MainIndex", strconv.Itoa(Listenerdata.RegistrationAddress.MailIndex), -1)
	doc.Replace("Region", Listenerdata.RegistrationAddress.Region, -1)
	doc.Replace("City", Listenerdata.RegistrationAddress.City, -1)
	doc.Replace("Street", Listenerdata.RegistrationAddress.Street, -1)
	doc.Replace("House", Listenerdata.RegistrationAddress.House, -1)
	doc.Replace("Building", Listenerdata.RegistrationAddress.Building, -1)
	doc.Replace("Appartaments", Listenerdata.RegistrationAddress.Apartment, -1)
	doc.Replace("SNILS", Listenerdata.Listener.SNILS, -1)
	doc.Replace("Phone", Listenerdata.Listener.ContactPhone, -1)
	doc.Replace("Email", Listenerdata.Listener.Email, -1)

	switch Listenerdata.EducationListener.LevelEducation {
	case "Среднее":
		doc.Replace("среднее ☐", "среднее ☒", -1)
	case "Среднее профессиональное":
		doc.Replace("среднее профессиональное ☐", "среднее профессиональное ☒", -1)
	case "Высшее":
		doc.Replace("высшее ☐", "высшее ☒", -1)
	}

	if Listenerdata.EducationListener != (models.EducationListenerDTO{}) {
		doc.Replace("диплом ☐", "диплом ☒", -1)
		doc.Replace("SDip", strconv.Itoa(Listenerdata.EducationListener.DiplomSeria), -1)
		doc.Replace("NDip", strconv.Itoa(Listenerdata.EducationListener.DiplomNumber), -1)
		dmy, err := time.Parse(time.RFC3339, Listenerdata.Listener.DateOfBirth)
		if err == nil {
			doc.Replace("DDip", fmt.Sprintf("%02d", dmy.Day()), -1)
			doc.Replace("MDip", fmt.Sprintf("%02d", dmy.Month()), -1)
			doc.Replace("YDip", fmt.Sprintf("%d", dmy.Year()), -1)
		}
		doc.Replace("CDip", Listenerdata.EducationListener.City, -1)
		doc.Replace("RDip", Listenerdata.EducationListener.Region, -1)
		doc.Replace("WhoDip", Listenerdata.EducationListener.EducationalInstitution, -1)
		doc.Replace("Speciality", Listenerdata.EducationListener.Speciality, -1)
	} else {
		doc.Replace("SDip", "_______", -1)
		doc.Replace("NDip", "_______", -1)
		doc.Replace("DDip", "_______", -1)
		doc.Replace("MDip", "_______", -1)
		doc.Replace("YDip", "_______", -1)

		doc.Replace("CDip", "_______", -1)
		doc.Replace("RDip", "_______", -1)
		doc.Replace("WhoDip", "_______", -1)
		doc.Replace("Speciality", "_______", -1)
	}

	if Listenerdata.PlaceWork != (models.PlaceWorkDTO{}) {
		doc.Replace("PlaceWork", Listenerdata.PlaceWork.NameCompany, -1)
		doc.Replace("Post", Listenerdata.PlaceWork.JobTitle, -1)
		doc.Replace("AllP", strconv.Itoa(Listenerdata.PlaceWork.AllExperience), -1)
		doc.Replace("OnP", strconv.Itoa(Listenerdata.PlaceWork.JobTitleExpirience), -1)
	} else {
		doc.Replace("PlaceWork", "_______", -1)
		doc.Replace("Post", "_______", -1)
		doc.Replace("AllP", "_______", -1)
		doc.Replace("OnP", "_______", -1)
	}

	doc.Replace("Division", EducationData.Division.Divisions, -1)

}
