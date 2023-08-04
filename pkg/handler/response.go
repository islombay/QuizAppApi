package handler

import (
	"QuizAppApi/pkg/service"
	"github.com/gin-gonic/gin"
	"log"
)

const (
	ErrorAuthBinding    = "Provide authorization credentials"
	ErrorRecordNotFound = "Not found"
)

type errorResponse struct {
	Message string `json:"message"`
}

type basicResponse struct {
	Message interface{} `json:"message"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string, opts ...string) {
	for _, opt := range opts {
		if opt == "auth binding" {
			message = ErrorAuthBinding
		}
	}
	if message == service.RecordNotFound {
		message = ErrorRecordNotFound
	}
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
	log.Printf("--- [ERROR] --- %s\n", message)
}
