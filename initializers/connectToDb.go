package initializers

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	//postgres://pdlmxsuw:3OoWbh-3Ku420uYtji-Kb3PQcGN-2hXy@peanut.db.elephantsql.com/pdlmxsuw
	dsn := os.Getenv("DB_DSN")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}