package config

import (
	// "database/sql"
	"fmt"
	"goguru/helper"

	// _ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "root"
	password = "Dar@1611#"
	dbName   = "DarshilDb"
)

func DatabaseConnection() *gorm.DB {
	// sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	// // dsn := "root:Dar@1611#@tcp(localhost:3306)/DarshilDb"
	// sqlDb, err := sql.Open("mysql", sqlInfo)
	// helper.ErrorPanic(err)
	// db, err := gorm.Open(sqlDb, &gorm.Config{})
	// helper.ErrorPanic(err)

	// return db
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
