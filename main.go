package main

import (
	"greenwelfare/artikel"
	"greenwelfare/auth"
	"greenwelfare/contact"
	_ "greenwelfare/docs"
	"greenwelfare/ecopedia"
	endpointcount "greenwelfare/endpointCount"
	"greenwelfare/event"
	"greenwelfare/feedback"
	"greenwelfare/handler"
	"greenwelfare/middleware"
	"greenwelfare/user"
	"greenwelfare/veganguide"
	"greenwelfare/workshop"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// @title Sweager Service API
// @description Sweager service API in Go using Gin framework
// @host https://backend-gwf-production.up.railway.app/

// @securitydefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	
	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); exists == false {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file:", err)
		}
	}

	dbUsername := os.Getenv("MYSQLUSER")
	dbPassword := os.Getenv("MYSQLPASSWORD")
	dbHost := os.Getenv("MYSQLHOST")
	dbPort := os.Getenv("MYSQLPORT")
	dbName := os.Getenv("MYSQLDATABASE")
	// secretKey := os.Getenv("SECRET_KEY")

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders: []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Access-Control-Allow-Origin , Origin , Accept , X-Requested-With , Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization"},
		AllowMethods: []string{"POST, OPTIONS, GET, PUT, DELETE"},
	  }))

	//add sweager
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// user
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	statisticsRepository := endpointcount.NewStatisticsRepository(db)
	// Inisialisasi service
	statisticsService := endpointcount.NewStatisticsService(statisticsRepository)
	// Inisialisasi handler
	statisticsHandler := handler.NewStatisticsHandler(statisticsService)
	//--//
	user := router.Group("/user")
	user.POST("/register", userHandler.RegisterUser)
	user.POST("/login", userHandler.Login)
	user.POST("/email_checkers", userHandler.CheckEmailAvailabilty)
	user.DELETE("/", middleware.AuthMiddleware(authService, userService), userHandler.DeletedUser)
	user.PUT("/:id", middleware.AuthMiddleware(authService, userService), userHandler.UpdateUser)

	// contact
	contactRepository := contact.NewRepository(db)
	contactService := contact.NewService(contactRepository)
	contactHandler := handler.NewContactHandler(contactService)
	//--//
	con := router.Group("/contact")
	con.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), contactHandler.SubmitContactForm)
	con.GET("/", contactHandler.GetContactSubmissionsHandler)
	con.GET("/:id", contactHandler.GetContactSubmissionHandler)
	con.DELETE("/:id", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), contactHandler.DeleteContactSubmissionHandler)

	// workshop
	workshopRepository := workshop.NewRepository(db)
	workshopService := workshop.NewService(workshopRepository)
	workshopHandler := handler.NewWorkshopHandler(workshopService, statisticsService)
	//--//
	work := router.Group("/workshop")
	work.POST("/", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), workshopHandler.CreateWorkshop)
	work.GET("/", workshopHandler.GetAllWorkshop)
	work.GET("/:id", workshopHandler.GetOneWorkshop)
	work.PUT("/:id", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), workshopHandler.UpdateWorkshop)
	work.DELETE("/:id", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), workshopHandler.DeleteWorkshop)

	// ecopedia
	ecopediaRepository := ecopedia.NewRepository(db)
	ecopediaService := ecopedia.NewService(ecopediaRepository)
	ecopediaHandler := handler.NewEcopediaHandler(ecopediaService, statisticsService)

	eco := router.Group("/ecopedia")
	eco.POST("/", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), ecopediaHandler.PostEcopediaHandler)
	eco.GET("/", ecopediaHandler.GetAllEcopedia)
	eco.GET("/:id", ecopediaHandler.GetEcopediaByID)
	eco.PUT("/:id", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), ecopediaHandler.UpdateEcopedia)
	eco.DELETE("/:id", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), ecopediaHandler.DeleteEcopedia)
	eco.POST("comment/:id", middleware.AuthMiddleware(authService, userService), ecopediaHandler.PostCommentEcopedia)

	// artikel
	artikelRepository := artikel.NewRepository(db)
	artikelService := artikel.NewService(artikelRepository)
	artikelHandler := handler.NewArtikelHandler(artikelService, statisticsService)
	//--//
	art := router.Group("/artikel")
	art.POST("/", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), artikelHandler.CreateArtikel)
	art.GET("/", artikelHandler.GetAllArtikel)
	art.GET("/:id", artikelHandler.GetOneArtikel)
	art.PUT("/:id", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), artikelHandler.UpdateArtikel)
	art.DELETE("/:id", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), artikelHandler.DeleteArtikel)

	// event
	eventRepository := event.NewRepository(db)
	eventService := event.NewService(eventRepository)
	eventHandler := handler.NewEventHandler(eventService, statisticsService)
	//--//
	eve := router.Group("/event")
	eve.POST("/", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), eventHandler.CreateEvent)
	eve.GET("/", eventHandler.GetAllEvent)
	eve.GET("/:id", eventHandler.GetOneEvent)
	eve.PUT(":id", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), eventHandler.UpdateEvent)
	eve.DELETE("/:id", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), eventHandler.DeleteEvent)

	// veganguide
	veganguideRepository := veganguide.NewRepository(db)
	veganguideService := veganguide.NewService(veganguideRepository)
	veganguideHandler := handler.NewVeganguideHandler(veganguideService, statisticsService)
	//--//
	veg := router.Group("/veganguide")
	veg.POST("/", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), veganguideHandler.PostVeganguideHandler)
	veg.GET("/", veganguideHandler.GetAllVeganguide)
	veg.GET("/:id", veganguideHandler.GetVeganguideByID)
	veg.PUT("/:id", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), veganguideHandler.UpdateVeganguide)
	veg.DELETE("/:id", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), veganguideHandler.DeleteVeganguide)

	// feedback
	feedbackRepository := feedback.NewRepository(db)
	feedbackService := feedback.NewService(feedbackRepository)
	feedbackHandler := handler.NewFeedbackHandler(feedbackService)
	//--//
	fee := router.Group("/feedback")
	fee.POST("/", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), feedbackHandler.PostFeedbackHandler)
	fee.GET("/", feedbackHandler.GetAllFeedback)
	fee.GET("/:id", feedbackHandler.GetFeedbackByID)
	fee.PUT("/:id", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), feedbackHandler.UpdateFeedback)
	fee.DELETE("/:id", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), feedbackHandler.DeleteFeedback)

	// statistics
	router.GET("/statistics", statisticsHandler.GetStatisticsHandler)

	// Port
	router.Run(":8080")
}

