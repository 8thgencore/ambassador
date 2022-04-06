package routes

import (
	"ambassador/src/controllers"
	"ambassador/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	// Group
	api := app.Group("api")

	// admin
	admin := api.Group("admin")
	admin.Post("register", controllers.Register)
	admin.Post("login", controllers.Login)

	adminAuthenticated := admin.Use(middleware.IsAuthenticated)
	adminAuthenticated.Get("user", controllers.User)
	adminAuthenticated.Get("logout", controllers.Logout)
	adminAuthenticated.Put("users/info", controllers.UpdateInfo)
	adminAuthenticated.Put("users/password", controllers.UpdatePassword)

	adminAuthenticated.Get("ambassadors", controllers.Ambassadors)

	adminAuthenticated.Get("products", controllers.Products)
	adminAuthenticated.Post("products", controllers.CreateProducts)
	adminAuthenticated.Get("products/:id", controllers.GetProduct)
	adminAuthenticated.Put("products/:id", controllers.UpdateProduct)
	adminAuthenticated.Delete("products/:id", controllers.DeleteProduct)

	adminAuthenticated.Get("users/:id/links", controllers.Link)
	adminAuthenticated.Get("orders", controllers.Orders)

	// ambassador
	ambassador := api.Group("ambassador")
	ambassador.Post("register", controllers.Register)
	ambassador.Post("login", controllers.Login)
	ambassador.Get("products/frontend", controllers.ProductsFrontend)
	ambassador.Get("products/backend", controllers.ProductsBackend)

	ambassadorAuthenticated := ambassador.Use(middleware.IsAuthenticated)
	ambassadorAuthenticated.Get("user", controllers.User)
	ambassadorAuthenticated.Get("logout", controllers.Logout)
	ambassadorAuthenticated.Put("users/info", controllers.UpdateInfo)
	ambassadorAuthenticated.Put("users/password", controllers.UpdatePassword)
	ambassadorAuthenticated.Post("links", controllers.CreateLink)
	ambassadorAuthenticated.Get("stats", controllers.Stats)
}
