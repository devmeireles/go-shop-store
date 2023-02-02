package routes

import (
	"github.com/devmeireles/go-shop-store/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
	product := app.Group("/product")
	product.Get("/", controllers.ListProducts)
	product.Get("/:id", controllers.GetProduct)
	product.Post("/", controllers.CreateProduct)

}
