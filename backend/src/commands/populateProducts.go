package main

import (
	"ambassador/src/database"
	"ambassador/src/models"
	"math/rand"
	"strconv"

	"github.com/bxcodec/faker/v3"
)

func main() {
	database.Connect()

	for i := 0; i < 30; i++ {

		title:= faker.Word()
		imageUrl := "https://via.placeholder.com/640x480.png/" + strconv.Itoa((rand.Intn(899999) + 100000)) + "?text=" + title

		ambassador := models.Product{
			Title:       title,
			Description: faker.Sentence(),
			Image:       imageUrl,
			Price:       float64(rand.Intn(90) + 10),
		}

		database.DB.Create(&ambassador)
	}
}
