package models

//Lawyer model definition
type Lawyer struct {
	ObjectType   string `json:"docType"`
	FullName     string `json:"fullName"`
	BuisnessName string `json:"buisnessName"`
	IsLicensed   string `json:"isLicensed"`
}
