package main

import (
	_ "greenwelfare/docs"
	"greenwelfare/handler"
)

// @title Sweager Service API
// @description Sweager service API in Go using Gin framework
// @host backend-gwf-production.up.railway.app
// @securitydefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	handler.StartApp()
}

