package dyflexis

import (
	"github.com/cydev/zero"
	"github.com/omniboost/go-dyflexis/omitempty"
)

type Offices []Office

type Office struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Informationstream struct {
	Day struct {
		Value int `json:"value"`
		Hours []struct {
			Hour  int `json:"hour"`
			Value int `json:"value"`
		} `json:"hours"`
	} `json:"day"`
}

type Employee struct {
	Initials              string `json:"initials"`
	FirstName             string `json:"firstName"`
	LastNamePrefix        string `json:"lastNamePrefix"`
	LastName              string `json:"lastName"`
	NameFormat            string `json:"nameFormat"`
	MaritalStatus         string `json:"maritalStatus"`
	PartnerInitials       string `json:"partnerInitials"`
	PartnerFirstName      string `json:"partnerFirstName"`
	PartnerLastNamePrefix string `json:"partnerLastNamePrefix"`
	PartnerLastName       string `json:"partnerLastName"`
	Gender                string `json:"gender"`
	PhoneNumber           string `json:"phoneNumber"`
	PhoneNumber2          string `json:"phoneNumber2"`
	Email                 string `json:"email"`
	DateOfBirth           Date   `json:"dateOfBirth"`
	PlaceOfBirth          string `json:"placeOfBirth"`
	StreetName            string `json:"streetName"`
	StreetNumber          string `json:"streetNumber"`
	PostalCode            string `json:"postalCode"`
	City                  string `json:"city"`
	Nationality           string `json:"nationality"`
	EmploymentStart       string `json:"employmentStart"`
	EmploymentEnd         string `json:"employmentEnd"`
	PersonnelNumber       string `json:"personnelNumber"`
	CostCenter            string `json:"costCenter"`
	ProbationDate         Date   `json:"probationDate,omitempty"`
	EmployerReferenceID   string `json:"employerReferenceId"`
	JobDescription        string `json:"jobDescription"`
}

func (e Employee) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(e)
}

func (e Employee) IsEmpty() bool {
	return zero.IsZero(e)
}

type Contracts []Contract

type Contract struct {
	ContractReference string  `json:"contractReference"`
	Office            int     `json:"office"`
	Type              int     `json:"type"`
	Start             Date    `json:"start"`
	End               Date    `json:"end"`
	Hours             float64 `json:"hours"`
	Days              float64 `json:"days"`
	Salary            float64 `json:"salary"`
}

func (c Contract) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(c)
}

func (c Contract) IsEmpty() bool {
	return zero.IsZero(c)
}

type ContractTypes []ContractType

type ContractType struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	FixedBase bool   `json:"fixedBase"`
}
