package model

type Address struct {
	Country  string `json:"country"`
	City     string `json:"city"`
	Street   string `json:"street"`
	Building int    `json:"building"`
	Entrance int    `json:"entrance"`
	ZipCode  string `json:"zip_code"`
}
