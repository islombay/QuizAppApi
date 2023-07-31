package handler

import (
	"QuizAppApi/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		subjects := api.Group("/subjects")
		{
			subjects.GET("/", h.getAllSubjects)    // return a list of all subjects
			subjects.GET("/:id", h.getSubjectById) // return specific subject with questions

			subjects.GET("/:id/questions") // get all questions for the subject
			subjects.GET("/:id/answer")    // get the correct answers for the subject questions.
		}

		admin := api.Group("/admin")
		{
			admin.POST("/subject", h.createNewSubject)
			admin.DELETE("/subject", h.deleteSubject)

			admin.PUT("/subject", h.updateSubject)
		}
	}

	return router
}
