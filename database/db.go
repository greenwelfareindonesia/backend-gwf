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

	// if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); exists == false {
	// 	if err := godotenv.Load(); err != nil {
	// 		log.Fatal("error loading .env file:", err)
	// 	}
	// }

	// dbUsername := os.Getenv("MYSQLUSER")
	// dbPassword := os.Getenv("MYSQLPASSWORD")
	// dbHost := os.Getenv("MYSQLHOST")
	// dbPort := os.Getenv("MYSQLPORT")
	// dbName := os.Getenv("MYSQLDATABASE")
	// // dbUrl := os.Getenv("DATABASE_URL")

	// // Gunakan nilai variabel lingkungan untuk koneksi database
	// // dsn := dbUsername + ":" + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	// fmt.Println("Nilai dbUsername:", dbUsername)
	// // fmt.Println("Nilai dbPassword:", dbPassword)
	// fmt.Println("Nilai dbHost:", dbHost)
	// fmt.Println("Nilai dbPort:", dbPort)
	// fmt.Println("Nilai dbName:", dbName)
	// // fmt.Println("String Koneksi (DSN):", dsn)

	// // dsn := "root:@tcp(127.0.0.1:3306)/mencoba?charset=utf8mb4&parseTime=True&loc=Local"
	// // db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("DB Connection Error")
	// }

	// if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); !exists {
	// 	if err := godotenv.Load(); err != nil {
	// 		log.Fatal("Error loading .env file:", err)
	// 	}
	// }

	// // Get DATABASE_URL from environment variables
	// databaseURL := os.Getenv("DATABASE_URL")
	// if databaseURL == "" {
	// 	log.Fatal("DATABASE_URL is not set in environment variables")
	// }

	// fmt.Println("Database URL:", databaseURL)

	// // Connect to the database using GORM
	// db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("DB Connection Error:", err)
	// }

	// Load environment variables if not in Railway environment
	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); !exists {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	// Get DATABASE_URL from environment variables
	databaseURL := "root:NyeiUNkCgnvPCOVYVnuEkGVwEeFoBPgY@roundhouse.proxy.rlwy.net:17131/railway?charset=utf8mb4&parseTime=True&loc=Local"
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set in environment variables")
	}

	// Print the database URL for debugging purposes
	fmt.Println("Database URL:", databaseURL)

	// Connect to the database using GORM
	db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{})
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

	return db, nil

}
