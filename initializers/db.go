package initializers

import (
	"fmt"
	"os"

	"github.com/DeepjyotiSarmah/go_MVC_with_React_tutorial/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error
	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database")
	}

}

func SyncDB() {
	DB.AutoMigrate(&models.Post{}) //creates table in our database
}
