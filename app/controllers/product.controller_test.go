package controllers_test

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/bgadrian/fastfaker/faker"
	"github.com/devmeireles/go-shop-store/app/config"
	"github.com/devmeireles/go-shop-store/app/models"
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

	t.Run("Create a product", func(t *testing.T) {
		var product = models.Product{
			Status:      1,
			Title:       faker.Global.BeerName(),
			Description: faker.Global.Paragraph(3, 5, 12, "\n"),
			Price:       faker.Global.Price(3, 25),
		}

		productSave, _ := json.Marshal(product)

		apitest.New().
			Handler(handler).
			Post("/product").
			JSON(productSave).
			Expect(t).
			Status(http.StatusCreated).
			End()
	})

	t.Run("Shouldn't create a product due a missing body", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Post("/product").
			Expect(t).
			Status(http.StatusBadRequest).
			End()
	})

	t.Run("Get a product by ID", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/product/1").
			Expect(t).
			Status(http.StatusOK).
			End()
	})

	t.Run("Shouldn't get a product by ID due an unexistent item", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/product/999999").
			Expect(t).
			Status(http.StatusNotFound).
			End()
	})

	t.Run("Shouldn't get a product by ID due a badly formatted ID", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/product/abc").
			Expect(t).
			Status(http.StatusBadRequest).
			End()
	})
}
