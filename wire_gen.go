// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"go-blog/config"
	"go-blog/controller"
	"go-blog/repository"
	"go-blog/service"
)

// Injectors from injector.go:

func InitializedUserController(filenames ...string) controller.UserController {
	configConfig := config.New(filenames...)
	db := config.NewPostgresDB(configConfig)
	userRepository := repository.NewUserRepository(db)
	socialAccountRepository := repository.NewSocialAccountRepository(db)
	connection := config.NewRabbitmqConn(configConfig)
	rabbitMqService := service.NewRabbitMqService(connection)
	userService := service.NewUserService(userRepository, socialAccountRepository, rabbitMqService)
	userController := controller.NewUserController(userService)
	return userController
}

func InitializedServer(filenames ...string) *fiber.App {
	app := NewApp()
	return app
}

// injector.go:

var userSet = wire.NewSet(repository.NewUserRepository, repository.NewSocialAccountRepository, service.NewUserService, controller.NewUserController)
