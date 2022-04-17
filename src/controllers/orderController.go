package controllers

import (
	"ambassador/src/database"
	"ambassador/src/internal"
	"ambassador/src/models"
	"context"
	"encoding/json"
	"fmt"
	"net/smtp"

	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/checkout/session"
)

func Orders(ctx *fiber.Ctx) error {
	var orders []models.Order

	database.DB.Preload("OrderItems").Find(&orders)

	for i, order := range orders {
		orders[i].Name = order.FullName()
		orders[i].Total = order.GetTotal()
	}

	return ctx.JSON(orders)
}

type CreateOrderRequest struct {
	Code      string
	FirstName string
	LastName  string
	Email     string
	Address   string
	Country   string
	City      string
	Zip       string
	Products  []map[string]int
}

func CreateOrder(ctx *fiber.Ctx) error {
	var request CreateOrderRequest

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	link := models.Link{
		Code: request.Code,
	}

	database.DB.Preload("User").First(&link)

	if link.Id == 0 {
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"message": "Invalid link!",
		})
	}

	order := models.Order{
		Code:            link.Code,
		UserId:          link.UserId,
		AmbassadorEmail: link.User.Email,
		FirstName:       link.User.FirstName,
		LastName:        link.User.LastName,
		Email:           request.Email,
		Address:         request.Address,
		Country:         request.Country,
		City:            request.City,
		Zip:             request.Zip,
	}

	tx := database.DB.Begin()

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var lineItems []*stripe.CheckoutSessionLineItemParams

	for _, requestProduct := range request.Products {
		product := models.Product{}
		product.Id = uint(requestProduct["product_id"])
		database.DB.First(&product)

		total := product.Price * float64(requestProduct["quantity"])

		item := models.OrderItem{
			OrderId:           order.Id,
			ProductTitle:      product.Title,
			Price:             product.Price,
			Quantity:          uint(requestProduct["quantity"]),
			AmbassadorRevenue: 0.1 * total,
			AdminRevenue:      0.9 * total,
		}

		if err := tx.Create(&item).Error; err != nil {
			tx.Rollback()
			ctx.Status(fiber.StatusBadRequest)
			return ctx.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		lineItems = append(lineItems, &stripe.CheckoutSessionLineItemParams{
			Name:        stripe.String(product.Title),
			Description: stripe.String(product.Description),
			Images:      []*string{stripe.String(product.Image)},
			Amount:      stripe.Int64(100 * int64(product.Price)),
			Currency:    stripe.String("usd"),
			Quantity:    stripe.Int64(int64(requestProduct["quantity"])),
		})
	}

	// stripe checkout
	stripe.Key = internal.StripeSecretKey()

	params := stripe.CheckoutSessionParams{
		SuccessURL:         stripe.String("http://localhost:5000/success?source={CHECKOUT_SESSION_ID}"),
		CancelURL:          stripe.String("http://localhost:5000/error"),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems:          lineItems,
	}

	source, err := session.New(&params)
	j, err := json.Marshal(&source)
	raw := json.RawMessage(string(j))
	if err != nil {
		panic(err)
	}

	if err != nil {
		tx.Rollback()
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	order.TransactionId = source.ID

	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	tx.Commit()

	return ctx.JSON(raw)
}

func CompleteOrder(ctx *fiber.Ctx) error {
	var data map[string]string

	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	order := models.Order{}

	database.DB.Preload("OrderItems").First(&order, models.Order{
		TransactionId: data["source"],
	})

	if order.Id == 0 {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"message": "Order not found",
		})
	}

	order.Complete = true
	database.DB.Save(&order)

	go func(order models.Order) {
		ambassadorRevenue := 0.0
		adminRevenue := 0.0

		for _, item := range order.OrderItems {
			ambassadorRevenue += item.AmbassadorRevenue
			adminRevenue += item.AdminRevenue
		}

		user := models.User{}
		user.Id = order.UserId

		database.DB.First(&user)
		// https://redis.io/commands/zincrby
		database.Cache.ZIncrBy(
			context.Background(),
			"rankings",
			ambassadorRevenue,
			user.Name(),
		)

		// Email to ambassador
		ambassadorMessage := []byte(
			fmt.Sprintf("You earned $%f Link: #%s",
				ambassadorRevenue,
				order.Code))
		// Mailhog
		smtp.SendMail(
			"host.docker.internal:1025",     // Adress
			nil,                             // Authentication
			"no-reply@email.com",            // from
			[]string{order.AmbassadorEmail}, // To
			ambassadorMessage)               // Message

		// Email to admin
		adminMessage := []byte(
			fmt.Sprintf("Order: #%d, Total: %f",
				order.Id,
				adminRevenue))
		// Mailhog
		smtp.SendMail(
			"host.docker.internal:1025", // Address
			nil,                         // Authentication
			"no-reply@email.com",        // from
			[]string{"admin@admin.com"}, // To
			adminMessage)                // Message
	}(order)

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}
