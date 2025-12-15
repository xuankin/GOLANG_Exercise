package controller

import (
	"Learn_Gin/model"
	"Learn_Gin/service"
	"Learn_Gin/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

// Get / users
func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.userService.GetUsers()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var userResponse []model.UserResponse
	for _, user := range users {
		userResponse = append(userResponse, model.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	utils.SuccessResponse(c, http.StatusOK, userResponse)
}

// Get /user/:id
func (uc *UserController) GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := uc.userService.GetUserById(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if user == nil {
		utils.ErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}
	userResponse := model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
	utils.SuccessResponse(c, http.StatusOK, userResponse)
}

//Post /users

func (uc *UserController) CreateUser(c *gin.Context) {
	var newUser model.CreateUserRequest
	if err := c.ShouldBindJSON(&newUser); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	createdUser, err := uc.userService.CreateUser(newUser)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userResponse := model.UserResponse{
		ID:    createdUser.ID,
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}
	utils.SuccessResponse(c, http.StatusOK, userResponse)
}

// PUT / users/:id
func (uc *UserController) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var updateUser model.UpdateUserRequest
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := uc.userService.UpdateUser(id, updateUser); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "User updated successfully")
}

// Delete /user/:id
func (uc *UserController) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := uc.userService.DeleteUser(id); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, http.StatusOK, "User deleted successfully")
}
