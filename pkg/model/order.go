package model

type Order struct {
	Id       string `gorm:"primary_key" json:"id"`
	Uid      string `json:"uid"`
	Start    int32  `json:"start"`
	Dest     int32  `json:"dest"`
	Finished bool   `json:"finished"`
}
