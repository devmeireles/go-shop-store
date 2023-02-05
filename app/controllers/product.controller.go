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

func CreateProduct(c *fiber.Ctx) error {
	data := new(models.Product)

	c.BodyParser(data)

	errors := utils.ValidateStruct(*data)
	if errors != nil {
		res := utils.ResErrorValidation(errors)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	product, err := services.SaveProduct(data)

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(product)
	return c.Status(fiber.StatusCreated).JSON(res)
}

func ListProducts(c *fiber.Ctx) error {
	products, err := services.GetProducts()

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(products)
	return c.Status(fiber.StatusOK).JSON(res)
}

func GetProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(res)
	}

	product, err := services.GetProductByID(id)

	if err != nil {
		res := utils.ResError(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(res)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(product)
	return c.Status(fiber.StatusOK).JSON(res)
}

func UpdateProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	payload := new(models.Product)
	c.BodyParser(payload)

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(res)
	}

	_, err = services.GetProductByID(id)
	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusNotFound).JSON(res)
	}

	payloadFmt, err := json.Marshal(payload)
	product := models.Product{}
	err = json.Unmarshal(payloadFmt, &product)

	err = services.UpdateProduct(product, id)
	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(nil)
	return c.Status(fiber.StatusOK).JSON(res)
}

func DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(res)
	}

	product, err := services.GetProductByID(id)

	if err != nil {
		res := utils.ResError(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(res)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	err = services.DeleteProduct(product, id)
	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(nil)
	return c.Status(fiber.StatusOK).JSON(res)
}
