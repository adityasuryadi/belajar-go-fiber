package model

import "time"

type RoleCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type RoleResponse struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:created_at`
}
