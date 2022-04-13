package controllers

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

func Ambassadors(ctx *fiber.Ctx) error {
	var users []models.User

	database.DB.Where("is_ambassador = true").Find(&users)

	return ctx.JSON(users)
}

// func Rankings(ctx *fiber.Ctx) error {
// 	var users []models.User
// 	database.DB.Find(&users, models.User{
// 		IsAmbassador: true,
// 	})
// 	var result []interface{}
// 	for _, user := range users {
// 		ambassador := models.Ambassador(user)
// 		ambassador.CalculateRevenue(database.DB)
// 		result = append(result, fiber.Map{
// 			user.Name(): ambassador.Revenue,
// 		})
// 	}
// 	return ctx.JSON(result)
// }

func Rankings(ctx *fiber.Ctx) error {
	rankings, err := database.Cache.ZRevRangeByScoreWithScores(context.Background(), "rankings", &redis.ZRangeBy{
		Min: "-inf",
		Max: "+inf",
	}).Result()

	if err != nil {
		return err
	}

	result := make(map[string]float64)

	for _, ranking := range rankings {
		result[ranking.Member.(string)] = ranking.Score
	}

	return ctx.JSON(result)
}
