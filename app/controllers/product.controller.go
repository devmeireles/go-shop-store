package controllers

import (
	"github.com/devmeireles/go-shop-store/app/services"
	"github.com/devmeireles/go-shop-store/app/utils"
	"github.com/gofiber/fiber/v2"
)

func ListProducts(c *fiber.Ctx) error {
	products, err := services.GetProducts()

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(products)
	return c.Status(fiber.StatusOK).JSON(res)
}
