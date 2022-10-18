package initializers

import "mygram/models"

func SyncToDb() {
	DB.AutoMigrate(&models.User{})
}