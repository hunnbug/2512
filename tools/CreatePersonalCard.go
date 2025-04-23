package tools

import (
	"fmt"
	"main/models"
	"strconv"
	"time"

	"github.com/nguyenthenguyen/docx"
)

func CreatePersonalCard(Listenerdata *models.FullListenerDataDTO, EducationData *models.EducationData) {
	templatePath := "./documents/PersonalCard.docx"
	newFilePath := "./documents/output.docx"

	r, err := docx.ReadDocxFile(templatePath)
	if err != nil {
		panic(fmt.Sprintf("Не удалось открыть шаблон: %v", err))
	}
	defer r.Close()

	doc := r.Editable()

	replacePlaceholders(doc, Listenerdata, EducationData)

	err = doc.WriteToFile(newFilePath)
	if err != nil {
		panic(fmt.Sprintf("Не удалось сохранить документ: %v", err))
	}

	fmt.Printf("Документ успешно изменен и сохранен как %s\n", newFilePath)
}

func replacePlaceholders(doc *docx.Docx, Listenerdata *models.FullListenerDataDTO, EducationData *models.EducationData) {
	doc.Replace("}}", "", -1)
	doc.Replace("{{", "", -1)
	doc.Replace("ProgramEducation", EducationData.ProgramEducation.NameProfEducation, -1)
	doc.Replace("TypeOfEducation", EducationData.TypeEducation.TypeName, -1)

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

	if Listenerdata.Passport.Gender == "мужской" {
		doc.Replace("мужской ☐", "мужской ☑", -1)
	} else if Listenerdata.Passport.Gender == "женский" {
		doc.Replace("женский ☐", "женский ☑", -1)
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
	case "среднее":
		doc.Replace("среднее ☐", "среднее ☑", -1)
	case "среднее профессиональное":
		doc.Replace("среднее профессиональное ☐", "среднее профессиональное ☑", -1)
	case "высшее":
		doc.Replace("высшее ☐", "высшее ☑", -1)
	}

	// doc.Replace("", , -1)
	// doc.Replace("", , -1)
	// doc.Replace("", , -1)

	// doc.Replace("среднее ☐", "мужской ☒", -1)
}
