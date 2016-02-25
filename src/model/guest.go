package model

type Guest struct {
	Common `bson:",inline"`
	FullName string `json:"fullname"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	RegTime string `json:"regtime"`
	Address string `json:"address"`
	HomeTown string `json:"hometown"`
}