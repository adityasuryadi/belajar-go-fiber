package controller

import (
	"go-blog/config"
	"go-blog/exception"
	"go-blog/model"
	"go-blog/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func NewUserController(userService service.UserService) UserController {
	return UserController{UserService: userService}
}

type UserController struct {
	UserService service.UserService
}

func (controller *UserController) Route(app *fiber.App) {
	api := app.Group("api")

	api.Post("/login", controller.Login)

	// api.Use(middleware.Verify())
	api.Get("/users", controller.List)
	api.Get("/test", controller.restricted)
	api.Post("/users", controller.Create)
	api.Put("/users/:id", controller.Edit)
	api.Get("/users/:id", controller.Find)
	api.Delete("/users/:id", controller.Delete)

	// sign in google
	api.Get("/signin", controller.LoginOauth)
	// api.Get(("/signup"))

	app.Get("/auth/:provider/callback", controller.OAuthCallback)
}

func (controller *UserController) List(ctx *fiber.Ctx) error {
	responses := controller.UserService.List()
	return ctx.JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   responses,
	})
}

func (controller *UserController) Create(ctx *fiber.Ctx) error {
	validate := validator.New()
	var request model.UserCreateRequest
	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)
	err = validate.Struct(&request)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		out := make([]exception.ErrorMessage, len(validationErrors))
		for i, fieldError := range validationErrors {
			out[i] = exception.ErrorMessage{fieldError.Field(), exception.GetErrorMsg(fieldError)}
		}
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(model.WebResponse{
				Code:   fiber.StatusBadRequest,
				Status: "BAD_REQUEST",
				Data:   out,
			})
	}

	response := controller.UserService.Create(request)
	return ctx.JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserController) Edit(ctx *fiber.Ctx) error {
	validate := validator.New()
	var request model.UserUpdateRequest
	err := ctx.BodyParser(&request)
	exception.PanicIfNeeded(err)

	id := ctx.Params("id")

	err = validate.Struct(&request)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		out := make([]exception.ErrorMessage, len(validationErrors))
		for i, fieldError := range validationErrors {
			out[i] = exception.ErrorMessage{fieldError.Field(), exception.GetErrorMsg(fieldError)}
		}
		return ctx.
			Status(fiber.StatusBadRequest).
			JSON(model.WebResponse{
				Code:   fiber.StatusBadRequest,
				Status: "BAD_REQUEST",
				Data:   out,
			})
	}

	response := controller.UserService.Edit(id, request)
	return ctx.JSON(model.WebResponse{
		Code:   fiber.StatusOK,
		Status: "OK",
		Data:   response,
	})
}

func (controller *UserController) Find(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, errorCode := controller.UserService.Find(id)
	if errorCode == "404" {
		model.NotFoundResponse(ctx, nil)
		return nil
	} else if errorCode == "500" {
		model.InternalServerErrorResponse(ctx, nil)
		return nil
	}

	model.SuccessResponse(ctx, user)
	return nil
}

func (controller *UserController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	errorCode := controller.UserService.Delete(id)
	if errorCode == "404" {
		model.NotFoundResponse(ctx, nil)
		return nil
	} else if errorCode == "500" {
		model.InternalServerErrorResponse(ctx, nil)
		return nil
	}

	model.SuccessResponse(ctx, nil)
	return nil
}

func (controller *UserController) Login(ctx *fiber.Ctx) error {
	var request model.UserLoginRequest

	// user := ctx.Locals("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// email := claims["email"].(string)
	// return email

	err := ctx.BodyParser(&request)
	if err != nil {
		exception.PanicIfNeeded(err)
	}

	var data = map[string]string{"token": ""}

	token, errorCode := controller.UserService.Login(request)
	if errorCode == "404" {
		model.NotFoundResponse(ctx, nil)
		return nil
	} else if errorCode == "500" {
		model.InternalServerErrorResponse(ctx, nil)
		return nil
	} else {
		data["token"] = token
		model.SuccessResponse(ctx, data)
		return nil
	}

}

// signin oauth
func (controller *UserController) LoginOauth(ctx *fiber.Ctx) error {
	path := config.ConfigGoogle()
	url := path.AuthCodeURL("state")
	return ctx.Redirect(url)
}

func (controller *UserController) OAuthCallback(ctx *fiber.Ctx) error {
	provider := ctx.Params("provider")
	token, err := config.ConfigGoogle().Exchange(ctx.Context(), ctx.FormValue("code"))

	if err != nil {
		exception.PanicIfNeeded(err)
	}
	client := config.GetClient(token.AccessToken)
	jwtToken, errorCode := controller.UserService.FindOrCreateUser(client, provider)
	var data = map[string]string{"token": ""}
	if errorCode == "404" {
		model.NotFoundResponse(ctx, nil)
		return nil
	} else if errorCode == "500" {
		model.InternalServerErrorResponse(ctx, nil)
		return nil
	} else {
		data["token"] = jwtToken
		model.SuccessResponse(ctx, data)
		return nil
	}
}

func (controller *UserController) restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["email"].(string)
	return c.SendString("Welcome " + name)
}
