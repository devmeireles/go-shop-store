package controllers

import (
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
		return c.Status(fiber.StatusBadRequest).JSON(errors)
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
