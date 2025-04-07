package models

import "github.com/google/uuid"

type PassportRequest struct {
	PlaceBirth    string `json:"placeBirth"`
	Citizenship   string `json:"citizenship"`
	Gender        string `json:"gender"`
	Seria         int    `json:"seria"`
	Number        int    `json:"number"`
	PassportGiven string `json:"passportGiven"`
	DateGiven     string `json:"dateGiven"`
	Code          string `json:"code"`
}

type RegAddressRequest struct {
	MailIndex int    `json:"mailIndex"`
	Region    string `json:"region"`
	City      string `json:"city"`
	Street    string `json:"street"`
	House     string `json:"house"`
	Building  string `json:"building"`
	Apartment string `json:"apartment"`
}

type ListenerEducationRequest struct {
	Education string `json:"education"`
}

type EducationListenerRequest struct {
	DiplomSeria            int    `json:"diplomSeria"`
	DiplomNumber           int    `json:"diplomNumber"`
	DateGiven              string `json:"dateGiven"`
	City                   string `json:"city"`
	Region                 string `json:"region"`
	EducationalInstitution string `json:"educationalInstitution"`
	Speciality             string `json:"speciality"`
	LevelEducation         string `json:"levelEducation"`
}

type PlaceWorkRequest struct {
	NameCompany        string `json:"nameCompany"`
	JobTitle           string `json:"jobTitle"`
	AllExperience      int    `json:"allExperience"`
	JobTitleExpirience int    `json:"jobTitleExpirience"`
}

type ProgramRequest struct {
	NameProfEducation string `json:"nameProfEducation"`
	TypeOfEducation   string `json:"typeOfEducation"`
	TimeEducation     int    `json:"timeEducation"`
}

type CreateListenerRequest struct {
	FirstName           string                   `json:"firstName"`
	SecondName          string                   `json:"secondName"`
	MiddleName          string                   `json:"middleName"`
	DateOfBirth         string                   `json:"dateOfBirth"`
	SNILS               string                   `json:"snils"`
	ContactPhone        string                   `json:"contactPhone"`
	Email               string                   `json:"email"`
	Passport            PassportRequest          `json:"passport"`
	RegistrationAddress RegAddressRequest        `json:"registrationAddress"`
	EducationListener   EducationListenerRequest `json:"education"`
	PlaceWork           PlaceWorkRequest         `json:"placeWork"`
	ProgramEducation    ProgramRequest           `json:"programEducation"`
}

//
//DTO (которое можно перенести в отдельный файл)
//

type ListenerIDDTO struct {
	ID_Passport          uuid.UUID
	ID_RegAddress        uuid.UUID
	ID_EducationListener uuid.UUID
	ID_PlaceWork         uuid.UUID
}

type ListenerDTO struct {
	FirstName    string
	SecondName   string
	MiddleName   string
	DateOfBirth  string
	SNILS        string
	ContactPhone string
	Email        string
}

type PassportDTO struct {
	PlaceBirth    string
	Citizenship   string
	Gender        string
	Seria         int
	Number        int
	PassportGiven string
	DateGiven     string
	Code          string
}

type RegistrationAddressDTO struct {
	MailIndex int
	Region    string
	City      string
	Street    string
	House     string
	Building  string
	Apartment string
}

type LevelEducationDTO struct {
	Education string
}

type EducationListenerDTO struct {
	ID_EducationListener   string
	DiplomSeria            int
	DiplomNumber           int
	DateGiven              string
	City                   string
	Region                 string
	EducationalInstitution string
	Speciality             string
	LevelEducation         string
}

type PlaceWorkDTO struct {
	NameCompany        string
	JobTitle           string
	AllExperience      int
	JobTitleExpirience int
}

type ProgramEducationDTO struct {
	NameProfEducation string
	TypeOfEducation   string
	TimeEducation     int
}
