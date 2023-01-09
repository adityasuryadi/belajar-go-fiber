package service

import (
	"errors"
	"go-blog/entity"
	helpers "go-blog/helper"
	"go-blog/model"
	"go-blog/repository"
	utils "go-blog/util"

	"gorm.io/gorm"
)

func NewUserService(repository repository.UserRepository, socialAccountRepository repository.SocialAccountRepository) UserService {
	return &UserServiceImpl{
		UserRepository:          repository,
		SocialAccountRepository: socialAccountRepository,
	}
}

type UserServiceImpl struct {
	UserRepository          repository.UserRepository
	SocialAccountRepository repository.SocialAccountRepository
}

// List implements UserService
func (service *UserServiceImpl) List() (responses []model.UserResponse) {
	users := service.UserRepository.GetAll()
	for _, user := range users {
		responses = append(responses, model.UserResponse{
			Id:        user.Id.String(),
			Name:      user.Name,
			Email:     user.Email,
			UserName:  user.UserName,
			Address:   user.Address,
			CreatedAt: user.CreatedAt,
		})
	}
	return responses
}

func (service *UserServiceImpl) Create(request model.UserCreateRequest) model.UserResponse {
	user := entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: helpers.GetHash([]byte(request.Password)),
		UserName: request.UserName,
		Address:  request.Address,
	}
	service.UserRepository.Insert(user)
	response := model.UserResponse{
		Name:      user.Name,
		Email:     user.Email,
		UserName:  user.UserName,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
	}

	return response
}

func (service *UserServiceImpl) Edit(id string, request model.UserUpdateRequest) model.UserResponse {
	user := entity.User{
		Name:    request.Name,
		Address: request.Address,
	}
	service.UserRepository.Update(id, user)
	response := model.UserResponse{
		Name:      user.Name,
		Email:     user.Email,
		UserName:  user.UserName,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
	}
	return response
}

func (service *UserServiceImpl) Find(id string) (model.UserResponse, string) {
	errorCode := make(chan string, 1)
	user, err := service.UserRepository.Get(id)
	response := model.UserResponse{
		Id:        user.Id.String(),
		Name:      user.Name,
		Email:     user.Email,
		UserName:  user.UserName,
		Address:   user.Address,
		CreatedAt: user.CreatedAt,
	}
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		errorCode <- "404"
	} else if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		errorCode <- "500"
	} else {
		errorCode <- "nil"
	}
	return response, <-errorCode
}

func (service *UserServiceImpl) Delete(id string) string {
	errorCode := make(chan string, 1)
	result, err := service.UserRepository.Destroy(id)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		errorCode <- "500"
	} else if result == 0 && errors.Is(err, gorm.ErrRecordNotFound) {
		errorCode <- "404"
	} else {
		errorCode <- "200"
	}
	return <-errorCode
}

func (service *UserServiceImpl) Login(request model.UserLoginRequest) (tokenJwt, errCode string) {
	errorCode := make(chan string, 1)
	user := entity.User{
		Email:    request.Email,
		Password: request.Password,
	}
	var t string

	result, err := service.UserRepository.Auth(user)

	if err != nil {
		errorCode <- "404"
	} else {
		// Create the Claims
		t, err = utils.ClaimToken(result.Email)

		if err != nil {
			errorCode <- "500"
		} else {
			errorCode <- "nil"
		}
	}

	return t, <-errorCode
}

func (service *UserServiceImpl) LoginOAuth(request model.UserLoginRequest) {
	// path:= config.ConfigGoogle()
	// url:=path.AuthCodeURL("state")
	panic("implement me")
}

func (service *UserServiceImpl) OAuthCallback(email string, provider string) {

	panic("implement me")
}

func (service *UserServiceImpl) FindOrCreateUser(client model.GoogleResponse, provider string) (jwtToken, errCode string) {
	email := client.Email
	providerId := client.ID
	userEntity, errUser := service.UserRepository.FindUserBySlug("email", email)
	socialAccount, err := service.SocialAccountRepository.FindByProviderId(providerId)

	errorCode := make(chan string, 1)
	var token string
	var errToken error
	var user entity.User

	// jika social account exist

	if err != "404" {
		user = socialAccount.User
	} else {
		if errUser != "404" {
			user = userEntity
		} else {
			user = entity.User{
				Email:    email,
				Provider: provider,
				SocialAccount: []entity.SocialAccount{
					{ProviderName: provider, ProviderId: providerId},
				},
			}
			// insert user
			service.UserRepository.Insert(user)
		}
	}

	token, errToken = utils.ClaimToken(user.Email)
	if errToken != nil {
		errorCode <- "500"
	} else {
		errorCode <- "nil"
	}
	return token, <-errorCode
}
