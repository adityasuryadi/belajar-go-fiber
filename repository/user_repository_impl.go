package repository

import (
	"errors"
	"go-blog/entity"
	"go-blog/exception"
	utils "go-blog/util"

	"gorm.io/gorm"
)

func NewUserRepository(database *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: database,
	}
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (repository *UserRepositoryImpl) Insert(user entity.User) {
	result := repository.db.Create(&user)
	exception.PanicIfNeeded(result.Error)
}

func (repository *UserRepositoryImpl) GetAll() (users []entity.User) {
	var user []entity.User
	result := repository.db.Find(&user)
	exception.PanicIfNeeded(result.Error)
	for _, value := range user {
		users = append(users, entity.User{
			Id:        value.Id,
			Name:      value.Name,
			Email:     value.Email,
			Password:  value.Password,
			UserName:  value.UserName,
			Address:   value.Address,
			CreatedAt: value.CreatedAt,
		})
	}
	return users
}

func (repository *UserRepositoryImpl) Update(id string, user entity.User) {
	var entityUser entity.User
	checkedUser := repository.db.First(&entityUser, "id = ?", id)
	if checkedUser.RowsAffected < 1 {
		return
	}
	exception.PanicIfNeeded(checkedUser.Error)
	entityUser.Name = user.Name
	entityUser.Address = user.Address
	repository.db.Save(&entityUser)

}

func (repository *UserRepositoryImpl) Get(id string) (user entity.User, err error) {
	var entityUser entity.User
	result := repository.db.First(&entityUser, "id = ?", id)
	return entityUser, result.Error
}

func (repository *UserRepositoryImpl) Destroy(id string) (result int64, err error) {
	var user entity.User
	db := repository.db.Where("id = ?", id).First(&user)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return 0, db.Error
	}
	db.Delete(&user)
	return 1, nil
}

func (repository *UserRepositoryImpl) Auth(user entity.User) (entity.User, error) {
	var userEntity entity.User
	db := repository.db.Where("email = ?", user.Email).First(&userEntity)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return userEntity, db.Error
	}

	hashPassword := userEntity.Password
	err := utils.ComparePassword(hashPassword, user.Password)
	if err != nil {
		return entity.User{}, err
	}

	return userEntity, nil

}

func (repository *UserRepositoryImpl) FindUserBySlug(slug string, value interface{}) (user entity.User, errCode string) {
	var userEntity entity.User
	db := repository.db.Where(slug+" = ?", value).Preload("SocialAccount").First(&userEntity)
	if db.Error != nil || errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return userEntity, "404"
	}
	return userEntity, "nil"
}

func (repository *UserRepositoryImpl) FindUserByProvider(email string, provider string) (user entity.User, errCOde string) {
	var userEntity entity.User

	db := repository.db.
		Where("email = ?", email).
		Where("provider = ?", provider).
		First(&userEntity)

	if db.Error != nil || db.RowsAffected == 0 {
		return userEntity, "404"
	}
	return userEntity, "nil"
}
