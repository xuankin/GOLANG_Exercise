package service

import (
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
func (s *UserService) GetUsers() ([]model.User, error) {
	return s.repo.FindAll()
}
func (s *UserService) GetUserById(id int) (*model.User, error) {
	return s.repo.FindById(id)
}
func (s *UserService) CreateUser(user *model.User) error {
	users, _ := s.repo.FindAll()
	maxId := 0
	for i := 0; i < len(users); i++ {
		if users[i].ID > maxId {
			maxId = users[i].ID
		}
	}
	user.ID = maxId + 1
	return s.repo.Create(*user)
}
func (s *UserService) UpdateUser(id int, user model.User) error {
	return s.repo.Update(id, user)
}
func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
