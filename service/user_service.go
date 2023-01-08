package service

import (
	"go-blog/model"
)

type UserService interface {
	List() (responses []model.UserResponse)
	Create(user model.UserCreateRequest) model.UserResponse
	Edit(id string, user model.UserUpdateRequest) model.UserResponse
	Find(id string) (model.UserResponse, string)
	Delete(id string) (errorCode string)
	Login(user model.UserLoginRequest) (tokenJwt, errCode string)
	LoginOAuth(request model.UserLoginRequest)
	OAuthCallback(email string, provider string)
	FindOrCreateUser(client model.GoogleResponse, provider string) (jwtToken, errorCode string)
}
