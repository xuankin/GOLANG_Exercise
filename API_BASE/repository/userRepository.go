package repository

import (
	"API_BASE/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindAll() ([]entity.User, error)
	FindById(id uuid.UUID) (*entity.User, error)
	Update(user *entity.User) error
	Delete(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	CreateSession(session *entity.Session) error
	FindSession(token string) (*entity.Session, error)
	DeleteSession(token string) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}
func (r *userRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	return users, err
}
func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
func (r *userRepository) FindSession(token string) (*entity.Session, error) {
	var session entity.Session
	err := r.db.Where("refresh_token=? AND is_revoked= ?", token, false).First(&session).Error
	return &session, err
}
func (r *userRepository) DeleteSession(token string) error {
	return r.db.Model(&entity.Session{}).Where("refresh_token=?", token).Update("is_revoked", true).Error
}
func (r *userRepository) FindById(id uuid.UUID) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error
	return &user, err
}
func (r *userRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}
func (r *userRepository) Delete(user *entity.User) error {
	return r.db.Delete(user).Error
}
func (r *userRepository) CreateSession(session *entity.Session) error {
	return r.db.Create(session).Error
}
