package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Client struct {
	Username string
	Password string
	Host     string
	Port     int
	DbName   string
}

func GetMySQLDB(client Client) *gorm.DB {
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		client.Username,
		client.Password,
		client.Host,
		client.Port,
		client.DbName)
	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func AutoMigrate(db *gorm.DB, objs ...*interface{}) {
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")
	if err := db.AutoMigrate(&objs); err != nil {
		panic(err.Error)
	}
}
