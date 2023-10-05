package main

import (
	"fmt"
	"greenwelfare/artikel"
	"greenwelfare/auth"
	"greenwelfare/contact"
	"greenwelfare/ecopedia"
	"greenwelfare/event"
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
	db.AutoMigrate(&ecopedia.Ecopedia{})
	db.AutoMigrate(&contact.Contact{})
	db.AutoMigrate(&workshop.Workshop{})
	db.AutoMigrate(&event.Event{})
	db.AutoMigrate(&artikel.Artikel{})
	db.AutoMigrate(&workshop.Workshop{})
	db.AutoMigrate(&veganguide.Veganguide{})

	// fmt.Println("Database Connection Success") //

	router := gin.Default()

	// user
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)
	//--//
	api := router.Group("/users")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailabilty)
	api.DELETE("/", authMiddleware(authService, userService), userHandler.DeletedUser)
	api.PUT("/:id", userHandler.UpdateUser)

	// contact
	contactRepository := contact.NewRepository(db)
	contactService := contact.NewService(contactRepository)
	contactHandler := handler.NewContactHandler(contactService)
	//--//
	router.POST("/contact", contactHandler.SubmitContactForm)
	router.GET("/contact", contactHandler.GetContactSubmissionsHandler)
	router.GET("/contact/:id", contactHandler.GetContactSubmissionHandler)
	router.DELETE("/contact/:id", contactHandler.DeleteContactSubmissionHandler)

	// workshop
	workshopRepository := workshop.NewRepository(db)
	workshopService := workshop.NewService(workshopRepository)
	//--//
	workshopHandler := handler.NewWorkshopHandler(workshopService)
	router.POST("/workshop", workshopHandler.CreateWorkshop)
	router.GET("/workshop", workshopHandler.GetAllWorkshop)
	router.GET("/workshop/:id", workshopHandler.GetOneWorkshop)
	router.PUT("/workshop/:id", workshopHandler.UpdateWorkshop)
	router.DELETE("/workshop/:id", workshopHandler.DeleteWorkshop)

	// ecopedia
	ecopediaRepository := ecopedia.NewRepository(db)
	ecopediaService := ecopedia.NewService(ecopediaRepository)
	ecopediaHandler := handler.NewEcopediaHandler(ecopediaService)
	//--//
	eco := router.Group("/eco")
	eco.GET("/ecopedias", ecopediaHandler.GetAllEcopedia)
	eco.POST("/create", ecopediaHandler.PostEcopediaHandler)
	eco.GET("/ecopedias/:id", ecopediaHandler.GetEcopediaByID)
	eco.DELETE("/delete/:id", ecopediaHandler.DeleteEcopedia)
	eco.PUT("/update/:id", ecopediaHandler.UpdateEcopedia)

	// artikel
	artikelRepository := artikel.NewRepository(db)
	artikelService := artikel.NewService(artikelRepository)
	artikelHandler := handler.NewArtikelHandler(artikelService)
	//--//
	apiArtikel := router.Group("/artikel")
	apiArtikel.POST("/", artikelHandler.CreateArtikel)
	apiArtikel.GET("/", artikelHandler.GetAllArtikel)
	apiArtikel.DELETE("/delete/:id", artikelHandler.DeleteArtikel)
	apiArtikel.GET("/:id", artikelHandler.GetOneArtikel)
	apiArtikel.PUT("/update/:id", artikelHandler.UpdateArtikel)

	// event
	eventRepository := event.NewRepository(db)
	eventService := event.NewService(eventRepository)
	eventHandler := handler.NewEventHandler(eventService)
	//--//
	apiEvent := router.Group("/event")
	apiEvent.POST("/", eventHandler.CreateEvent)
	apiEvent.GET("/", eventHandler.GetAllEvent)
	apiEvent.DELETE("/delete/:id", eventHandler.DeleteEvent)
	apiEvent.GET("/:id", eventHandler.GetOneEvent)
	apiEvent.PUT("update/:id", eventHandler.UpdateEvent)

	// veganguide
	veganguideRepository := veganguide.NewRepository(db)
	veganguideService := veganguide.NewService(veganguideRepository)
	veganguideHandler := handler.NewVeganguideHandler(veganguideService)
	//--//
	router.POST("/veganguide", veganguideHandler.PostVeganguideHandler)
	router.GET("/veganguide", veganguideHandler.GetAllVeganguide)
	router.GET("/veganguide/:id", veganguideHandler.GetVeganguideByID)
	router.PUT("/veganguide/:id", veganguideHandler.UpdateVeganguide)
	router.DELETE("/veganguide/:id", veganguideHandler.DeleteVeganguide)

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
