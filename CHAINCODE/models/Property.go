package models

//Property model definition
type Property struct {
	ObjectType            string                `json:"docType"`
	PropertyType          string                `json:"propertyType"`
	Area                  string                `json:"area"`
	Owner                 string                `json:"owner"`
	Size                  float32               `json:"size"`
	Price                 float32               `json:"price"`
	IsGovermentOwned      bool                  `json:"isGovermentOwned"`
	Tyr                   string                `json:"tyr"`
	PropertySpecification PropertySpecification `json:"propertySpecification"`
}
