package mysql

import (
	"ddbs-final/internal/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetMySQLDB(c *config.MySQLConf) *gorm.DB {
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.DBName)
	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func MigrateTables(db *gorm.DB, objs ...interface{}) {
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")
	if err := db.AutoMigrate(objs...).Error; err != nil {
		panic(err.Error)
	}
}
