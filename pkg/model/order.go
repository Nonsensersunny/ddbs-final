package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct {
	Id       string    `gorm:"primary_key" json:"id"`
	Uid      string    `json:"uid"`
	Start    string    `json:"start"`
	Dest     string    `json:"dest"`
	Finished bool      `json:"finished"`
	Time     time.Time `gorm:"default:now()" json:"time"`
}

func (o *Order) CreateOrder(db *gorm.DB) (string, error) {
	err := db.Table("orders").Create(&o).Error
	return o.Id, err
}

func GetOrderById(db *gorm.DB, id string) (*Order, error) {
	o := &Order{}
	err := db.Table("orders").Where("id = ?", id).Scan(o).Error
	return o, err
}

func GetOrdersByUid(db *gorm.DB, uid string) ([]*Order, error) {
	var orders []*Order
	err := db.Table("orders").Where("uid = ?", uid).Scan(&orders).Error
	return orders, err
}

func GetAllOrders(db *gorm.DB) ([]*Order, error) {
	var orders []*Order
	err := db.Table("orders").Find(&orders).Error
	return orders, err
}

func FinishOrder(db *gorm.DB, id string) error {
	fmt.Println(id)
	return db.Table("orders").Debug().Where("id = ?", id).Update("finished", true).Error
}
