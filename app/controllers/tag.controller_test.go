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

func setupTagTestRoutes() *fiber.App {
	os.Setenv("ENVIRONMENT", "test")
	config.ConnectDb()
	app := fiber.New()
	routes.TagRoutes(app)

	return app
}

func TestTags(t *testing.T) {
	handler := utils.FiberToHandlerFunc(setupTagTestRoutes())

	t.Run("Should return an empty list of tags because there's no data", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/tag").
			Expect(t).
			Assert(jsonpath.Len("data", 0)).
			Assert(jsonpath.NotPresent("message")).
			Assert(jsonpath.Equal("success", true)).
			Status(http.StatusOK).
			End()
	})

	t.Run("Create a tag", func(t *testing.T) {
		var tag = models.Tag{
			Status:      1,
			Name:        faker.Global.BeerStyle(),
			Description: faker.Global.Paragraph(3, 5, 12, "\n"),
		}

		tagSave, _ := json.Marshal(tag)

		apitest.New().
			Handler(handler).
			Post("/tag").
			JSON(tagSave).
			Expect(t).
			Assert(jsonpath.Present("data")).
			Assert(jsonpath.NotPresent("message")).
			Assert(jsonpath.Equal("success", true)).
			Status(http.StatusCreated).
			End()
	})

	t.Run("Get all tags", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/tag").
			Expect(t).
			Assert(jsonpath.Present("data")).
			Assert(jsonpath.NotPresent("message")).
			Assert(jsonpath.Equal("success", true)).
			Status(http.StatusOK).
			End()
	})

	t.Run("Shouldn't create a tag due a missing body", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Post("/tag").
			Expect(t).
			Assert(jsonpath.Present("message")).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.Equal("success", false)).
			Status(http.StatusBadRequest).
			End()
	})

	t.Run("Get a tag by ID", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/tag/1").
			Expect(t).
			Assert(jsonpath.Present("data")).
			Assert(jsonpath.NotPresent("message")).
			Assert(jsonpath.Equal("success", true)).
			Status(http.StatusOK).
			End()
	})

	t.Run("Shouldn't get a tag by ID due an unexistent item", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/tag/999999").
			Expect(t).
			Assert(jsonpath.Present("message")).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.Equal("success", false)).
			Status(http.StatusNotFound).
			End()
	})

	t.Run("Shouldn't get a tag by ID due a badly formatted ID", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Get("/tag/abc").
			Expect(t).
			Assert(jsonpath.Present("message")).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.Equal("success", false)).
			Status(http.StatusBadRequest).
			End()
	})

	t.Run("Update a tag", func(t *testing.T) {
		var tag = models.Tag{
			Description: faker.Global.Paragraph(3, 5, 12, "\n"),
		}

		tagUpdate, _ := json.Marshal(tag)

		apitest.New().
			Handler(handler).
			Put("/tag/1").
			JSON(tagUpdate).
			Expect(t).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.NotPresent("message")).
			Assert(jsonpath.Equal("success", true)).
			Status(http.StatusOK).
			End()
	})

	t.Run("Shouldn't update a tag due an unexistent item", func(t *testing.T) {
		var tag = models.Tag{
			Description: faker.Global.Paragraph(3, 5, 12, "\n"),
		}

		tagUpdate, _ := json.Marshal(tag)

		apitest.New().
			Handler(handler).
			Put("/tag/999991").
			JSON(tagUpdate).
			Expect(t).
			Assert(jsonpath.Present("message")).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.Equal("success", false)).
			Status(http.StatusNotFound).
			End()
	})

	t.Run("Delete a tag", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Delete("/tag/1").
			Expect(t).
			Assert(jsonpath.NotPresent("data")).
			Assert(jsonpath.NotPresent("message")).
			Assert(jsonpath.Equal("success", true)).
			Status(http.StatusOK).
			End()
	})

	t.Run("Shouldn't delete a tag due an unexistent item", func(t *testing.T) {
		apitest.New().
			Handler(handler).
			Delete("/tag/999991").
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
