package repository

import (
	"API_BASE/config"
	"API_BASE/entity"
	"API_BASE/models"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindUsers(params models.UserQueryParams) ([]entity.User, int64, error)
	FindById(id uuid.UUID) (*entity.User, error)
	Update(user *entity.User) error
	Delete(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)

	SetSession(session *entity.Session, duration time.Duration) error
	FindSession(token string) (*entity.Session, error)
	DeleteSession(token string) error

	GetSearchCacheUsers(key string) ([]models.UserResponse, int64, error)
	SetSearchCacheUsers(key string, users []models.UserResponse, total int64) error

	ClearAllSearchCache() error

	ClearUserObjectCache(id uuid.UUID) error
}
type userRepository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewUserRepository(db *gorm.DB, rdb *redis.Client) UserRepository {
	return &userRepository{db, rdb}
}

type SearchCachePayload struct {
	Data  []models.UserResponse `json:"data"`
	Total int64                 `json:"total"`
}

func (r *userRepository) GetSearchCacheUsers(key string) ([]models.UserResponse, int64, error) {
	val, err := r.rdb.Get(config.Ctx, key).Result()
	if err != nil {
		return nil, 0, err
	}
	var payload SearchCachePayload
	err = json.Unmarshal([]byte(val), &payload)
	return payload.Data, payload.Total, err
}
func (r *userRepository) SetSearchCacheUsers(key string, users []models.UserResponse, total int64) error {
	payload := SearchCachePayload{Data: users, Total: total}
	data, _ := json.Marshal(payload)
	return r.rdb.Set(config.Ctx, key, string(data), 5*time.Minute).Err()
}
func (r *userRepository) ClearAllSearchCache() error {
	ctx := config.Ctx
	var cursor uint64
	iter := r.rdb.Scan(ctx, cursor, "search:ids:*", 0).Iterator()

	for iter.Next(ctx) {
		r.rdb.Del(ctx, iter.Val())
	}
	return iter.Err()
}
func (r *userRepository) GetManyUserCache(ids []uuid.UUID) (map[uuid.UUID]models.UserResponse, []uuid.UUID, error) {
	if len(ids) == 0 {
		return nil, nil, nil

	}
	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = "user:" + id.String()
	}
	vals, err := r.rdb.MGet(config.Ctx, keys...).Result()
	if err != nil {
		return nil, nil, err
	}
	result := make(map[uuid.UUID]models.UserResponse)
	var missedIds []uuid.UUID
	for i, val := range vals {
		if val == nil {
			missedIds = append(missedIds, ids[i])
			continue
		}
		var user models.UserResponse
		if strVal, ok := val.(string); ok {
			_ = json.Unmarshal([]byte(strVal), &user)
			result[ids[i]] = user
		} else {
			missedIds = append(missedIds, ids[i])
		}

	}
	return result, missedIds, nil
}
func (r *userRepository) SetSession(session *entity.Session, duration time.Duration) error {
	data, _ := json.Marshal(session)
	return r.rdb.Set(config.Ctx, "session"+session.RefreshToken, data, duration).Err()
}
func (r *userRepository) FindSession(token string) (*entity.Session, error) {
	val, err := r.rdb.Get(config.Ctx, "session"+token).Result()
	if err != nil {
		return nil, err
	}
	var session entity.Session
	err = json.Unmarshal([]byte(val), &session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}
func (r *userRepository) SetUserObjectCache(user *models.UserResponse) error {
	data, _ := json.Marshal(user)
	key := "user:" + user.ID.String()
	return r.rdb.Set(config.Ctx, key, string(data), 30*time.Minute).Err()
}
func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}
func (r *userRepository) ClearUserObjectCache(id uuid.UUID) error {
	return r.rdb.Del(config.Ctx, "user:"+id.String()).Err()
}
func (r *userRepository) FindUsers(params models.UserQueryParams) ([]entity.User, int64, error) {
	var users []entity.User
	var total int64
	db := r.db.Model(&entity.User{})
	if params.Keyword != "" {
		keyword := "%" + params.Keyword + "%"
		db = db.Where("name ILIKE ? OR email ILIKE ?", keyword, keyword)
	}
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	orderClause := fmt.Sprintf("%s %s", params.SortBy, params.Order)
	db = db.Order(orderClause)
	offset := (params.Page - 1) * params.Limit
	err := db.Limit(params.Limit).Offset(offset).Find(&users).Error
	return users, total, err
}
func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) DeleteSession(token string) error {
	return r.rdb.Del(config.Ctx, "session"+token).Err()
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
