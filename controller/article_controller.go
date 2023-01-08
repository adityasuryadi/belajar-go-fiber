package controller

import (
	"go-blog/exception"
	"go-blog/middleware"
	"go-blog/model"
	"go-blog/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewArticleController(articleService service.ArticleService) ArticleController {
	return ArticleController{
		ArticleService: articleService,
	}
}

type ArticleController struct {
	ArticleService service.ArticleService
}

func (controller *ArticleController) Route(app *fiber.App) {
	api := app.Group("api/articles").Use(middleware.Verify())
	api.Post("/", controller.Create)
	api.Get("/user", controller.GetArticleUser)
}

func (controller *ArticleController) Create(ctx *fiber.Ctx) error {
	var request model.ArticleCreateRequest
	err := ctx.BodyParser(&request)

	// validasi Input
	validate := validator.New()

	err = validate.Struct(&request)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		out := make([]exception.ErrorMessage, len(validationErrors))
		for i, fieldError := range validationErrors {
			out[i] = exception.ErrorMessage{fieldError.Field(), exception.GetErrorMsg(fieldError)}
		}
		model.BadRequestResponse(ctx, out)
		return nil
	}
	user, errorCode := controller.ArticleService.Create(ctx)
	switch errorCode {
	case "500":
		model.InternalServerErrorResponse(ctx, nil)
	case "404":
		model.NotFoundResponse(ctx, nil)
	default:
		model.SuccessResponse(ctx, user)
	}
	return nil
}

func (controller *ArticleController) GetArticleUser(ctx *fiber.Ctx) error {
	response, errorCode := controller.ArticleService.GetArticlesByUser(ctx)
	switch errorCode {
	case "500":
		model.InternalServerErrorResponse(ctx, nil)
		return nil
	default:
		model.SuccessResponse(ctx, response)
		return nil
	}
}
