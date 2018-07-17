package models

//PropertySpecification model definition
type PropertySpecification struct {
	ObjectType        string `json:"docType"`
	NumberOfRooms     int    `json:"numberOfRooms"`
	NumberOfToilets   int    `json:"numberOfToilets"`
	NumberLivingRooms int    `json:"numberOflivingRooms"`
	NumberOfKitchens  int    `json:"numberOfKitchens"`
}
