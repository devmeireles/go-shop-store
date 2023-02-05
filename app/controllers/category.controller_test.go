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

func setupCategoryTestRoutes() *fiber.App {
	os.Setenv("ENVIRONMENT", "test")
	config.ConnectDb()
	app := fiber.New()
	routes.CategoryRoutes(app)

	return app
}

func TestListCategories(t *testing.T) {
	handler := utils.FiberToHandlerFunc(setupCategoryTestRoutes())

	t.Run("Should return an empty list of categories because there's no data", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/category").
			Expect(t).
			Assert(jsonpath.Len("data", 0)).
			Assert(jsonpath.NotPresent("message")).
			Assert(jsonpath.Equal("success", true)).
			Status(http.StatusOK).
			End()
	})

	t.Run("Create a category", func(t *testing.T) {
		var category = models.Category{
			Status:      1,
			Name:        faker.Global.BeerStyle(),
			Description: faker.Global.Paragraph(3, 5, 12, "\n"),
		}

		categorySave, _ := json.Marshal(category)

		apitest.New().
			Handler(handler).
			Post("/category").
			JSON(categorySave).
			Expect(t).
			Assert(jsonpath.Present("data")).
			Assert(jsonpath.NotPresent("message")).
			Assert(jsonpath.Equal("success", true)).
			Status(http.StatusCreated).
			End()
	})

	t.Run("Get all categories", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/category").
			Expect(t).
			Assert(jsonpath.Present("data")).
			Assert(jsonpath.NotPresent("message")).
			Assert(jsonpath.Equal("success", true)).
			Status(http.StatusOK).
			End()
	})

	t.Run("Shouldn't create a category due a missing body", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Post("/category").
			Expect(t).
			Assert(jsonpath.Present("message")).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.Equal("success", false)).
			Status(http.StatusBadRequest).
			End()
	})

	t.Run("Get a category by ID", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/category/1").
			Expect(t).
			Assert(jsonpath.Present("data")).
			Assert(jsonpath.NotPresent("message")).
			Assert(jsonpath.Equal("success", true)).
			Status(http.StatusOK).
			End()
	})

	t.Run("Shouldn't get a category by ID due an unexistent item", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/category/999999").
			Expect(t).
			Assert(jsonpath.Present("message")).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.Equal("success", false)).
			Status(http.StatusNotFound).
			End()
	})

	t.Run("Shouldn't get a category by ID due a badly formatted ID", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/category/abc").
			Expect(t).
			Assert(jsonpath.Present("message")).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.Equal("success", false)).
			Status(http.StatusBadRequest).
			End()
	})

	t.Run("Update a category", func(t *testing.T) {
		var category = models.Category{
			Description: faker.Global.Paragraph(3, 5, 12, "\n"),
		}

		categoryUpdate, _ := json.Marshal(category)

		apitest.New().
			Handler(handler).
			Put("/category/1").
			JSON(categoryUpdate).
			Expect(t).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.NotPresent("message")).
			Assert(jsonpath.Equal("success", true)).
			Status(http.StatusOK).
			End()
	})

	t.Run("Shouldn't update a category due an unexistent item", func(t *testing.T) {
		var category = models.Category{
			Description: faker.Global.Paragraph(3, 5, 12, "\n"),
		}

		categoryUpdate, _ := json.Marshal(category)

		apitest.New().
			Handler(handler).
			Put("/category/999991").
			JSON(categoryUpdate).
			Expect(t).
			Assert(jsonpath.Present("message")).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.Equal("success", false)).
			Status(http.StatusNotFound).
			End()
	})

	t.Run("Delete a category", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Delete("/category/1").
			Expect(t).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.NotPresent("message")).
			Assert(jsonpath.Equal("success", true)).
			Status(http.StatusOK).
			End()
	})

	t.Run("Shouldn't delete a category due an unexistent item", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Delete("/category/999991").
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
