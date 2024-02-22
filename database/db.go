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

func InitDb() (*gorm.DB, error) {

	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); exists == false {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file:", err)
		}
	}

	dbUsername := os.Getenv("MYSQLUSER")
	// dbPassword := os.Getenv("MYSQLPASSWORD")
	dbHost := os.Getenv("MYSQLHOST")
	dbPort := os.Getenv("MYSQLPORT")
	dbName := os.Getenv("MYSQLDATABASE")

	// Gunakan nilai variabel lingkungan untuk koneksi database
	dsn := dbUsername + ":" + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	// dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println("Nilai dbUsername:", dbUsername)
	// fmt.Println("Nilai dbPassword:", dbPassword)
	fmt.Println("Nilai dbHost:", dbHost)
	fmt.Println("Nilai dbPort:", dbPort)
	fmt.Println("Nilai dbName:", dbName)
	fmt.Println("String Koneksi (DSN):", dsn)

	// dsn := "root:@tcp(127.0.0.1:3306)/mencoba?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	// Auto Migration

	db.AutoMigrate(&artikel.Artikel{})
	db.AutoMigrate(&ecopedia.Ecopedia{})
	db.AutoMigrate(&contact.Contact{})
	db.AutoMigrate(&workshop.Workshop{})
	db.AutoMigrate(&event.Event{})
	db.AutoMigrate(&veganguide.Veganguide{})
	db.AutoMigrate(&feedback.Feedback{})
	db.AutoMigrate(&endpointcount.Statistics{})
	db.AutoMigrate(&gallery.Gallery{})
	db.AutoMigrate(&gallery.GalleryImages{})
	db.AutoMigrate(&ecopedia.EcopediaImage{})
	db.AutoMigrate(&user.User{})

	return db, nil

}
