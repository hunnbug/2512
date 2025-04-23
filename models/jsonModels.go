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
	DiplomSeria            int       `json:"diplomSeria"`
	DiplomNumber           int       `json:"diplomNumber"`
	DateGiven              string    `json:"dateGiven"`
	City                   string    `json:"city"`
	Region                 string    `json:"region"`
	EducationalInstitution string    `json:"educationalInstitution"`
	Speciality             string    `json:"speciality"`
	ID_LevelEducation      uuid.UUID `json:"levelEducation"`
}

type PlaceWorkRequest struct {
	NameCompany        string `json:"nameCompany"`
	JobTitle           string `json:"jobTitle"`
	AllExperience      int    `json:"allExperience"`
	JobTitleExpirience int    `json:"jobTitleExpirience"`
}

type ProgramEducationRequest struct {
	NameProfEducation     string    `json:"nameProfEducation"`
	TimeEducation         int       `json:"timeEducation"`
	IndividualPrice       float32   `json:"individualprice"`
	GroupPrice            float32   `json:"groupprice"`
	CampusPrice           float32   `json:"campusprice"`
	ID_DivisionsEducation uuid.UUID `json:"id_divisionseducation"`
	ID_EducationType      uuid.UUID `json:"id_educationtype"`
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
}

type DivisionsEducationRequests struct {
	ID_DivisionsEducation string `json:"id_divisionseducation"`
}

type EnrollmentsRequests struct {
	ID_Listener uuid.UUID `json:"id_listener"`
	ID_Program  uuid.UUID `json:"id_program"`
	StartDate   string    `json:"startdate"`
	EndDate     string    `json:"enddate"`
	// PersonalCard string    `json:"personalcard"`
}

//update

type EducationListenerUpdateRequest struct {
	ID_EducationListener   string    `json:"id_educationlistener"`
	DiplomSeria            int       `json:"diplomSeria"`
	DiplomNumber           int       `json:"diplomNumber"`
	DateGiven              string    `json:"dateGiven"`
	City                   string    `json:"city"`
	Region                 string    `json:"region"`
	EducationalInstitution string    `json:"educationalInstitution"`
	Speciality             string    `json:"speciality"`
	ID_LevelEducation      uuid.UUID `json:"id_levelEducation"`
}

type PlaceWorkUpdateRequest struct {
	ID_Placework       string `json:"id_placework"`
	NameCompany        string `json:"nameCompany"`
	JobTitle           string `json:"jobTitle"`
	AllExperience      int    `json:"allExperience"`
	JobTitleExpirience int    `json:"jobTitleExpirience"`
}

type RegAddressUpdateRequest struct {
	ID_RegAddress string `json:"id_regaddress"`
	MailIndex     int    `json:"mailIndex"`
	Region        string `json:"region"`
	City          string `json:"city"`
	Street        string `json:"street"`
	House         string `json:"house"`
	Building      string `json:"building"`
	Apartment     string `json:"apartment"`
}

type PassportUpdateRequest struct {
	ID_Passport   string `json:"id_passport"`
	PlaceBirth    string `json:"placeBirth"`
	Citizenship   string `json:"citizenship"`
	Gender        string `json:"gender"`
	Seria         int    `json:"seria"`
	Number        int    `json:"number"`
	PassportGiven string `json:"passportGiven"`
	DateGiven     string `json:"dateGiven"`
	Code          string `json:"code"`
}

type UpdateListenerRequest struct {
	FirstName           string                         `json:"firstName"`
	SecondName          string                         `json:"secondName"`
	MiddleName          string                         `json:"middleName"`
	DateOfBirth         string                         `json:"dateOfBirth"`
	SNILS               string                         `json:"snils"`
	ContactPhone        string                         `json:"contactPhone"`
	Email               string                         `json:"email"`
	Passport            PassportUpdateRequest          `json:"passport"`
	RegistrationAddress RegAddressUpdateRequest        `json:"registrationAddress"`
	EducationListener   EducationListenerUpdateRequest `json:"education"`
	PlaceWork           PlaceWorkUpdateRequest         `json:"placeWork"`
}
