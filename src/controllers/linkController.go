package controllers

import (
	"ambassador/src/database"
	"ambassador/src/middleware"
	"ambassador/src/models"
	"strconv"

	"github.com/bxcodec/faker/v3"
	"github.com/gofiber/fiber/v2"
)

func Link(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	var links []models.Link

	database.DB.Where("user_id = ?", id).Find(&links)

	for i, link := range links {
		var orders []models.Order

		database.DB.Where("code = ? and complete = true", link.Code).Find(&orders)

		links[i].Orders = orders
	}

	return ctx.JSON(links)
}

type CreateLinkRequest struct {
	Products []int
}

func CreateLink(ctx *fiber.Ctx) error {
	var request CreateLinkRequest

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	id, _ := middleware.GetUserId(ctx)

	link := models.Link{
		UserId: id,
		Code:   faker.Username(),
	}

	for _, productID := range request.Products {
		product := models.Product{}
		product.Id = uint(productID)
		link.Products = append(link.Products, product)
	}

	database.DB.Create(&link)

	return ctx.JSON(link)
}

func Stats(ctx *fiber.Ctx) error {
	id, _ := middleware.GetUserId(ctx)

	var links []models.Link

	database.DB.Find(&links, models.Link{
		UserId: id,
	})

	var result []interface{}

	var orders []models.Order

	for _, link := range links {
		database.DB.Preload("OrderItems").Find(&orders, &models.Order{
			Code:     link.Code,
			Complete: true,
		})

		revenue := 0.0

		for _, order := range orders {
			revenue += order.GetTotal()
		}

		result = append(result, fiber.Map{
			"code":    link.Code,
			"count":   len(orders),
			"revenue": revenue,
		})
	}

	return ctx.JSON(result)
}
