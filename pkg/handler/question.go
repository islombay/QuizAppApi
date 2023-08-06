package handler

import (
	"QuizAppApi"
	"QuizAppApi/pkg/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) updateQuestion(c *gin.Context) {
	var updatingQuestion QuizAppApi.QuestionResponse
	if err := c.BindJSON(&updatingQuestion); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateQuestion(updatingQuestion); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, basicResponse{
		true,
	})
}

type deleteQuestionBody struct {
	SubjectId  int `json:"sid"`
	QuestionId int `json:"qid"`
}

func (h *Handler) deleteQuestion(c *gin.Context) {
	var deleteQ deleteQuestionBody
	if err := c.BindJSON(&deleteQ); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Question.DeleteQuestion(deleteQ.SubjectId, deleteQ.QuestionId); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		log.Fatalln("error during the deleting the question")
	}

	c.JSON(http.StatusOK, basicResponse{
		true,
	})
}

func (h *Handler) addQuestion(c *gin.Context) {
	var newQ QuizAppApi.CreateQuestionBody
	if err := c.BindJSON(&newQ); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.services.Subject.GetSubject(int(newQ.SubjectId))

	if err != nil {
		if err.Error() == service.RecordNotFound {
			NewErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}
	_, err = h.services.Question.AddQuestion(newQ)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		log.Fatalln("could not add new question")
	}
	c.JSON(http.StatusOK, basicResponse{
		true,
	})
}
