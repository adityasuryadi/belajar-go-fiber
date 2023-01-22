//go:build wireinject
// +build wireinject

package main

import (
	// "go-blog/config"

	"go-blog/config"
	"go-blog/controller"
	"go-blog/repository"
	"go-blog/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

// var userSet = wire.NewSet(repository.NewUserRepository, service.NewUserService, controller.NewUserController)

// var articleSet = wire.NewSet(repository.NewArticleRepository, repository.NewUserRepository, service.NewArticleService, controller.NewArticleController)

// func InitializedUserController(filenames ...string) controller.UserController {
// 	wire.Build(
// 		config.New,
// 		config.NewMongoDatabase,
// 		userSet,
// 	)
// 	return controller.UserController{}
// }

// func InitializedArticleController(filenames ...string) controller.ArticleController {
// 	wire.Build(
// 		config.New,
// 		config.NewMongoDatabase,
// 		articleSet,
// 	)
// 	return controller.ArticleController{}
// }

// func InitializedServer(filenames ...string) *fiber.App {
// 	wire.Build(
// 		NewApp,
// 	)
// 	return nil
// }

var userSet = wire.NewSet(repository.NewUserRepository, repository.NewSocialAccountRepository, service.NewUserService, controller.NewUserController)

func InitializedUserController(filenames ...string) controller.UserController {
	wire.Build(
		config.New,
		config.NewPostgresDB,
		config.NewRabbitmqConn,
		service.NewRabbitMqService,
		// repository.NewUserRepository,
		// service.NewUserService,
		userSet,
	)
	return controller.UserController{}
}

func InitializedServer(filenames ...string) *fiber.App {
	wire.Build(
		NewApp,
	)
	return nil
}
