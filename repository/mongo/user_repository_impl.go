package repository

import (
	"go-blog/config"
	"go-blog/entity"
	"go-blog/exception"
	helpers "go-blog/helper"
	utils "go-blog/util"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) User {
	return &UserRepositoryImpl{
		Collection: database.Collection("users"),
	}
}

type UserRepositoryImpl struct {
	Collection *mongo.Collection
}

// insert user
func (repository *UserRepositoryImpl) Insert(user entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := repository.Collection.InsertOne(ctx, bson.M{
		// "_id":       user.Id,
		"user_name":  user.UserName,
		"email":      user.Email,
		"password":   helpers.GetHash([]byte(user.Password)),
		"name":       user.Name,
		"address":    user.Address,
		"created_at": time.Now().UTC(),
	})
	exception.PanicIfNeeded(err)
}

// GetAll implements UserRepository
func (repository *UserRepositoryImpl) GetAll() (users []entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	cursor, err := repository.Collection.Find(ctx, bson.M{})
	exception.PanicIfNeeded(err)

	var datas []bson.M
	err = cursor.All(ctx, &datas)
	exception.PanicIfNeeded(err)
	for _, data := range datas {
		users = append(users, entity.User{
			Id:        helpers.PrimitveObjToString(data["_id"]),
			Name:      data["name"].(string),
			Email:     data["email"].(string),
			UserName:  data["user_name"].(string),
			CreatedAt: (data["created_at"].(primitive.DateTime)).Time(),
		})
	}

	return users
}

func (repository *UserRepositoryImpl) Update(id string, user entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)
	selector := bson.M{"_id": objId}
	_, err := repository.Collection.UpdateOne(ctx, selector, bson.M{
		"$set": bson.M{
			"name":    user.Name,
			"address": user.Address,
		},
	})
	// res, err := repository.Collection.Find(ctx, selector)
	exception.PanicIfNeeded(err)
}

func (repository *UserRepositoryImpl) Get(id string) (user entity.User, err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	var result bson.M
	objId, _ := primitive.ObjectIDFromHex(id)
	selector := bson.M{"_id": objId}
	err = repository.Collection.FindOne(ctx, selector).Decode(&result)
	// exception.PanicIfNeeded(err)
	var response entity.User
	if err == nil {
		response = entity.User{
			Id:        helpers.PrimitveObjToString(result["_id"]),
			Name:      result["name"].(string),
			Email:     result["email"].(string),
			UserName:  result["user_name"].(string),
			CreatedAt: (result["created_at"].(primitive.DateTime)).Time(),
		}
	}
	return response, err
}

func (repository *UserRepositoryImpl) Destroy(id string) (result int64, err error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)
	selector := bson.M{"_id": objId}
	res, err := repository.Collection.DeleteOne(ctx, selector)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, err
}

func (repository *UserRepositoryImpl) Auth(user entity.User) (entity.User, error) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	selector := bson.M{
		"email": user.Email,
	}
	var result bson.M
	err := repository.Collection.FindOne(ctx, selector).Decode(&result)
	if err != nil {
		// exception.PanicIfNeeded(err)
		return entity.User{}, err
	}
	// logrus.Warn(user.Password)
	hashPassword := result["password"].(string)
	err = utils.ComparePassword(hashPassword, user.Password)
	if err != nil {
		// exception.PanicIfNeeded(err)
		return entity.User{}, err
	}
	user = entity.User{
		Id:        helpers.PrimitveObjToString(result["_id"]),
		Name:      result["name"].(string),
		Email:     result["email"].(string),
		UserName:  result["user_name"].(string),
		Address:   result["address"].(string),
		CreatedAt: (result["created_at"].(primitive.DateTime)).Time(),
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindUserBySlug(slug string, value interface{}) (user entity.User, errCode string) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	errorCode := make(chan string, 1)

	var result bson.M
	selector := bson.M{slug: value}
	err := repository.Collection.FindOne(ctx, selector).Decode(&result)
	var response entity.User
	if err != nil && err != mongo.ErrNoDocuments {
		errorCode <- "500"
	} else if err != nil && err == mongo.ErrNoDocuments {
		errorCode <- "404"
	} else {
		response = entity.User{
			Id:        helpers.PrimitveObjToString(result["_id"]),
			Name:      result["name"].(string),
			Email:     result["email"].(string),
			UserName:  result["user_name"].(string),
			CreatedAt: (result["created_at"].(primitive.DateTime)).Time(),
		}
		errorCode <- "nill"
	}

	return response, <-errorCode
}
