package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"context"
	"encoding/json"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	products_frontend = "products_frontend"
	products_backend  = "products_backend"
)

func Products(ctx *fiber.Ctx) error {
	var products []models.Product

	database.DB.Find(&products)

	return ctx.JSON(products)
}

func CreateProducts(ctx *fiber.Ctx) error {
	var product models.Product

	if err := ctx.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Create(&product)

	go database.ClearCache(products_frontend, products_backend)

	return ctx.JSON(product)
}

func GetProduct(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	var product models.Product

	product.Id = uint(id)

	database.DB.Find(&product)

	return ctx.JSON(product)
}

func UpdateProduct(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	product := models.Product{}
	product.Id = uint(id)

	if err := ctx.BodyParser(&product); err != nil {
		return err
	}

	database.DB.Model(&product).Updates(&product)

	go database.ClearCache(products_frontend, products_backend)

	return ctx.JSON(product)
}

func DeleteProduct(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	product := models.Product{}
	product.Id = uint(id)

	database.DB.Delete(&product)

	go database.ClearCache(products_frontend, products_backend)

	return nil
}

func ProductsFrontend(ctx *fiber.Ctx) error {
	var products []models.Product
	var c = context.Background()

	result, err := database.Cache.Get(c, products_frontend).Result()

	if err != nil {
		database.DB.Find(&products)

		bytes, err := json.Marshal(products)
		if err != nil {
			panic(err)
		}

		if errKey := database.Cache.Set(c, products_frontend, bytes, 30*time.Minute).Err(); errKey != nil {
			panic(errKey)
		}
	} else {
		json.Unmarshal([]byte(result), &products)
	}

	return ctx.JSON(products)
}

func ProductsBackend(ctx *fiber.Ctx) error {
	var products []models.Product
	var c = context.Background()

	result, err := database.Cache.Get(c, products_backend).Result()

	if err != nil {
		database.DB.Find(&products)

		bytes, err := json.Marshal(products)
		if err != nil {
			panic(err)
		}

		if errKey := database.Cache.Set(c, products_backend, bytes, 30*time.Minute).Err(); errKey != nil {
			panic(errKey)
		}
	} else {
		json.Unmarshal([]byte(result), &products)
	}

	// search
	var searchedProducts []models.Product

	if s := ctx.Query("s"); s != "" {
		lower := strings.ToLower(s)
		for _, product := range products {
			if strings.Contains(strings.ToLower(product.Title), lower) ||
				strings.Contains(strings.ToLower(product.Description), lower) {
				searchedProducts = append(searchedProducts, product)
			}
		}
	} else {
		searchedProducts = products
	}

	// sorted
	if sortParam := ctx.Query("sort"); sortParam != "" {
		sortLower := strings.ToLower(sortParam)
		if sortLower == "asc" {
			sort.Slice(searchedProducts, func(i, j int) bool {
				return searchedProducts[i].Price < searchedProducts[j].Price
			})
		} else if sortLower == "desc" {
			sort.Slice(searchedProducts, func(i, j int) bool {
				return searchedProducts[i].Price > searchedProducts[j].Price
			})
		}
	}

	// paginating
	var total = len(searchedProducts)
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	perPage := 9

	var data []models.Product = searchedProducts

	if total <= page*perPage && total >= (page-1)*perPage {
		data = searchedProducts[(page-1)*perPage : total]
	} else if total >= page*perPage {
		data = searchedProducts[(page-1)*perPage : page*perPage]
	} else {
		data = []models.Product{}
	}

	return ctx.JSON(fiber.Map{
		"data":      data,
		"total":     total,
		"page":      page,
		"last_page": total/perPage + 1,
	})
}
