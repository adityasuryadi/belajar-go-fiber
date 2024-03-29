package config

import (
	"go-blog/exception"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type Testing interface {
// 	SetupRouter() *
// }

func NewTestPostgresDB(configuration Config) *gorm.DB {
	// host := configuration.Get("POSTGRE_HOST")
	// user := configuration.Get("POSTGRE_USER")
	// password := configuration.Get("POSTGRE_PASSWORD")
	// port := configuration.Get("POSTGRE_PORT")
	// db_name := configuration.Get("POSTGRE_DB_NAME")

	host := "localhost"
	user := "postgres"
	password := "postgres"
	port := "5433"
	db_name := "blog_test"

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + db_name + " port=" + port + " sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	exception.PanicIfNeeded(err)

	return db
}
