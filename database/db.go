package database

// func InitDb() (*gorm.DB, error){

// 	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); exists == false {
// 		if err := godotenv.Load(); err != nil {
// 			log.Fatal("error loading .env file:", err)
// 		}
// 	}

// 		dbUsername := os.Getenv("MYSQLUSER")
// 		dbPassword := os.Getenv("MYSQLPASSWORD")
// 		dbHost := os.Getenv("MYSQLHOST")
// 		dbPort := os.Getenv("MYSQLPORT")
// 		dbName := os.Getenv("MYSQLDATABASE")

// 		dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
// 		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 		if err != nil {
// 			return nil, err
// 		}

// 		// Auto Migration
// 		db.AutoMigrate(&user.User{})
// 		db.AutoMigrate(&artikel.Artikel{})
// 		db.AutoMigrate(&ecopedia.Ecopedia{})
// 		db.AutoMigrate(&contact.Contact{})
// 		db.AutoMigrate(&workshop.Workshop{})
// 		db.AutoMigrate(&event.Event{})
// 		db.AutoMigrate(&veganguide.Veganguide{})
// 		db.AutoMigrate(&feedback.Feedback{})
// 		db.AutoMigrate(&endpointcount.Statistics{})
// 		db.AutoMigrate(&ecopedia.Comment{})
// 		db.AutoMigrate(&ecopedia.IsLike{})
// 		db.AutoMigrate(&gallery.Gallery{})

// 		return db, nil

// }