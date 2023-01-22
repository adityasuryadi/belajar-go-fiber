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

	// route untuk test rabbitmq

	// koneksi ke rabbitmq
	// connRabbitMq, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	// if err != nil {
	// 	log.Fatalf("%s: %s", "failed to connect rabbitMQ", err)
	// }

	// app.Get("/send", func(c *fiber.Ctx) error {
	// 	if c.Query("msg") == "" {
	// 		log.Println("Missing 'msq' query parameter")
	// 	}

	// 	// buka koneksi ke rabbitmq
	// 	ch, err := connRabbitMq.Channel()
	// 	exception.PanicIfNeeded(err)
	// 	defer ch.Close()

	// 	// deklarasi queue yang akan di publish dan di subscribe

	// 	_, err = ch.QueueDeclare(
	// 		"TestQueue", //nama
	// 		false,       //durable
	// 		false,       //delete when unused
	// 		false,       //exclusive
	// 		false,       // no-wait
	// 		nil,         //arguments
	// 	)

	// 	exception.PanicIfNeeded(err)

	// 	dt := make(map[string]interface{})

	// 	dt["name"] = "Adit"
	// 	dt["user_name"] = "adit33"
	// 	dt["email"] = "adit@mail.com"
	// 	body, err := json.Marshal(dt)

	// 	err = ch.Publish(
	// 		"",
	// 		"TestQueue",
	// 		false,
	// 		false,
	// 		amqp.Publishing{
	// 			ContentType: "text/plain",
	// 			Body:        []byte(body),
	// 		},
	// 	)

	// 	exception.PanicIfNeeded(err)

	// 	return nil
	// })

	app.Listen(":3001")
}
