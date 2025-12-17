package service

import (
	"API_BASE/entity"
	"API_BASE/models"
	"API_BASE/repository"
	"errors"
)

type UserService interface {
	CreateUser(req *models.UserRequest) (*models.UserResponse, error)
	GetAllUsers() ([]models.UserResponse, error)
	GetUserByID(id uint) (*models.UserResponse, error)
	UpdateUserByID(id uint, req *models.UserUpdate) error
	DeleteUserByID(id uint) error
}
type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}
func mapToResponse(user *entity.User) models.UserResponse {
	return models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
func (s *userService) CreateUser(req *models.UserRequest) (*models.UserResponse, error) {
	userEntity := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	if err := s.repo.Create(&userEntity); err != nil {
		return nil, err
	}
	response := mapToResponse(&userEntity)
	return &response, nil
}
func (s *userService) GetAllUsers() ([]models.UserResponse, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var response []models.UserResponse
	for _, user := range users {
		response = append(response, mapToResponse(&user))
	}
	return response, nil
}
func (s *userService) GetUserByID(id uint) (*models.UserResponse, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	response := mapToResponse(user)
	return &response, nil
}
func (s *userService) UpdateUserByID(id uint, req *models.UserUpdate) error {
	user, err := s.repo.FindById(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Password != "" {
		user.Password = req.Password
	}
	return s.repo.Update(user)
}
func (s *userService) DeleteUserByID(id uint) error {
	user, err := s.repo.FindById(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	return s.repo.Delete(user)
}
