package routes

import (
	"github.com/devmeireles/go-shop-store/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App) {
	category := app.Group("/category")
	category.Get("/", controllers.ListCategories)
	category.Get("/:id", controllers.GetCategory)
	category.Post("/", controllers.CreateCategory)
	category.Put("/:id", controllers.UpdateCategory)
	category.Delete("/:id", controllers.DeleteCategory)
}
