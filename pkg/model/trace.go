package model

import "time"

type Trace struct {
	Id      string    `gorm:"primary_key" json:"id"`
	Time    time.Time `json:"time"`
	Oid     string    `json:"oid"`
	Current int32     `json:"current"`
	Next    int32     `json:"next"`
}
