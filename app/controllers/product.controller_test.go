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
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func setupProductTestRoutes() *fiber.App {
	// os.Remove("../../db_test.db")
	os.Setenv("ENVIRONMENT", "test")
	config.ConnectDb()
	app := fiber.New()
	routes.ProductRoutes(app)

	return app
}

func TestCategories(t *testing.T) {
	handler := utils.FiberToHandlerFunc(setupProductTestRoutes())

	t.Run("Should return an empty list of products because there's no data", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/product").
			Expect(t).
			Assert(jsonpath.Len("data", 0)).
			Assert(jsonpath.NotPresent("message")).
			Assert(jsonpath.Equal("success", true)).
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
			Assert(jsonpath.Present("data")).
			Assert(jsonpath.NotPresent("message")).
			Assert(jsonpath.Equal("success", true)).
			Status(http.StatusCreated).
			End()
	})

	t.Run("Get all products", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/product").
			Expect(t).
			Assert(jsonpath.Present("data")).
			Assert(jsonpath.NotPresent("message")).
			Assert(jsonpath.Equal("success", true)).
			Status(http.StatusOK).
			End()
	})

	t.Run("Shouldn't create a product due a missing body", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Post("/product").
			Expect(t).
			Assert(jsonpath.Present("message")).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.Equal("success", false)).
			Status(http.StatusBadRequest).
			End()
	})

	t.Run("Get a product by ID", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/product/1").
			Expect(t).
			Assert(jsonpath.Present("data")).
			Assert(jsonpath.NotPresent("message")).
			Assert(jsonpath.Equal("success", true)).
			Status(http.StatusOK).
			End()
	})

	t.Run("Shouldn't get a product by ID due an unexistent item", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/product/999999").
			Expect(t).
			Assert(jsonpath.Present("message")).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.Equal("success", false)).
			Status(http.StatusNotFound).
			End()
	})

	t.Run("Shouldn't get a product by ID due a badly formatted ID", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/product/abc").
			Expect(t).
			Assert(jsonpath.Present("message")).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.Equal("success", false)).
			Status(http.StatusBadRequest).
			End()
	})

	t.Run("Update a product", func(t *testing.T) {
		var product = models.Product{
			Price: faker.Global.Price(3, 25),
		}

		productUpdate, _ := json.Marshal(product)

		apitest.New().
			Handler(handler).
			Put("/product/1").
			JSON(productUpdate).
			Expect(t).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.NotPresent("message")).
			Assert(jsonpath.Equal("success", true)).
			Status(http.StatusOK).
			End()
	})

	t.Run("Shouldn't update a product due an unexistent item", func(t *testing.T) {
		var product = models.Product{
			Price: faker.Global.Price(3, 25),
		}

		productUpdate, _ := json.Marshal(product)

		apitest.New().
			Handler(handler).
			Put("/product/999991").
			JSON(productUpdate).
			Expect(t).
			Assert(jsonpath.Present("message")).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.Equal("success", false)).
			Status(http.StatusNotFound).
			End()
	})

	t.Run("Delete a product", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Delete("/product/1").
			Expect(t).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.NotPresent("message")).
			Assert(jsonpath.Equal("success", true)).
			Status(http.StatusOK).
			End()
	})

	t.Run("Shouldn't delete a product due an unexistent item", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Delete("/product/999991").
			Expect(t).
			Assert(jsonpath.Present("message")).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.Equal("success", false)).
			Status(http.StatusNotFound).
			End()
	})

	t.Cleanup(func() {
		os.Remove("../../db_test.db")
	})
}
