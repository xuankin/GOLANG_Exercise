package controller

import (
	"API_BASE/config"
	"API_BASE/models"
	"API_BASE/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserController struct {
	userService service.UserService
	config      *config.Config
}

func NewUserController(userService service.UserService, config *config.Config) *UserController {
	return &UserController{
		userService: userService,
		config:      config,
	}
}
func (uc *UserController) CreateUser(c *gin.Context) {
	var req models.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userResponse, err := uc.userService.CreateUser(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"userResponse": userResponse})
}
func (uc *UserController) GetAllUsers(c *gin.Context) {
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
func (uc *UserController) GetUserByID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))

	user, err := uc.userService.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})

}
func (uc *UserController) UpdateUser(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))

	var req models.UserUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err, status := uc.userService.UpdateUserByID(id, &req); err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cap nhat thanh cong"})
}
func (uc *UserController) DeleteUserByID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	err := uc.userService.DeleteUserByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"Delete user": id})
}
func (uc *UserController) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := uc.userService.Login(&req, uc.config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Login": res})
}
func (uc *UserController) RefreshToken(c *gin.Context) {
	var req models.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newAccesToken, err := uc.userService.RefreshToken(req.RefreshToken, uc.config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"RefreshToken": newAccesToken})
}
func (uc *UserController) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	refreshToken := strings.TrimPrefix(authHeader, "Bearer ")
	if err := uc.userService.Logout(refreshToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Khong the dang xuat"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Logout": "Dang xuat thanh cong"})
}
