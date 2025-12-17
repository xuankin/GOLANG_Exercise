package models

import "time"

type UserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}
type UserUpdate struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"required,min=6"`
}
type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
