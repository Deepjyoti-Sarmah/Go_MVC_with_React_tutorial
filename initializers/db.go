package initializers

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	// dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open("DB"), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database")
	}

}

// func SyncDB() {
// 	DB.AutoMigrate(&models.Post{}) //creates table in our database
// }
