package controllers_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/devmeireles/go-shop-store/app/config"
	"github.com/devmeireles/go-shop-store/app/routes"
	"github.com/devmeireles/go-shop-store/app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/steinfletcher/apitest"
)

func setupTestRoutes() *fiber.App {
	os.Setenv("ENVIRONMENT", "test")
	config.ConnectDb()
	app := fiber.New()
	routes.ProductRoutes(app)

	return app
}

func TestListProducts(t *testing.T) {
	handler := utils.FiberToHandlerFunc(setupTestRoutes())

	t.Run("Get all products", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/product").
			Expect(t).
			Status(http.StatusOK).
			End()
	})
}
