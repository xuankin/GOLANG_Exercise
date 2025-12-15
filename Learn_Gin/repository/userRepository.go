package repository

import (
	"Learn_Gin/entity"
	"encoding/json"
	"os"
)

type UserRepository struct {
	filePath string
}

func NewUserRepository(filePath string) *UserRepository {
	return &UserRepository{
		filePath: filePath,
	}
}

// Ham doc du lieu tu file
func (r *UserRepository) loadData() ([]entity.User, error) {
	var users []entity.User
	if _, err := os.Stat(r.filePath); os.IsNotExist(err) {
		return users, nil
	}
	fileData, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}
	if len(fileData) == 0 {
		return users, nil
	}
	err = json.Unmarshal(fileData, &users)
	return users, err
}

// Ham ghu du lieu vao file
func (r *UserRepository) saveData(users []entity.User) error {
	data, err := json.MarshalIndent(users, "", "   ")
	if err != nil {
		return err
	}
	return os.WriteFile(r.filePath, data, 0644)
}
func (r *UserRepository) FindAll() ([]entity.User, error) {
	return r.loadData()
}
func (r *UserRepository) FindById(id int) (*entity.User, error) {
	users, err := r.loadData()
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, nil
}
func (r *UserRepository) Create(user entity.User) error {
	users, err := r.loadData()
	if err != nil {
		return err
	}
	users = append(users, user)
	return r.saveData(users)
}
func (r *UserRepository) Update(id int, updatedUser entity.User) error {
	users, err := r.loadData()
	if err != nil {
		return err
	}
	for i, user := range users {
		if user.ID == id {
			updatedUser.ID = id
			users[i] = updatedUser
			return r.saveData(users)
		}
	}
	return nil
}
func (r *UserRepository) Delete(id int) error {
	users, err := r.loadData()
	if err != nil {
		return err
	}
	newUsers := make([]entity.User, 0)
	for _, user := range users {
		if user.ID != id {
			newUsers = append(newUsers, user)
		}
	}
	return r.saveData(newUsers)
}
