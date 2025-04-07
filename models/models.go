package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Username string
	Password string
}

type ErrorResponse struct {
	Err     error
	Message string
}

type Passport struct {
	ID_Passport   uuid.UUID `gorm:"primaryKey;column:id_passport"`
	PlaceBirth    string    `gorm:"column:placebirth"`
	Citizenship   string    `gorm:"column:citizenship"`
	Gender        string    `gorm:"column:gender"`
	Seria         int       `gorm:"column:seria"`
	Number        int       `gorm:"column:number;unique"`
	PassportGiven string    `gorm:"column:passportgiven"`
	DateGiven     string    `gorm:"column:dategiven;type:date"`
	Code          string    `gorm:"column:code"`
}

func (Passport) TableName() string {
	return "passport"
}

type RegistrationAddress struct {
	ID_regAddress uuid.UUID `gorm:"column:id_regaddress;primaryKey"`
	MailIndex     int       `gorm:"column:mailindex"`
	Region        string    `gorm:"column:region"`
	City          string    `gorm:"column:city"`
	Street        string    `gorm:"column:street"`
	House         string    `gorm:"column:house"`
	Building      string    `gorm:"column:building"`
	Apartment     string    `gorm:"column:apartment"`
}

func (RegistrationAddress) TableName() string {
	return "registrationaddress"
}

type LevelEducation struct {
	ID_LevelEducation uuid.UUID `gorm:"column:id_leveleducation;primaryKey"`
	Education         string    `gorm:"column:education"`
}

func (LevelEducation) TableName() string {
	return "leveleducation"
}

type EducationListener struct {
	ID_EducationListener   uuid.UUID `gorm:"column:id_educationlistener;primaryKey"`
	DiplomSeria            int       `gorm:"column:diplomseria;unique"`
	DiplomNumber           int       `gorm:"column:diplomnumber;unique"`
	DateGiven              string    `gorm:"column:dategiven;type:date"`
	City                   string    `gorm:"column:city"`
	Region                 string    `gorm:"column:region"`
	EducationalInstitution string    `gorm:"column:educationalinstitution"`
	Speciality             string    `gorm:"column:speciality"`
	ID_LevelEducation      uuid.UUID `gorm:"column:id_leveleducation"`
}

func (EducationListener) TableName() string {
	return "educationlistener"
}

type PlaceWork struct {
	ID_PlaceWork       uuid.UUID `gorm:"column:id_placework;primaryKey"`
	NameCompany        string    `gorm:"column:namecompany"`
	JobTitle           string    `gorm:"column:jobtitle"`
	AllExperience      int       `gorm:"column:allexperience"`
	JobTitleExpirience int       `gorm:"column:jobtitleexperience"`
}

func (PlaceWork) TableName() string {
	return "placework"
}

type DivisionsEducation struct {
	ID_DivisionsEducation uuid.UUID `gorm:"column:id_divisionseducation;primaryKey"`
	Divisions             string    `gorm:"column:divisions;unique"`
}

func (DivisionsEducation) TableName() string {
	return "divisionseducation"
}

type ProgramEducation struct {
	ID_ProgramEducation   uuid.UUID          `gorm:"column:id_programeducation;primaryKey"`
	NameProfEducation     string             `gorm:"column:nameprofeducation"`
	TypeOfEducation       string             `gorm:"column:typeofeducation"`
	TimeEducation         int                `gorm:"column:timeeducation"`
	ID_DivisionsEducation uuid.UUID          `gorm:"column:id_divisionseducation"`
	Division              DivisionsEducation `gorm:"foreignKey:ID_DivisionsEducation"`
}

func (ProgramEducation) TableName() string {
	return "programeducation"
}

type Listener struct {
	ID_Listener          uuid.UUID           `gorm:"column:id_listener;primaryKey"`
	FirstName            string              `gorm:"column:firstname"`
	SecondName           string              `gorm:"column:secondname"`
	MiddleName           string              `gorm:"column:middlename"`
	DateOfBirth          string              `gorm:"column:dateofbirth;type:date"`
	SNILS                string              `gorm:"column:snils;unique"`
	ContactPhone         string              `gorm:"column:contactphone;unique"`
	Email                string              `gorm:"column:email;unique"`
	ID_passport          uuid.UUID           `gorm:"column:id_passport;"`
	Passport             Passport            `gorm:"foreignKey:ID_passport"`
	ID_regAddress        uuid.UUID           `gorm:"column:id_regaddress"`
	RegistrationAddress  RegistrationAddress `gorm:"foreignKey:ID_regAddress"`
	ID_EducationListener uuid.UUID           `gorm:"column:id_educationlistener"`
	EducationListener    EducationListener   `gorm:"foreignKey:ID_EducationListener"`
	ID_PlaceWork         uuid.UUID           `gorm:"column:id_placework"`
	PlaceWork            PlaceWork           `gorm:"foreignKey:ID_PlaceWork"`
}

func (Listener) TableName() string {
	return "listener"
}

type ListenerProgramEducation struct {
	ID_Listener         uuid.UUID        `gorm:"primaryKey;column:id_listener"`
	ID_ProgramEducation uuid.UUID        `gorm:"primaryKey;column:id_programeducation"`
	Listener            Listener         `gorm:"foreignKey:ID_Listener"`
	ProgramEducation    ProgramEducation `gorm:"foreignKey:ID_ProgramEducation"`
}

func (ListenerProgramEducation) TableName() string {
	return "listenerprogrameducation"
}
