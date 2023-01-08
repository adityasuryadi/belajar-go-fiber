package service

import (
	"go-blog/model"
	"go-blog/repository"

	"github.com/gofiber/fiber/v2"
)

func NewArticleService(repository repository.ArticleRepository, userRepository repository.UserRepository) ArticleService {
	return &ArticleServiceImpl{
		ArticleRepository: repository,
		UserRepository:    userRepository,
	}
}

type ArticleServiceImpl struct {
	ArticleRepository repository.ArticleRepository
	UserRepository    repository.UserRepository
}

func (service *ArticleServiceImpl) Create(ctx *fiber.Ctx) (model.ArticleReponse, string) {
	// var article entity.Article
	// var request model.ArticleCreateRequest
	// err := ctx.BodyParser(&request)

	// errorCode := make(chan string, 1)

	// tokenUser := ctx.Locals("user").(*jwt.Token)
	// claims := tokenUser.Claims.(jwt.MapClaims)
	// email := claims["email"].(string)
	// // cari user id by email

	// userAuth, _ := service.UserRepository.FindUserBySlug("email", email)

	// article = entity.Article{
	// 	Title:       request.Title,
	// 	Description: request.Description,
	// 	UserId:      userAuth.Id,
	// }
	// user, err := service.ArticleRepository.Insert(article)
	// response := model.ArticleReponse{
	// 	Title:       user.Title,
	// 	Description: user.Description,
	// }

	// if err != nil {
	// 	errorCode <- "500"
	// } else {
	// 	errorCode <- "nil"
	// }

	// return response, <-errorCode
	panic("")
}

func (service *ArticleServiceImpl) GetArticlesByUser(ctx *fiber.Ctx) ([]model.ArticleReponse, string) {
	// var responses []model.ArticleReponse
	// user := ctx.Locals("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// email := claims["email"].(string)

	// userAuth, _ := service.UserRepository.FindUserBySlug("email", email)

	// articles, errorCode := service.ArticleRepository.FindArticleUser(userAuth.Id)

	// for _, article := range articles {
	// 	responses = append(responses, model.ArticleReponse{
	// 		Id:          article.Id,
	// 		Title:       article.Title,
	// 		Description: article.Description,
	// 		AuthorName:  article.UserId,
	// 		CreatedAt:   article.CreatedAt,
	// 	})
	// }

	// return responses, errorCode
	panic("")

}
