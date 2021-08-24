package model

type User struct {
	ID      uint `json:"Id" gorm:"primaryKey" `
	Username string `json:"Username"`
	Password string `json:"Password"`
}