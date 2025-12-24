package models

import (
	"time"

	"github.com/google/uuid"
)

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
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
type LoginResponse struct {
	AccessToken  string       `json:"access_token"`
	RefreshToken string       `json:"refresh_token"`
	User         UserResponse `json:"user"`
}
type UserQueryParams struct {
	Keyword string `form:"keyword"`
	Page    int    `form:"page,default=1"`
	Limit   int    `form:"limit,default=5"`
	SortBy  string `form:"sort_by,default=created_at"`
	Order   string `form:"order,default=desc"`
}
type PaginatedUserResponse struct {
	Data       []UserResponse `json:"data"`
	Pagination Pagination     `json:"pagination"`
}
type Pagination struct {
	Current    int `json:"current_page"`
	TotalPage  int `json:"total_page"`
	ToTalItems int `json:"to_tal_items"`
	Limit      int `json:"limit"`
}
