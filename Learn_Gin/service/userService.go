package service

import (
	"Learn_Gin/entity"
	"Learn_Gin/model"
	"Learn_Gin/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}
func (s *UserService) GetUsers() ([]entity.User, error) {
	return s.repo.FindAll()
}
func (s *UserService) GetUserById(id int) (*entity.User, error) {
	return s.repo.FindById(id)
}
func (s *UserService) CreateUser(user model.CreateUserRequest) (*entity.User, error) {
	users, _ := s.repo.FindAll()
	maxId := 0
	for i := 0; i < len(users); i++ {
		if users[i].ID > maxId {
			maxId = users[i].ID
		}
	}
	newUser := entity.User{
		ID:    maxId + 1,
		Name:  user.Name,
		Email: user.Email,
	}
	err := s.repo.Create(newUser)
	return &newUser, err
}
func (s *UserService) UpdateUser(id int, req model.UpdateUserRequest) error {
	updateData := entity.User{
		Name:  req.Name,
		Email: req.Email,
	}
	return s.repo.Update(id, updateData)
}
func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
