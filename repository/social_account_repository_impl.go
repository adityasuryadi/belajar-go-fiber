package repository

import (
	"go-blog/entity"

	"gorm.io/gorm"
)

func NewSocialAccountRepository(database *gorm.DB) SocialAccountRepository {
	return &SocialAccountRepositoryImpl{
		db: database,
	}
}

type SocialAccountRepositoryImpl struct {
	db *gorm.DB
}

// InsertSocialAccount implements SocialAccountRepository
func (*SocialAccountRepositoryImpl) InsertSocialAccount(socialAccount entity.SocialAccount) error {
	panic("unimplemented")
}

// FindByProviderId implements SocialAccountRepository
func (repository *SocialAccountRepositoryImpl) FindByProviderId(providerId string) (entity.SocialAccount, string) {
	var socialAccount entity.SocialAccount
	db := repository.db.Where("provider_id = ?", providerId).First(&socialAccount)

	if db.Error != nil || db.RowsAffected == 0 {
		return socialAccount, "404"
	}

	return socialAccount, "nil"
}
