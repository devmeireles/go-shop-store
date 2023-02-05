package main

import (
	database "github.com/devmeireles/go-shop-store/app/config"
	"github.com/devmeireles/go-shop-store/app/routes"
	"github.com/gofiber/fiber/v2"
)

func setupDatabase() {
	database.ConnectDb()
}

func setupRoutes() {
	app := fiber.New()
	routes.ProductRoutes(app)
	routes.CategoryRoutes(app)
	routes.TagRoutes(app)
	app.Listen(":3000")
}
