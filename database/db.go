package database

import (
	"fmt"
	"greenwelfare/artikel"
	"greenwelfare/contact"
	"greenwelfare/ecopedia"
	endpointcount "greenwelfare/endpointCount"
	"greenwelfare/event"
	"greenwelfare/feedback"
	"greenwelfare/gallery"
	"greenwelfare/user"
	"greenwelfare/veganguide"
	"greenwelfare/workshop"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() (*gorm.DB, error){

	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file:", err)
	}
	
	dbUsername, exists := os.LookupEnv("MYSQLUSER")
	if !exists {
		log.Fatal("MYSQLUSER environment variable is not set")
	}
	
	dbPassword, exists := os.LookupEnv("MYSQLPASSWORD")
	if !exists {
		log.Fatal("MYSQLPASSWORD environment variable is not set")
	}
	
	dbHost, exists := os.LookupEnv("MYSQLHOST")
	if !exists {
		log.Fatal("MYSQLHOST environment variable is not set")
	}
	
	dbPort, exists := os.LookupEnv("MYSQLPORT")
	if !exists {
		log.Fatal("MYSQLPORT environment variable is not set")
	}
	
	dbName, exists := os.LookupEnv("MYSQLDATABASE")
	if !exists {
		log.Fatal("MYSQLDATABASE environment variable is not set")
	}
	
	// Gunakan nilai variabel lingkungan untuk koneksi database
	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	
		// dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
		fmt.Println("Nilai dbUsername:", dbUsername)
		fmt.Println("Nilai dbPassword:", dbPassword)
		fmt.Println("Nilai dbHost:", dbHost)
		fmt.Println("Nilai dbPort:", dbPort)
		fmt.Println("Nilai dbName:", dbName)
		fmt.Println("String Koneksi (DSN):", dsn)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("DB Connection Error")
		}

		// Auto Migration
		db.AutoMigrate(&user.User{})
		db.AutoMigrate(&artikel.Artikel{})
		db.AutoMigrate(&ecopedia.Ecopedia{})
		db.AutoMigrate(&contact.Contact{})
		db.AutoMigrate(&workshop.Workshop{})
		db.AutoMigrate(&event.Event{})
		db.AutoMigrate(&veganguide.Veganguide{})
		db.AutoMigrate(&feedback.Feedback{})
		db.AutoMigrate(&endpointcount.Statistics{})
		db.AutoMigrate(&ecopedia.Comment{})
		db.AutoMigrate(&ecopedia.IsLike{})
		db.AutoMigrate(&gallery.Gallery{})

		return db, nil

}