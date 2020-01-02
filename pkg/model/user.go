package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Id       string `gorm:"primary_key" json:"id"`
	Phone    string `json:"phone"`
	Area     string `json:"area"`
	Password string `json:"password"`
	Role     int32  `json:"role"`
}

func (u *User) CreateUser(db *gorm.DB) error {
	u.Role = 1
	return db.Table("users").Create(&u).Error
}

func (u *User) CreateAdmin(db *gorm.DB) error {
	u.Role = 2
	return db.Table("users").Create(&u).Error
}

func GetUserById(db *gorm.DB, id string) (*User, error) {
	u := &User{}
	err := db.Table("users").Where("id = ?", id).Scan(&u).Error
	return u, err
}

func GetUsersByArea(db *gorm.DB, area string) ([]*User, error) {
	var users []*User
	err := db.Table("users").Where("area = ?", area).Scan(&users).Error
	return users, err
}

func UserExistsWithPhoneAndPwd(db *gorm.DB, phone, password string) bool {
	u := &User{}
	err := db.Table("users").Where("phone = ? and password = ?", phone, password).Scan(u).Error
	if err != nil || u.Area == "" {
		return false
	}
	return true
}

func GetUserByPhoneAndPwd(db *gorm.DB, phone, password string) (*User, error) {
	u := &User{}
	err := db.Table("users").Where("phone = ? and password = ?", phone, password).Scan(u).Error
	return u, err
}

func UserExistsWithPhone(db *gorm.DB, phone string) bool {
	u := &User{}
	db.Table("users").Where("phone = ?", phone).Scan(u)
	if u.Id != "" {
		return true
	}
	return false
}
