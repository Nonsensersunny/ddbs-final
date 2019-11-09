package model

type User struct {
	Id       string `gorm:"primary_key" json:"id"`
	Area     int32  `json:"area"`
	Password string `json:"password"`
}