package entity

import "time"

type User struct {
	Id        string
	Name      string
	Email     string
	Password  string
	UserName  string
	Address   string
	CreatedAt time.Time `bson:created_at`
}
