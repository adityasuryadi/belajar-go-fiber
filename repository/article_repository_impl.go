package repository

import (
	"go-blog/config"
	"go-blog/entity"
	"go-blog/exception"
	helpers "go-blog/helper"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewArticleRepository(database *mongo.Database) ArticleRepository {
	return &ArticleRespositoryImpl{Collection: database.Collection("articles")}
}

type ArticleRespositoryImpl struct {
	Collection *mongo.Collection
}

func (repository *ArticleRespositoryImpl) Insert(article entity.Article) (entity.Article, error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	// insert to db proses
	result, err := repository.Collection.InsertOne(ctx, bson.M{
		"title":       article.Title,
		"description": article.Description,
		"user_id":     article.UserId,
		"created_at":  time.Now().UTC(),
		"updated_at":  nil,
	})

	response := entity.Article{
		Id:          helpers.PrimitveObjToString(result.InsertedID),
		Title:       article.Title,
		Description: article.Description,
		UserId:      article.UserId,
		CreatedAt:   time.Now().UTC(),
	}
	return response, err
}

func (repository *ArticleRespositoryImpl) FindArticleUser(userId string) ([]entity.Article, string) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	// objId, x := primitive.ObjectIDFromHex(userId)
	// exception.PanicIfNeeded(x)
	// filter := bson.M{
	// 	"user_id": userId,
	// }
	pipeline := mongo.Pipeline{
		{{
			"$lookup", bson.M{
				"from":         "users",
				"localField":   "user_id",
				"foreignField": "_id",
				"as":           "user_article",
			},
		}},
		{{
			"$project", bson.M{
				"user_article": bson.M{
					"_id":  1,
					"name": 1,
				},
				"title":       1,
				"description": 1,
				"created_at":  1,
			},
		}},
	}
	cursor, err := repository.Collection.Aggregate(ctx, pipeline)

	var datas []bson.M
	var articles []entity.Article
	errorCode := make(chan string, 1)

	err = cursor.All(ctx, &datas)
	logrus.Warn(datas)
	if err != nil {
		errorCode <- "500"
	} else {
		errorCode <- "nil"
	}

	for _, data := range datas {
		articles = append(articles, entity.Article{
			Id:          helpers.PrimitveObjToString(data["_id"]),
			Title:       data["title"].(string),
			Description: data["description"].(string),
			UserId:      data["user_id"].(string),
			CreatedAt:   (data["created_at"].(primitive.DateTime)).Time(),
		})
	}

	if err != nil {
		exception.PanicIfNeeded(err)
	}
	return articles, <-errorCode
}
