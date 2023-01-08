package service

import (
	"go-blog/model"

	"github.com/gofiber/fiber/v2"
)

type ArticleService interface {
	Create(ctx *fiber.Ctx) (model.ArticleReponse, string)
	GetArticlesByUser(ctx *fiber.Ctx) ([]model.ArticleReponse, string)
}
