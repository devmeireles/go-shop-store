package controllers

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/devmeireles/go-shop-store/app/models"
	"github.com/devmeireles/go-shop-store/app/services"
	"github.com/devmeireles/go-shop-store/app/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateCategory(c *fiber.Ctx) error {
	data := new(models.Category)

	c.BodyParser(data)

	errors := utils.ValidateStruct(*data)
	if errors != nil {
		res := utils.ResErrorValidation(errors)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	category, err := services.SaveCategory(data)

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(category)
	return c.Status(fiber.StatusCreated).JSON(res)
}

func ListCategories(c *fiber.Ctx) error {
	categories, err := services.GetCategories()

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(categories)
	return c.Status(fiber.StatusOK).JSON(res)
}

func GetCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(res)
	}

	category, err := services.GetCategoryByID(id)

	if err != nil {
		res := utils.ResError(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(res)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(category)
	return c.Status(fiber.StatusOK).JSON(res)
}

func UpdateCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	payload := new(models.Category)
	c.BodyParser(payload)

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(res)
	}

	_, err = services.GetCategoryByID(id)
	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusNotFound).JSON(res)
	}

	payloadFmt, err := json.Marshal(payload)
	category := models.Category{}
	err = json.Unmarshal(payloadFmt, &category)

	err = services.UpdateCategory(category, id)
	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(nil)
	return c.Status(fiber.StatusOK).JSON(res)
}

func DeleteCategory(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(res)
	}

	category, err := services.GetCategoryByID(id)

	if err != nil {
		res := utils.ResError(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(res)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	err = services.DeleteCategory(category, id)
	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(nil)
	return c.Status(fiber.StatusOK).JSON(res)
}
