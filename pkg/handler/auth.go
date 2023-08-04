package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type signInBody struct {
	Username string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var credentials signInBody
	if err := c.BindJSON(&credentials); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error(), "auth binding")
		return
	}

	token, err := h.services.Authorization.GenerateAdminToken(credentials.Username, credentials.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, basicResponse{
		token,
	})
}
