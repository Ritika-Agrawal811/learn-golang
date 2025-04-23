package config

import (
	"fmt"
	"log"

	"github.com/Ritika-Agrawal811/learn-golang/project-book-management/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbInstance *gorm.DB
)

func Connect() {
	dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=true&loc=Local",
		utils.GetEnv("DATABASE_USER", "user"),
		utils.GetEnv("DATABASE_PASSWORD", "password"),
		utils.GetEnv("DATABASE_NAME", "dbname"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to create db instance", err)
	}

	dbInstance = db
}

func GetDatabaseInstance() *gorm.DB {
	if dbInstance == nil {
		log.Fatal("Database instance is not initialized yet")
	}

	return dbInstance
}

func Close() {
	if dbInstance != nil {
		sqlDB, err := dbInstance.DB()
		if err != nil {
			log.Println("Error getting raw DB instance:", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Println("Error closing DB connection:", err)
		}
	}
}
