package models

//Buyer model Definition
type Buyer struct {
	ObjectType string  `json:"docType"`
	LastName   string  `json:"lastName"`
	FirstName  string  `json:"firstName"`
	Age        int     `json:"age"`
	Gender     string  `json:"Gender"`
	Email      string  `json:"email"`
	Guarantor  string  `json:"guarantor"`
	Occupation string  `json:"occupation"`
	Address    Address `json:"address"`
}
