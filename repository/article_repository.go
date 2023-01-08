package repository

import "go-blog/entity"

type ArticleRepository interface {
	Insert(entity.Article) (entity.Article, error)
	FindArticleUser(userId string) ([]entity.Article, string)
}
