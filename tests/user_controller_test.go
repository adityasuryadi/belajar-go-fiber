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

var jwtToken string

// var userId string
var userId string = "5485a216-6a3a-42c2-bb5a-0b6793cc9c02"

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
	assert.Equal(t, data["user_name"], "adit")
	assert.Equal(t, data["email"], "adit@mail.com")
	assert.Equal(t, data["address"], "Jl MOh Toha")
}

/**
* Test Create User
 */
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
		assert.Equal(t, "field tidak boleh kosong", value["message"])
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
		assert.Equal(t, "field tidak boleh kosong", value["message"])
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
		assert.Equal(t, "field tidak boleh kosong", value["message"])
	}
}

func TestCreateWrongEmail(t *testing.T) {
	app := fiber.New(config.NewFiberConfig())
	userController := Setup()
	userController.Route(app)

	payload := strings.NewReader(`{
		"user_name": "adit",
		"name": "Test",
		"email": "test@",
		"password": "1234567",
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
		assert.Equal(t, "format email salah", value["message"])
	}
}

/**
* List Test Users
 */

func TestUnAuthorizedGetUsers(t *testing.T) {
	app := fiber.New(config.NewFiberConfig())
	userController := Setup()
	userController.Route(app)

	request := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer 1234")
	res, _ := app.Test(request)
	body, _ := ioutil.ReadAll(res.Body)

	response := make(map[string]interface{})
	json.Unmarshal(body, &response)
	parse := response
	assert.Equal(t, 401, res.StatusCode)
	assert.Equal(t, "UNAUTHORIZE", parse["status"])
	assert.Equal(t, float64(401), parse["code"])
}

func TestEmptyTokenGetUsers(t *testing.T) {
	app := fiber.New(config.NewFiberConfig())
	userController := Setup()
	userController.Route(app)

	request := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	request.Header.Add("Content-Type", "application/json")
	res, _ := app.Test(request)
	body, _ := ioutil.ReadAll(res.Body)

	response := make(map[string]interface{})
	json.Unmarshal(body, &response)
	parse := response
	assert.Equal(t, 400, res.StatusCode)
	assert.Equal(t, "BAD_REQUEST", parse["status"])
	assert.Equal(t, float64(400), parse["code"])
}

func TestLoginSuccess(t *testing.T) {
	app := fiber.New(config.NewFiberConfig())
	userController := Setup()
	userController.Route(app)

	payload := strings.NewReader(`{
		"email": "adit@mail.com",
		"password": "admin"
	  }`)

	request := httptest.NewRequest(http.MethodPost, "/api/login", payload)
	request.Header.Add("Content-Type", "application/json")
	res, _ := app.Test(request)
	body, _ := ioutil.ReadAll(res.Body)

	response := make(map[string]interface{})
	json.Unmarshal(body, &response)
	parse := response
	data := response["data"].(map[string]interface{})
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, "OK", parse["status"])
	assert.Equal(t, float64(200), parse["code"])

	jwtToken = data["token"].(string)
}

func TestGestListUsersSuccess(t *testing.T) {
	app := fiber.New(config.NewFiberConfig())
	userController := Setup()
	userController.Route(app)

	request := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+jwtToken)
	res, _ := app.Test(request)
	body, _ := ioutil.ReadAll(res.Body)

	response := make(map[string]interface{})
	json.Unmarshal(body, &response)
	parse := response
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, "OK", parse["status"])
	assert.Equal(t, float64(200), parse["code"])
}

/*
* Get User By Id
 */

func TestGetUserByIdSuccess(t *testing.T) {
	app := fiber.New(config.NewFiberConfig())
	userController := Setup()
	userController.Route(app)

	request := httptest.NewRequest(http.MethodGet, "/api/users/b2b029d2-8b44-4b42-94f1-32811caa1ffc", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+jwtToken)
	res, _ := app.Test(request)
	body, _ := ioutil.ReadAll(res.Body)

	response := make(map[string]interface{})
	json.Unmarshal(body, &response)
	parse := response
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, "OK", parse["status"])
	assert.Equal(t, float64(200), parse["code"])
}

func TestGetUserByIdNotFound(t *testing.T) {
	app := fiber.New(config.NewFiberConfig())
	userController := Setup()
	userController.Route(app)

	request := httptest.NewRequest(http.MethodGet, "/api/users/b2b029d2-8b44-4b42-94f1-32811caa1ffd", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+jwtToken)
	res, _ := app.Test(request)
	body, _ := ioutil.ReadAll(res.Body)

	response := make(map[string]interface{})
	json.Unmarshal(body, &response)
	parse := response
	assert.Equal(t, 404, res.StatusCode)
	assert.Equal(t, "NOT_FOUND", parse["status"])
	assert.Equal(t, float64(404), parse["code"])
}

func TestGetUserUnAuthorized(t *testing.T) {
	app := fiber.New(config.NewFiberConfig())
	userController := Setup()
	userController.Route(app)

	request := httptest.NewRequest(http.MethodGet, "/api/users/b2b029d2-8b44-4b42-94f1-32811caa1ffd", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer 123")
	res, _ := app.Test(request)
	body, _ := ioutil.ReadAll(res.Body)

	response := make(map[string]interface{})
	json.Unmarshal(body, &response)
	parse := response
	assert.Equal(t, 401, res.StatusCode)
	assert.Equal(t, "UNAUTHORIZE", parse["status"])
	assert.Equal(t, float64(401), parse["code"])
}

/*
* Test Update User
*
 */

func TestUpdateUserSuccess(t *testing.T) {
	app := fiber.New(config.NewFiberConfig())
	userController := Setup()
	userController.Route(app)

	payload := strings.NewReader(`{
		"name": "Aditya Suryadi",
		"address": "Jl MOh Toha"
	  }`)

	request := httptest.NewRequest(http.MethodPut, "/api/users/"+userId, payload)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+jwtToken)
	res, _ := app.Test(request)
	body, _ := ioutil.ReadAll(res.Body)

	response := make(map[string]interface{})
	json.Unmarshal(body, &response)
	data := response["data"].(map[string]interface{})
	parse := response
	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, "OK", parse["status"])
	assert.Equal(t, float64(200), parse["code"])
	assert.Equal(t, data["name"], "Aditya Suryadi")
	assert.Equal(t, data["email"], "adit@mail.com")
	assert.Equal(t, data["address"], "Jl MOh Toha")
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
