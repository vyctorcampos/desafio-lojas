package main

import (
	"log"

	"desafio-lojas/database"
	"desafio-lojas/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	database.Connect()

	routes.RegisterRoutes(e)

	log.Println("Servidor rodando em http://localhost:8080")
	e.Logger.Fatal(e.Start(":8080"))
}
