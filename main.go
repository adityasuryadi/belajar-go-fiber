package main

import (
	"go-blog/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewApp() *fiber.App {
	app := fiber.New(config.NewFiberConfig())
	return app
}

func main() {

	// // setup controller
	userController := InitializedUserController()
	// // articleController := InitializedArticleController()

	// // setup fiber
	app := InitializedServer()
	app.Use(recover.New())
	app.Use(cors.New())

	// // Setup Routing
	userController.Route(app)
	// // articleController.Route(app)
	app.Get("/test", func(c *fiber.Ctx) error {
		// logger := logrus.New()

		// file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		// logger.SetOutput(file)

		// logger.Info("Hello Logging")
		// logger.Warn("Hello Logging")
		// logger.Error("Hello Logging")
		return c.SendString("Hello, Adit!")
	})
	app.Listen(":3001")
}
