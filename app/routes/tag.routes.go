package routes

import (
	"github.com/devmeireles/go-shop-store/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func TagRoutes(app *fiber.App) {
	tag := app.Group("/tag")
	tag.Get("/", controllers.ListTags)
	tag.Get("/:id", controllers.GetTag)
	tag.Post("/", controllers.CreateTag)
	tag.Put("/:id", controllers.UpdateTag)
	tag.Delete("/:id", controllers.DeleteTag)
}
