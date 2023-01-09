package tests

import (
	"encoding/json"
	"go-blog/config"
	"go-blog/controller"
	"go-blog/repository"
	"go-blog/service"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setup(app *fiber.App) {

}

func TestCreateUser(t *testing.T) {
	app := fiber.New(config.NewFiberConfig())
	userController := Setup()
	userController.Route(app)

	payload := strings.NewReader(`{
		"user_name": "adit",
		"name": "Aditya",
		"email": "adit@mail.com",
		"password": "admin",
		"address": "Jl MOh Toha"
	  }`)

	request := httptest.NewRequest(http.MethodPost, "/api/users", payload)
	request.Header.Add("Content-Type", "application/json")
	res, _ := app.Test(request)
	body, _ := ioutil.ReadAll(res.Body)

	response := make(map[string]interface{})
	json.Unmarshal(body, &response)
	data := response["data"].(map[string]interface{})
	parse := response
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, "OK", parse["status"])
	assert.Equal(t, float64(200), parse["code"])
	assert.Equal(t, data["name"], "Aditya")
}

func Setup() controller.UserController {
	configConfig := config.New()
	db := config.NewPostgresDB(configConfig)
	userRepository := repository.NewUserRepository(db)
	socialAccountRepository := repository.NewSocialAccountRepository(db)
	userService := service.NewUserService(userRepository, socialAccountRepository)
	userController := controller.NewUserController(userService)
	return userController
}
