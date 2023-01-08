package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SocialAccount struct {
	gorm.Model
	Id           int       `gorm:"primaryKey;" column:"id"`
	ProviderName string    `gorm:"column:provider_name"`
	ProviderId   string    `gorm:"column:provider_id"`
	UserId       uuid.UUID `gorm:"type:uuid" column:"user_id"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
	User         User      `gorm:"foreignKey:UserId"`
}

func (SocialAccount) TableName() string {
	return "social_account"
}

func (entity *SocialAccount) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *SocialAccount) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
