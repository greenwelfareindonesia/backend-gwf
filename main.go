package main

import (
	_ "greenwelfare/docs"
	"greenwelfare/handler"
)

func main() {
	// @title Sweager Service API
	// @description Sweager service API in Go using Gin framework
	// @host https://backend-gwf-production.up.railway.app/
	// @securitydefinitions.apikey BearerAuth
	// @in header
	// @name Authorization
	handler.StartApp()
}
