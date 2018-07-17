package models

//Address Model Definition
type Address struct {
	Line1 string `json:"line1"`
	Line2 string `json:"line2"`
	State string `json:"state"`
}
