package entity

import "time"

type Article struct {
	Id          string
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	UserId      string    `bson:"user_id"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}
