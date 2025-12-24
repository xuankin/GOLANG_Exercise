package service

import (
	"API_BASE/config"
	"API_BASE/entity"
	"API_BASE/models"
	"API_BASE/repository"
	"API_BASE/utils"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(req *models.UserRequest) (*models.UserResponse, error)
	GetAllUsers(params models.UserQueryParams) (*models.PaginatedUserResponse, error)
	GetUserByID(id uuid.UUID) (*models.UserResponse, error)
	UpdateUserByID(id uuid.UUID, req *models.UserUpdate) (error, int)
	DeleteUserByID(id uuid.UUID) error
	Login(req *models.LoginRequest, conf *config.Config) (*models.LoginResponse, error)
	RefreshToken(token string, conf *config.Config) (*models.LoginResponse, error)
	Logout(token string) error
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
	hashedPassword, _ := utils.HashPassword(req.Password)
	userEntity := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}
	if err := s.repo.Create(&userEntity); err != nil {
		return nil, err
	}

	response := mapToResponse(&userEntity)
	go s.repo.ClearAllSearchCache()
	return &response, nil
}
func generateCacheKey(p models.UserQueryParams) string {
	raw := fmt.Sprintf("k=%s|p=%s|l=%d|s=%s|o=%s", p.Keyword, p.Page, p.Limit, p.SortBy, p.Order)
	hasher := md5.New()
	hasher.Write([]byte(raw))
	return "search:ids:" + hex.EncodeToString(hasher.Sum(nil))
}
func (s *userService) GetAllUsers(params models.UserQueryParams) (*models.PaginatedUserResponse, error) {
	cacheKey := generateCacheKey(params)
	cachedUsers, totalItems, err := s.repo.GetSearchCacheUsers(cacheKey)

	if err != nil {
		userEntity, total, errDB := s.repo.FindUsers(params)
		if errDB != nil {
			return nil, errDB
		}
		totalItems = total
		cachedUsers = []models.UserResponse{}
		for _, user := range userEntity {
			cachedUsers = append(cachedUsers, mapToResponse(&user))
		}
		go s.repo.SetSearchCacheUsers(cacheKey, cachedUsers, total)
	}
	totalPages := int(math.Ceil(float64(totalItems) / float64(params.Limit)))
	if totalItems == 0 {
		totalPages = 0
	}
	return &models.PaginatedUserResponse{
		Data: cachedUsers,
		Pagination: models.Pagination{
			Current:    params.Page,
			TotalPage:  totalPages,
			ToTalItems: int(totalItems),
			Limit:      params.Limit,
		},
	}, nil
}
func (s *userService) GetUserByID(id uuid.UUID) (*models.UserResponse, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	response := mapToResponse(user)
	return &response, nil
}
func (s *userService) UpdateUserByID(id uuid.UUID, req *models.UserUpdate) (error, int) {
	user, err := s.repo.FindById(id)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user not found"), http.StatusNotFound
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
	err = s.repo.Update(user)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	go s.repo.ClearAllSearchCache()

	return nil, http.StatusOK
}
func (s *userService) DeleteUserByID(id uuid.UUID) error {
	user, err := s.repo.FindById(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}
	if err := s.repo.ClearUserObjectCache(id); err != nil {
		return err
	}
	return s.repo.ClearAllSearchCache()
}

func (s *userService) Login(req *models.LoginRequest, conf *config.Config) (*models.LoginResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid password")
	}
	acessToken, _ := utils.GenerateToken(user.ID, conf.JWTSecret, conf.JWTDuration)
	refreshToken, _ := utils.GenerateToken(user.ID, conf.JWTRefreshSecret, conf.JWTRefreshDuration)
	hashedRT := utils.HashToken(refreshToken)
	session := &entity.Session{
		UserID:       user.ID,
		RefreshToken: hashedRT,
		ExpiresAt:    time.Now().Add(conf.JWTRefreshDuration),
	}
	s.repo.SetSession(session, conf.JWTRefreshDuration)
	return &models.LoginResponse{
		AccessToken:  acessToken,
		RefreshToken: refreshToken,
		User:         mapToResponse(user),
	}, nil
}
func (s *userService) RefreshToken(token string, conf *config.Config) (*models.LoginResponse, error) {
	hashedRT := utils.HashToken(token)
	session, err := s.repo.FindSession(hashedRT)
	if err != nil {
		return nil, errors.New("Token expired or invalid")
	}
	if session.ExpiresAt.Before(time.Now()) {
		s.repo.DeleteSession(hashedRT)
		return nil, errors.New("Session expired")
	}
	s.repo.DeleteSession(hashedRT)
	newAccessToken, _ := utils.GenerateToken(session.ID, conf.JWTSecret, conf.JWTDuration)
	newRefreshToken, _ := utils.GenerateToken(session.ID, conf.JWTRefreshSecret, conf.JWTRefreshDuration)
	newHashedRT := utils.HashToken(newRefreshToken)
	newSession := &entity.Session{
		UserID:       session.UserID,
		RefreshToken: newHashedRT,
		ExpiresAt:    session.ExpiresAt,
		IsRevoked:    false,
	}
	s.repo.SetSession(newSession, conf.JWTRefreshDuration)
	user, _ := s.repo.FindById(session.UserID)
	return &models.LoginResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
		User:         mapToResponse(user),
	}, nil
}
func (s *userService) Logout(token string) error {
	hashedRT := utils.HashToken(token)
	return s.repo.DeleteSession(hashedRT)
}
