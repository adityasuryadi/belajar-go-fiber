package config

import (
	"go-blog/exception"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(configuration Config) *gorm.DB {
	host := configuration.Get("POSTGRE_HOST")
	user := configuration.Get("POSTGRE_USER")
	password := configuration.Get("POSTGRE_PASSWORD")
	port := configuration.Get("POSTGRE_PORT")
	db_name := configuration.Get("POSTGRE_DB_NAME")

	// host := "postgres"
	// user := "postgres"
	// password := "postgres"
	// port := "5432"
	// db_name := "blog"

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + db_name + " port=" + port + " sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	exception.PanicIfNeeded(err)

	return db
}
