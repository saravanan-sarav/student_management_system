package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sms/config"
	"sms/db/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type DB struct {
	*gorm.DB
}

func Init() {
	sqlDB, err := sql.Open("mysql", config.Conf.DatabaseURL)

	if err != nil {
		log.Fatal("Unable to connect Database", err)
		os.Exit(1)
	}

	sqlDB.SetConnMaxIdleTime(5)
	sqlDB.SetMaxOpenConns(config.Conf.MaxDBConn)
	sqlDB.SetConnMaxLifetime(10)

	db, err = gorm.Open(mysql.Open(config.Conf.DatabaseURL), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to create GORM content", err)
	}
	fmt.Println("DB CONNECTED SUCCESSFULLY")

	err = autoMigrate()

	if err != nil {
		log.Fatal("GORM Unable to create tables. Auto Migrate Failed.", err)
	}

	fmt.Println("AUTOMIGRATED SUCCESSFULLY")

}

func autoMigrate() error {
	return db.AutoMigrate(&models.Student{}, &models.Auth{}, &models.AttendanceEntry{}, &models.WorkDoneEntry{}, &models.LeaveEntry{})
}

func New() *DB {
	return &DB{
		DB: db,
	}
}
