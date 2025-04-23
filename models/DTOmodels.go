package models

import "github.com/google/uuid"

type ListenerIDDTO struct {
	ID_Passport          uuid.UUID
	ID_RegAddress        uuid.UUID
	ID_EducationListener *uuid.UUID
	ID_PlaceWork         *uuid.UUID
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
	ID_ProgramEducation   uuid.UUID
	NameProfEducation     string
	TimeEducation         int
	ID_DivisionsEducation uuid.UUID
	IndividualPrice       float32
	GroupPrice            float32
	CampusPrice           float32
	ID_EducationType      uuid.UUID
}
