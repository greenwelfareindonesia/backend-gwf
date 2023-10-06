package main

import (
	"fmt"
	"greenwelfare/artikel"
	"greenwelfare/auth"
	"greenwelfare/contact"
	"greenwelfare/ecopedia"
	"greenwelfare/endpointCount"
	"greenwelfare/event"
	"greenwelfare/feedback"
	"greenwelfare/handler"
	"greenwelfare/helper"
	"greenwelfare/user"
	"greenwelfare/veganguide"
	"greenwelfare/workshop"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Database connection
	dsn := "root:@tcp(127.0.0.1:3306)/greenwelfare?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db Connestion Error")
	}

	// Auto Migration
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&artikel.Artikel{})
	db.AutoMigrate(&ecopedia.Ecopedia{})
	db.AutoMigrate(&contact.Contact{})
	db.AutoMigrate(&workshop.Workshop{})
	db.AutoMigrate(&event.Event{})
	db.AutoMigrate(&workshop.Workshop{})
	db.AutoMigrate(&veganguide.Veganguide{})
	db.AutoMigrate(&feedback.Feedback{})
	db.AutoMigrate(&endpointcount.Statistics{})


	// fmt.Println("Database Connection Success") //

	router := gin.Default()
	// api := router.Group("/api") // penggunaan contoh: http://localhost:8080/api/user/login

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
	user.DELETE("/", authMiddleware(authService, userService), userHandler.DeletedUser)
	user.PUT("/:id", userHandler.UpdateUser)

	// contact
	contactRepository := contact.NewRepository(db)
	contactService := contact.NewService(contactRepository)
	contactHandler := handler.NewContactHandler(contactService)
	//--//
	con := router.Group("/contact")
	con.POST("/", contactHandler.SubmitContactForm)
	con.GET("/", contactHandler.GetContactSubmissionsHandler)
	con.GET("/:id", contactHandler.GetContactSubmissionHandler)
	con.DELETE("/:id", contactHandler.DeleteContactSubmissionHandler)

	// workshop
	workshopRepository := workshop.NewRepository(db)
	workshopService := workshop.NewService(workshopRepository)
	workshopHandler := handler.NewWorkshopHandler(workshopService, statisticsService)
	//--//
	work := router.Group("/workshop")
	work.POST("/", workshopHandler.CreateWorkshop)
	work.GET("/", workshopHandler.GetAllWorkshop)
	work.GET("/:id", workshopHandler.GetOneWorkshop)
	work.PUT("/:id", workshopHandler.UpdateWorkshop)
	work.DELETE("/:id", workshopHandler.DeleteWorkshop)

	// ecopedia
	ecopediaRepository := ecopedia.NewRepository(db)
	ecopediaService := ecopedia.NewService(ecopediaRepository)
	ecopediaHandler := handler.NewEcopediaHandler(ecopediaService, statisticsService)
	//--/
	eco := router.Group("/ecopedia")
	eco.POST("/", ecopediaHandler.PostEcopediaHandler)
	eco.GET("/", ecopediaHandler.GetAllEcopedia)
	eco.GET("/:id", ecopediaHandler.GetEcopediaByID)
	eco.PUT("/:id", ecopediaHandler.UpdateEcopedia)
	eco.DELETE("/:id", ecopediaHandler.DeleteEcopedia)

	// artikel
	artikelRepository := artikel.NewRepository(db)
	artikelService := artikel.NewService(artikelRepository)
	artikelHandler := handler.NewArtikelHandler(artikelService, statisticsService)
	//--//
	art := router.Group("/artikel")
	art.POST("/", artikelHandler.CreateArtikel)
	art.GET("/", artikelHandler.GetAllArtikel)
	art.GET("/:id", artikelHandler.GetOneArtikel)
	art.PUT("/:id", artikelHandler.UpdateArtikel)
	art.DELETE("/:id", artikelHandler.DeleteArtikel)

	// event
	eventRepository := event.NewRepository(db)
	eventService := event.NewService(eventRepository)
	eventHandler := handler.NewEventHandler(eventService, statisticsService)
	//--//
	eve := router.Group("/event")
	eve.POST("/", eventHandler.CreateEvent)
	eve.GET("/", eventHandler.GetAllEvent)
	eve.GET("/:id", eventHandler.GetOneEvent)
	eve.PUT(":id", eventHandler.UpdateEvent)
	eve.DELETE("/:id", eventHandler.DeleteEvent)

	// veganguide
	veganguideRepository := veganguide.NewRepository(db)
	veganguideService := veganguide.NewService(veganguideRepository)
	veganguideHandler := handler.NewVeganguideHandler(veganguideService, statisticsService)
	//--//
	veg := router.Group("/veganguide")
	veg.POST("/", veganguideHandler.PostVeganguideHandler)
	veg.GET("/", veganguideHandler.GetAllVeganguide)
	veg.GET("/:id", veganguideHandler.GetVeganguideByID)
	veg.PUT("/:id", veganguideHandler.UpdateVeganguide)
	veg.DELETE("/:id", veganguideHandler.DeleteVeganguide)

	// feedback
	feedbackRepository := feedback.NewRepository(db)
	feedbackService := feedback.NewService(feedbackRepository)
	feedbackHandler := handler.NewFeedbackHandler(feedbackService)
	//--//
	fee := router.Group("/feedback")
	fee.POST("/", feedbackHandler.PostFeedbackHandler)
	fee.GET("/", feedbackHandler.GetAllFeedback)
	fee.GET("/:id", feedbackHandler.GetFeedbackByID)
	fee.PUT("/:id", feedbackHandler.UpdateFeedback)
	fee.DELETE("/:id", feedbackHandler.DeleteFeedback)
	router.GET("/statistics", statisticsHandler.GetStatisticsHandler)

	// Port
	router.Run(":8080")
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		fmt.Println(authHeader)
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrToken := strings.Split(authHeader, " ")
		if len(arrToken) == 2 {
			//nah ini kalau emang ada dua key nya dan sesuai, maka tokenString tadi masuk ke arrtoken index ke1
			tokenString = arrToken[1]
		}
		token, err := authService.ValidasiToken(tokenString)
		fmt.Println(token, err)
		if err != nil {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		fmt.Println(claim, ok)
		if !ok || !token.Valid {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByid(userID)
		fmt.Println(user, err)
		if err != nil {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}
