package models

// Realtor model definition
type Realtor struct {
	ObjectType   string   `json:"docType"`
	BuisnessName string   `json:"buisnessName"`
	Address      Address  `json:"address"`
	IsRegistered bool     `json:"isRegistered"`
	Agents       []string `json:"agents"`
}
