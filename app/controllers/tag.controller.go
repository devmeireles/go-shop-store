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

func CreateTag(c *fiber.Ctx) error {
	data := new(models.Tag)

	c.BodyParser(data)

	errors := utils.ValidateStruct(*data)
	if errors != nil {
		res := utils.ResErrorValidation(errors)
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}

	tag, err := services.SaveTag(data)

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(tag)
	return c.Status(fiber.StatusCreated).JSON(res)
}

func ListTags(c *fiber.Ctx) error {
	tags, err := services.GetTags()

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(tags)
	return c.Status(fiber.StatusOK).JSON(res)
}

func GetTag(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(res)
	}

	tag, err := services.GetTagByID(id)

	if err != nil {
		res := utils.ResError(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(res)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(tag)
	return c.Status(fiber.StatusOK).JSON(res)
}

func UpdateTag(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	payload := new(models.Tag)
	c.BodyParser(payload)

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(res)
	}

	_, err = services.GetTagByID(id)
	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusNotFound).JSON(res)
	}

	payloadFmt, err := json.Marshal(payload)
	tag := models.Tag{}
	err = json.Unmarshal(payloadFmt, &tag)

	err = services.UpdateTag(tag, id)
	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(nil)
	return c.Status(fiber.StatusOK).JSON(res)
}

func DeleteTag(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.ErrBadRequest.Code).JSON(res)
	}

	tag, err := services.GetTagByID(id)

	if err != nil {
		res := utils.ResError(err)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(res)
		}

		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	err = services.DeleteTag(tag, id)
	if err != nil {
		res := utils.ResError(err)
		return c.Status(fiber.StatusInternalServerError).JSON(res)
	}

	res := utils.ResSuccess(nil)
	return c.Status(fiber.StatusOK).JSON(res)
}
