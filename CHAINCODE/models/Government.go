package models

//Government model definition
type Government struct {
	ObjectType   string `json:"docType"`
	State        string `json:"state"`
	GovernorName string `json:"governorName"`
}
