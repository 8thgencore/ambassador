package internal

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func StripeSecretKey() string {
	path, exists := os.LookupEnv("STRIPE_SECRET_KEY")

	if exists {
		return path
	}
	return ""
}
