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
		}

		admin := api.Group("/admin")
		{
			subject := admin.Group("/subject")
			{
				subject.POST("/", h.createNewSubject)
				subject.DELETE("/", h.deleteSubject)
				subject.PUT("/", h.updateSubject)
			}

			questionGroup := admin.Group("/question")
			{
				questionGroup.PUT("/", h.updateQuestion)
				questionGroup.DELETE("/", h.deleteQuestion)
				questionGroup.POST("/", h.addQuestion)
			}
		}
	}

	return router
}
