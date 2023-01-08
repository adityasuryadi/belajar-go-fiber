package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	processENV()

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// db, err := gorm.Open("mysql", dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=true")
	// dsn := "user:pass@tcp(127.0.0.1:3306)/"+dbName+"?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		logrus.Error("Cannot Connect to MySQL DB")
		panic(err)
	}

	// migrateDDL(db)

	return db
}

func processENV() {

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Error("Error loading env file")
	}
}
