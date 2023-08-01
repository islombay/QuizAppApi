package handler

import (
	"QuizAppApi"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type getAllSubjectsResponse struct {
	Subjects []QuizAppApi.SubjectResponse `json:"subjects"`
}

func (h *Handler) getAllSubjects(c *gin.Context) {
	subjects, err := h.services.Subject.GetAll()
	if err != nil {
		log.Fatalf("[DB] could not get list of subjects: %s", err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllSubjectsResponse{
		Subjects: subjects,
	})
}

func (h *Handler) getSubjectById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	subject, err := h.services.Subject.GetSubject(id)
	if err != nil {
		log.Fatalf("[DB] could not get single subject: %s", err.Error())
		return
	}
	questions, err := h.services.Question.GetQuestions(id)
	if err != nil {
		log.Fatalf("[DB] could not get single question: %s", err.Error())
		return
	}

	subject.Questions = questions

	c.JSON(http.StatusOK, subject)
}

func (h *Handler) createNewSubject(c *gin.Context) {
	var newSubject QuizAppApi.CreateNewSubjectBody
	if err := c.BindJSON(&newSubject); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	subject, questions := h.services.Subject.ConvertCreate(newSubject)
	subjectId, err := h.services.Subject.Create(subject, questions)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		log.Fatalln("could not create a subject in db")
	}

	c.JSON(http.StatusOK, basicResponse{
		fmt.Sprintf("%d", subjectId),
	})
}

type deleteSubjectBody struct {
	Id int `json:"id"`
}

func (h *Handler) deleteSubject(c *gin.Context) {
	var deleteSub deleteSubjectBody
	if err := c.BindJSON(&deleteSub); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Println(deleteSub)

	err := h.services.DeleteSubject(deleteSub.Id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		log.Fatalln("could not delete object in db")
	}

	c.JSON(http.StatusOK, basicResponse{
		strconv.FormatBool(true),
	})
}

func (h *Handler) updateSubject(c *gin.Context) {
	var updatingSubject QuizAppApi.SubjectResponse
	if err := c.BindJSON(&updatingSubject); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.UpdateSubject(updatingSubject); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, basicResponse{
		strconv.FormatBool(true),
	})
}
