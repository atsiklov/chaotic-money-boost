package handlers

import (
	myerr "backend/internal/errors"
	enums "backend/internal/layers"
	assignment "backend/internal/layers/database/challenge/assignmnt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AssignmentHandler struct {
	asgnRepo assignment.Repository
}

type CreateAssignmentRequest struct {
	UserId int64 `json:"userId" binding:"required"`
}

func (handler *AssignmentHandler) CreateAssignment(c *gin.Context) {
	instId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": myerr.INVALID_REQUEST})
		return
	}
	var req CreateAssignmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": myerr.INVALID_REQUEST})
		return
	}
	if err := handler.asgnRepo.Create(c.Request.Context(), &assignment.Assignment{
		Status: enums.ASGN_IN_PROGRESS,
		InstID: instId,
		UserID: req.UserId,
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": myerr.INTERNAL_ERROR})
		return
	}
	c.Status(http.StatusCreated)
}

type UpdateAssignmentRequest struct {
	UserId     int64  `json:"userId" binding:"required"`
	Submission string `json:"submission" binding:"required"`
}

func (handler *AssignmentHandler) UpdateAssignment(c *gin.Context) {
	instId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": myerr.INVALID_REQUEST})
		return
	}

	var req UpdateAssignmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": myerr.INVALID_REQUEST})
		return
	}
	if err := handler.asgnRepo.Update(c.Request.Context(), &assignment.Assignment{
		Status:     enums.ASGN_SUBMITTED,
		UserID:     req.UserId,
		InstID:     instId,
		Submission: req.Submission,
	}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": myerr.INTERNAL_ERROR})
		return
	}
	c.Status(http.StatusOK)
}

func NewAsgnHandler(asgnRepo assignment.Repository) *AssignmentHandler {
	return &AssignmentHandler{asgnRepo: asgnRepo}
}
