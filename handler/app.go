package handler

import (
	"greenwelfare/auth"
	"greenwelfare/database"
	_ "greenwelfare/docs"
	endpointcount "greenwelfare/endpointCount"
	"greenwelfare/middleware"
	"greenwelfare/repository"
	"greenwelfare/service"
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
	userRepository := repository.NewRepositoryUser(db)
	userService := service.NewServiceUser(userRepository)
	authService := auth.NewService()
	userHandler := NewUserHandler(userService, authService)

	statisticsRepository := endpointcount.NewStatisticsRepository(db)
	// Inisialisasi service
	statisticsService := endpointcount.NewStatisticsService(statisticsRepository)
	// Inisialisasi handler
	// statisticsHandler := NewStatisticsHandler(statisticsService)
	//--//
	user := router.Group("/api/user")
	user.POST("/register", userHandler.RegisterUser)
	user.POST("/login", userHandler.Login)
	user.DELETE("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), userHandler.DeletedUser)
	user.PUT("/:slug", middleware.AuthMiddleware(authService, userService), userHandler.UpdateUser)

	{
		// contact
		contactRepository := repository.NewRepositoryContact(db)
		contactService := service.NewServiceContact(contactRepository)
		contactHandler := NewContactHandler(contactService)
		//--//
		con := router.Group("/api/contact")
		con.POST("/", contactHandler.SubmitContactForm)
		con.GET("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), contactHandler.GetContactSubmissionsHandler)
		con.GET("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), contactHandler.GetContactSubmissionHandler)
		con.DELETE("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), contactHandler.DeleteContactSubmissionHandler)
	}

	{
		// workshop
		workshopRepository := repository.NewRepositoryWorkshop(db)
		workshopService := service.NewServiceWorkshop(workshopRepository)
		workshopHandler := NewWorkshopHandler(workshopService, statisticsService)
		//--//
		work := router.Group("/api/workshop")
		work.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), workshopHandler.CreateWorkshop)
		work.GET("/", workshopHandler.GetAllWorkshop)
		work.GET("/:slug", workshopHandler.GetOneWorkshop)
		work.PUT("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), workshopHandler.UpdateWorkshop)
		work.DELETE("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), workshopHandler.DeleteWorkshop)
	}

	{
		// ecopedia
		ecopediaRepository := repository.NewRepositoryEcopedia(db)
		ecopediaService := service.NewServiceEcopedia(ecopediaRepository)
		ecopediaHandler := NewEcopediaHandler(ecopediaService, statisticsService)

		eco := router.Group("/api/ecopedia")
		eco.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), ecopediaHandler.PostEcopediaHandler)
		eco.GET("/", ecopediaHandler.GetAllEcopedia)
		eco.GET("/:slug", ecopediaHandler.GetEcopediaByID)
		eco.PUT("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), ecopediaHandler.UpdateEcopedia)
		eco.DELETE("/:ID", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), ecopediaHandler.DeleteEcopedia)
	}

	{
		// artikel
		artikelRepository := repository.NewRepositoryArtikel(db)
		artikelService := service.NewServiceArtikel(artikelRepository)
		artikelHandler := NewArtikelHandler(artikelService, statisticsService)
		//--//
		art := router.Group("/api/article")
		art.POST("/", artikelHandler.CreateArtikel)
		art.GET("/", artikelHandler.GetAllArtikel)
		art.GET("/:slug", artikelHandler.GetOneArtikel)
		art.PUT("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), artikelHandler.UpdateArtikel)
		art.DELETE("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), artikelHandler.DeleteArtikel)
	}

	{
		// event
		eventRepository := repository.NewRepositoryEvent(db)
		eventService := service.NewServiceEvent(eventRepository)
		eventHandler := NewEventHandler(eventService, statisticsService)
		//--//
		eve := router.Group("/api/event")
		eve.POST("/", eventHandler.CreateEvent)
		eve.GET("/", eventHandler.GetAllEvent)
		eve.GET("/:slug", eventHandler.GetOneEvent)
		eve.PUT(":slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), eventHandler.UpdateEvent)
		eve.DELETE("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), eventHandler.DeleteEvent)
	}

	{ //veganGuide
		veganguideRepository := repository.NewRepositoryVeganguide(db)
		veganguideService := service.NewServiceVeganguide(veganguideRepository)
		veganguideHandler := NewVeganguideHandler(veganguideService, statisticsService)
		//--//
		veg := router.Group("/api/veganguide")
		veg.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), veganguideHandler.PostVeganguideHandler)
		veg.GET("/", veganguideHandler.GetAllVeganguide)
		veg.GET("/:slug", veganguideHandler.GetVeganguideByID)
		veg.PUT("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), veganguideHandler.UpdateVeganguide)
		veg.DELETE("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), veganguideHandler.DeleteVeganguide)
	}
	// veganguide

	{ //feedback
		feedbackRepository := repository.NewRepositoryFeedback(db)
		feedbackService := service.NewServiceFeedback(feedbackRepository)
		feedbackHandler := NewFeedbackHandler(feedbackService)
		//--//
		fee := router.Group("/api/feedback")
		fee.POST("/", feedbackHandler.PostFeedbackHandler)
		fee.GET("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), feedbackHandler.GetAllFeedback)
		fee.GET("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), feedbackHandler.GetFeedbackBySlug)
		// fee.PUT("/:id", middleware.AuthMiddleware(authService, userService),  middleware.AuthRole(authService, userService), feedbackHandler.UpdateFeedback)
		fee.DELETE("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), feedbackHandler.DeleteFeedback)
	}
	// feedback

	{ //gallery
		galleryRepository := repository.NewRepositoryGallery(db)
		galleryService := service.NewServiceGallery(galleryRepository)
		galleryHandler := NewGalleryHandler(galleryService, statisticsService)

		gallerys := router.Group("/api/gallery")

		gallerys.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), galleryHandler.CreateGallery)
		gallerys.GET("/", galleryHandler.GetAllGallery)
		gallerys.GET("/:slug", galleryHandler.GetOneGallery)
		gallerys.PUT("/:slug", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), galleryHandler.UpdateGallery)
		gallerys.DELETE("/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), galleryHandler.DeleteGallery)
	}

	{ //products
		productsRepository := repository.NewRepositoryProduct(db)
		productsService := service.NewServiceProduct(productsRepository)
		productHandler := NewProductHandler(productsService, statisticsService)

		products := router.Group("/api/product")
		//get all
		products.Use(middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService))
		products.POST("/", productHandler.CreateProduct)
		products.GET("/:slug", productHandler.ReadProductBySlug)
		products.PUT("/:slug", productHandler.UpdateProductBySlug)
		products.DELETE("/:slug", productHandler.DeleteProductBySlug)

		//shoping
		shoppingCartRepository := repository.NewRepositoryShoppingCart(db)
		shoppingCartService := service.NewServiceShoppingCart(shoppingCartRepository, productsRepository)
		shoppingCartHandler := NewShoppingCartHandler(shoppingCartService)

		shoppingCarts := router.Group("/api/shopping-cart")

		shoppingCarts.Use(middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService))
		shoppingCarts.POST("/", shoppingCartHandler.CreateShoppingCart)
		shoppingCarts.GET("/", shoppingCartHandler.GetShoppingCarts)
		shoppingCarts.GET("/:id", shoppingCartHandler.GetShoppingCartById)
		shoppingCarts.PUT("/:id", shoppingCartHandler.UpdateShoppingCartById) // update qty and total price
		shoppingCarts.DELETE("/:id", shoppingCartHandler.DeleteShoppingCartById)
	}

	{ // hrd
		hrd := router.Group("/api/hrd")
		hrdRepository := repository.NewRepositoryHRD(db)
		hrdService := service.NewServiceHrd(hrdRepository)
		hrdHandler := NewHrdHandler(hrdService)

		hrd.Use(
			middleware.AuthMiddleware(authService, userService),
			middleware.AuthRole(authService, userService),
		)

		hrd.POST("/", hrdHandler.CreateHrd)
		hrd.GET("/", hrdHandler.GetAllHrd)
		hrd.GET("/status/:status", hrdHandler.GetAllByStatus)
		hrd.GET("/departement/:departement", hrdHandler.GetAllByDepartement)
		hrd.GET("/:slug", hrdHandler.GetOneHrd)
		hrd.PUT("/:slug", hrdHandler.UpdateHrd)
		hrd.DELETE("/:slug", hrdHandler.DeleteHrd)
	}

	{ // kajian
		kajian := router.Group("/api/kajian")
		kajianRepository := repository.NewRepositoryKajian(db)
		kajianService := service.NewServiceKajian(kajianRepository)
		kajianHandler := NewKajianHandler(kajianService)

		kajian.Use(
			middleware.AuthMiddleware(authService, userService),
			middleware.AuthRole(authService, userService),
		)
		
		kajian.POST("/", kajianHandler.CreateKajian)
		kajian.GET("/", kajianHandler.GetAllKajian)
		kajian.GET("/:slug", kajianHandler.GetOneKajian)
		kajian.PUT("/:slug", kajianHandler.UpdateKajian)
		kajian.DELETE("/:slug", kajianHandler.DeleteKajian)
	}

	// statistics
	// router.GET("/statistics", statisticsHandler.GetStatisticsHandler)

	// Port
	router.Run(":8080")
}
