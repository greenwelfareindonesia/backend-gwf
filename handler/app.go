package handler

import (
	"greenwelfare/artikel"
	"greenwelfare/auth"
	"greenwelfare/contact"
	"greenwelfare/database"
	_ "greenwelfare/docs"
	"greenwelfare/ecopedia"
	endpointcount "greenwelfare/endpointCount"
	"greenwelfare/event"
	"greenwelfare/feedback"
	"greenwelfare/gallery"
	"greenwelfare/middleware"
	"greenwelfare/user"
	"greenwelfare/veganguide"
	"greenwelfare/workshop"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartApp() {
	db, err := database.InitDb()
	if err != nil {
		log.Fatal("Eror Db Connection")
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Origin", "Origin , Accept , X-Requested-With , Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization"},
		AllowMethods:    []string{"POST, OPTIONS, GET, PUT, DELETE"},
	}))

	//add sweager
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// user
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := NewUserHandler(userService, authService)

	statisticsRepository := endpointcount.NewStatisticsRepository(db)
	// Inisialisasi service
	statisticsService := endpointcount.NewStatisticsService(statisticsRepository)
	// Inisialisasi handler
	statisticsHandler := NewStatisticsHandler(statisticsService)
	//--//
	user := router.Group("/api/user")
	user.POST("/register", userHandler.RegisterUser)
	user.POST("/login", userHandler.Login)
	user.DELETE("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), userHandler.DeletedUser)
	user.PUT("/:slug", middleware.AuthMiddleware(authService, userService), userHandler.UpdateUser)

	// contact
	contactRepository := contact.NewRepository(db)
	contactService := contact.NewService(contactRepository)
	contactHandler := NewContactHandler(contactService)
	//--//
	con := router.Group("/api/contact")
	con.POST("/", contactHandler.SubmitContactForm)
	con.GET("/", contactHandler.GetContactSubmissionsHandler)
	con.GET("/:slug", contactHandler.GetContactSubmissionHandler)
	con.DELETE("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), contactHandler.DeleteContactSubmissionHandler)

	// workshop
	workshopRepository := workshop.NewRepository(db)
	workshopService := workshop.NewService(workshopRepository)
	workshopHandler := NewWorkshopHandler(workshopService, statisticsService)
	//--//
	work := router.Group("/api/workshop")
	work.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), workshopHandler.CreateWorkshop)
	work.GET("/", workshopHandler.GetAllWorkshop)
	work.GET("/:slug", workshopHandler.GetOneWorkshop)
	work.PUT("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), workshopHandler.UpdateWorkshop)
	work.DELETE("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), workshopHandler.DeleteWorkshop)

	// ecopedia
	ecopediaRepository := ecopedia.NewRepository(db)
	ecopediaService := ecopedia.NewService(ecopediaRepository)
	ecopediaHandler := NewEcopediaHandler(ecopediaService, statisticsService)

	eco := router.Group("/api/ecopedia")
	eco.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), ecopediaHandler.PostEcopediaHandler)
	eco.GET("/", ecopediaHandler.GetAllEcopedia)
	eco.GET("/:slug", ecopediaHandler.GetEcopediaByID)
	eco.PUT("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), ecopediaHandler.UpdateEcopedia)
	eco.DELETE("/:ID", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), ecopediaHandler.DeleteEcopedia)

	// artikel
	artikelRepository := artikel.NewRepository(db)
	artikelService := artikel.NewService(artikelRepository)
	artikelHandler := NewArtikelHandler(artikelService, statisticsService)
	//--//
	art := router.Group("/api/article")
	art.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), artikelHandler.CreateArtikel)
	art.GET("/", artikelHandler.GetAllArtikel)
	art.GET("/:slug", artikelHandler.GetOneArtikel)
	art.PUT("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), artikelHandler.UpdateArtikel)
	art.DELETE("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), artikelHandler.DeleteArtikel)

	// event
	eventRepository := event.NewRepository(db)
	eventService := event.NewService(eventRepository)
	eventHandler := NewEventHandler(eventService, statisticsService)
	//--//
	eve := router.Group("/api/event")
	eve.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), eventHandler.CreateEvent)
	eve.GET("/", eventHandler.GetAllEvent)
	eve.GET("/:slug", eventHandler.GetOneEvent)
	eve.PUT(":slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), eventHandler.UpdateEvent)
	eve.DELETE("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), eventHandler.DeleteEvent)

	// veganguide
	veganguideRepository := veganguide.NewRepository(db)
	veganguideService := veganguide.NewService(veganguideRepository)
	veganguideHandler := NewVeganguideHandler(veganguideService, statisticsService)
	//--//
	veg := router.Group("/api/veganguide")
	veg.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), veganguideHandler.PostVeganguideHandler)
	veg.GET("/", veganguideHandler.GetAllVeganguide)
	veg.GET("/:slug", veganguideHandler.GetVeganguideByID)
	veg.PUT("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), veganguideHandler.UpdateVeganguide)
	veg.DELETE("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), veganguideHandler.DeleteVeganguide)

	// feedback
	feedbackRepository := feedback.NewRepository(db)
	feedbackService := feedback.NewService(feedbackRepository)
	feedbackHandler := NewFeedbackHandler(feedbackService)
	//--//
	fee := router.Group("/api/feedback")
	fee.POST("/", feedbackHandler.PostFeedbackHandler)
	fee.GET("/", feedbackHandler.GetAllFeedback)
	fee.GET("/:slug", feedbackHandler.GetFeedbackBySlug)
	// fee.PUT("/:id", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), feedbackHandler.UpdateFeedback)
	fee.DELETE("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), feedbackHandler.DeleteFeedback)

	galleryRepository := gallery.NewRepository(db)
	galleryService := gallery.NewService(galleryRepository)
	galleryHandler := NewGalleryHandler(galleryService, statisticsService)
	gallerys := router.Group("/api/gallery")
	gallerys.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), galleryHandler.CreateGallery)
	gallerys.GET("/", galleryHandler.GetAllGallery)
	gallerys.GET("/:slug", galleryHandler.GetOneGallery)
	gallerys.PUT("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), galleryHandler.UpdateGallery)
	gallerys.DELETE("/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), galleryHandler.DeleteGallery)

	// statistics
	router.GET("/statistics", statisticsHandler.GetStatisticsHandler)

	// Port
	router.Run(":8080")
}
