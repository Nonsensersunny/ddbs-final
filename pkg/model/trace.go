package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Trace struct {
	Id      string    `gorm:"primary_key" json:"id"`
	Time    time.Time `gorm:"default:now()" json:"time"`
	Oid     string    `json:"oid"`
	Current string    `json:"current"`
	Next    string    `json:"next"`
}

func (t *Trace) CreateTrace(db *gorm.DB) error {
	return db.Table("traces").Create(&t).Error
}

func GetTraceById(db *gorm.DB, id string) (trace *Trace, err error) {
	err = db.Table("traces").Where("id = ?", id).Scan(&trace).Error
	return
}

func (t *Trace) UpdateNext(db *gorm.DB, next string) error {
	return db.Table("traces").Where("id = ?", t.Id).Update("next", next).Error
}

func GetByOid(db *gorm.DB, oid string) (*Trace, error) {
	trace := &Trace{}
	err := db.Table("traces").Where("oid = ?", oid).Scan(trace).Error
	return trace, err
}
