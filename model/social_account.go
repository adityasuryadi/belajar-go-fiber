package model

import "time"

type SocialAccount struct {
	Id           int       `json:"id"`
	ProviderName string    `json:"provider_name"`
	ProviderId   string    `json:"provider_id"`
	UserId       string    `json:"user_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
