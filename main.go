package main

import (
	_ "greenwelfare/docs"
	"greenwelfare/handler"
)

func main() {
	// @title Sweager Service API
	// @description Sweager service API in Go using Gin framework
	// @host localhost:8080/
	// @securitydefinitions.apikey BearerAuth
	// @in header
	// @name Authorization
	handler.StartApp()
}
