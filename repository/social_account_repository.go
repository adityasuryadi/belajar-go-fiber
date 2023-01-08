package repository

import (
	"go-blog/entity"
)

type SocialAccountRepository interface {
	FindByProviderId(providerId string) (entity.SocialAccount, string)
	InsertSocialAccount(socialAccount entity.SocialAccount) error
}
