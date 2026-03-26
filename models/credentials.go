package models

// Credentials used to authenticate
type Credentials struct {
	Id       string `json:"id"`
	System   string `json:"system"`
	Lang     string `json:"lang"`
	User     string `json:"user"`
	Password string `json:"password"`
	Apitype  string `json:"apitype"`
	Apikey   string `json:"apikey"`
}
