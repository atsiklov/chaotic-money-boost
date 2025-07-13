package handlers

import (
	myerr "backend/internal/errors"
	user "backend/internal/layers/database/userslove"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepo user.Repository
}

func (handler *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": myerr.INVALID_REQUEST})
		return
	}
	userSct, err := handler.userRepo.FindByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, myerr.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": myerr.INTERNAL_ERROR})
		}
		return
	}
	c.JSON(http.StatusOK, userSct)
}

type CreateUserRequest struct {
	Nickname string `json:"nickname" binding:"required,min=3,max=32"`
	Email    string `json:"email" binding:"required,email,max=100"`
}

func (handler *UserHandler) CreateUser(c *gin.Context) {
	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": myerr.INVALID_REQUEST})
		return
	}
	if err := handler.userRepo.Create(c.Request.Context(), &user.User{
		Nickname: req.Nickname,
		Email:    req.Email,
	}); err != nil {
		if errors.Is(err, myerr.ErrUserAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": myerr.INTERNAL_ERROR})
		}
		return
	}
	c.Status(http.StatusCreated)
}

func NewUserHandler(userRepo user.Repository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}
