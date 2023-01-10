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
	assert.Equal(t, data["user_name"], "Adit")
	assert.Equal(t, data["email"], "adit@mail.com")
	assert.Equal(t, data["address"], "Jl MOh Toha")
}

func TestCreateEmptyUserName(t *testing.T) {
	app := fiber.New(config.NewFiberConfig())
	userController := Setup()
	userController.Route(app)

	payload := strings.NewReader(`{
		"user_name": "",
		"name": "Test",
		"email": "test@mail.com",
		"password": "test",
		"address": "Jl MOh Toha"
	  }`)

	request := httptest.NewRequest(http.MethodPost, "/api/users", payload)
	request.Header.Add("Content-Type", "application/json")
	res, _ := app.Test(request)
	body, _ := ioutil.ReadAll(res.Body)

	response := make(map[string]interface{})
	json.Unmarshal(body, &response)
	data := response["data"].([]interface{})
	parse := response
	assert.Equal(t, 400, res.StatusCode)
	assert.Equal(t, "BAD_REQUEST", parse["status"])
	assert.Equal(t, float64(400), parse["code"])

	for _, val := range data {
		value := val.(map[string]interface{})
		assert.Equal(t, "user_name", value["field"])
		assert.Equal(t, "This field is required", value["message"])
	}
}

func TestCreateEmptyEmail(t *testing.T) {
	app := fiber.New(config.NewFiberConfig())
	userController := Setup()
	userController.Route(app)

	payload := strings.NewReader(`{
		"user_name": "adit",
		"name": "Test",
		"email": "",
		"password": "test",
		"address": "Jl MOh Toha"
	  }`)

	request := httptest.NewRequest(http.MethodPost, "/api/users", payload)
	request.Header.Add("Content-Type", "application/json")
	res, _ := app.Test(request)
	body, _ := ioutil.ReadAll(res.Body)

	response := make(map[string]interface{})
	json.Unmarshal(body, &response)
	data := response["data"].([]interface{})
	parse := response
	assert.Equal(t, 400, res.StatusCode)
	assert.Equal(t, "BAD_REQUEST", parse["status"])
	assert.Equal(t, float64(400), parse["code"])

	for _, val := range data {
		value := val.(map[string]interface{})
		assert.Equal(t, "email", value["field"])
		assert.Equal(t, "This field is required", value["message"])
	}
}

func TestCreateEmptyPassword(t *testing.T) {
	app := fiber.New(config.NewFiberConfig())
	userController := Setup()
	userController.Route(app)

	payload := strings.NewReader(`{
		"user_name": "adit",
		"name": "Test",
		"email": "test@mail.com",
		"password": "",
		"address": "Jl MOh Toha"
	  }`)

	request := httptest.NewRequest(http.MethodPost, "/api/users", payload)
	request.Header.Add("Content-Type", "application/json")
	res, _ := app.Test(request)
	body, _ := ioutil.ReadAll(res.Body)

	response := make(map[string]interface{})
	json.Unmarshal(body, &response)
	data := response["data"].([]interface{})
	parse := response
	assert.Equal(t, 400, res.StatusCode)
	assert.Equal(t, "BAD_REQUEST", parse["status"])
	assert.Equal(t, float64(400), parse["code"])

	for _, val := range data {
		value := val.(map[string]interface{})
		assert.Equal(t, "password", value["field"])
		assert.Equal(t, "This field is required", value["message"])
	}
}

func Setup() controller.UserController {
	configConfig := config.New()
	db := config.NewTestPostgresDB(configConfig)
	userRepository := repository.NewUserRepository(db)
	socialAccountRepository := repository.NewSocialAccountRepository(db)
	userService := service.NewUserService(userRepository, socialAccountRepository)
	userController := controller.NewUserController(userService)
	return userController
}
