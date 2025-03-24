package models

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
