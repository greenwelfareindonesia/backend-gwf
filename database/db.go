package database

import (
	"fmt"
	endpointcount "greenwelfare/endpointCount"
	"greenwelfare/entity"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error) {

	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); exists == false {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file:", err)
		}
	}

	// dbUsername := os.Getenv("MYSQLUSER")
	// dbPassword := os.Getenv("MYSQLPASSWORD")
	// dbHost := os.Getenv("MYSQLHOST")
	// dbPort := os.Getenv("MYSQLPORT")
	// dbName := os.Getenv("MYSQLDATABASE")
	dbUrl := os.Getenv("DATABASE_URL")

	// Connect to the database using GORM
	db, err := gorm.Open(mysql.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB Connection Error: %v", err)
	}

	// Further database operations
	fmt.Println("Database connection successful")
	
	// Auto Migration

	db.AutoMigrate(&entity.Artikel{})
	db.AutoMigrate(&entity.Ecopedia{})
	db.AutoMigrate(&entity.Contact{})
	db.AutoMigrate(&entity.Workshop{})
	db.AutoMigrate(&entity.Event{})
	db.AutoMigrate(&entity.Veganguide{})
	db.AutoMigrate(&entity.Feedback{})
	db.AutoMigrate(&endpointcount.Statistics{})
	db.AutoMigrate(&entity.Gallery{})
	db.AutoMigrate(&entity.GalleryImages{})
	db.AutoMigrate(&entity.EcopediaImage{})
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Product{})
	db.AutoMigrate(&entity.ShoppingCart{})
	db.AutoMigrate(&entity.Banner{})
	db.AutoMigrate(&entity.Hrd{})
	db.AutoMigrate(&entity.Kajian{})
	db.AutoMigrate(&entity.KajianImage{})

	return db, nil

}
