package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBconfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBconfig() *DBconfig {
	dbconfig := DBconfig{
		Host:     "localhost",
		Port:     3306,
		User:     "root",
		Password: "1234",
		DBName:   "first_go",
	}
	return &dbconfig
}
func DbURL(dbconfig *DBconfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbconfig.User,
		dbconfig.Password,
		dbconfig.Host,
		dbconfig.Port,
		dbconfig.DBName,
	)
}
