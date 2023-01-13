package repository

import "go-blog/entity"

type UserRepository interface {
	Insert(user entity.User)
	GetAll() (users []entity.User)
	Update(id string, user entity.User) (userEntity entity.User, errorCode string)
	Get(id string) (user entity.User, err error)
	Destroy(id string) (result int64, err error)
	Auth(user entity.User) (entity.User, error)
	FindUserBySlug(slug string, value interface{}) (user entity.User, errorCode string)
	FindUserByProvider(email string, provider string) (user entity.User, errCOde string)
}
