package handlers

import (
	myerr "backend/internal/errors"
	"backend/internal/layers/database/challenge/showcase"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ShowcaseHandler struct {
	showRepo showcase.Repository
}

func (handler *ShowcaseHandler) GetShowcases(c *gin.Context) {
	chgesShow, err := handler.showRepo.FindAll(c.Request.Context()) // len == 0 ошибкой не считаем
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": myerr.INTERNAL_ERROR})
		return
	}
	c.JSON(http.StatusOK, chgesShow)
}

func (handler *ShowcaseHandler) GetShowcase(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": myerr.INVALID_REQUEST})
		return
	}
	chgeShow, err := handler.showRepo.FindByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, myerr.ErrChgeShowcaseNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "challenge not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": myerr.INTERNAL_ERROR})
		}
		return
	}
	c.JSON(http.StatusOK, chgeShow)
}

func (handler *ShowcaseHandler) GetChallengeStatus(c *gin.Context) {
	// todo - пока под вопросом
}

func NewShowHandler(showRepo showcase.Repository) *ShowcaseHandler {
	return &ShowcaseHandler{showRepo: showRepo}
}
