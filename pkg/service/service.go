package service

import (
	"QuizAppApi"
	"QuizAppApi/pkg/repository"
)

type Authorization interface {
	GenerateAdminToken(user, password string) (string, error)
	AdminTokenValid(adminToken string) error
}

type Question interface {
	GetQuestions(sId int) ([]QuizAppApi.QuestionResponse, error)

	UpdateQuestion(qs QuizAppApi.QuestionResponse) error
	DeleteQuestion(sID, qID int) error

	AddQuestion(su QuizAppApi.CreateQuestionBody) (uint, error)
}

type Subject interface {
	GetAll() ([]QuizAppApi.SubjectResponse, error)
	Create(s QuizAppApi.SubjectModel, qs []QuizAppApi.QuestionModel) (int, error)
	ConvertCreate(sb QuizAppApi.CreateNewSubjectBody) (QuizAppApi.SubjectModel, []QuizAppApi.QuestionModel)
	DeleteSubject(f int) error

	GetSubject(sId int) (QuizAppApi.SubjectSingleResponse, error)

	UpdateSubject(su QuizAppApi.SubjectResponse) error
}

type Service struct {
	Authorization
	Question
	Subject
}

func NewService(repository *repository.Repository) *Service {

	return &Service{
		Subject:       NewSubjectService(repository.Subject),
		Question:      NewQuestionService(repository.Question),
		Authorization: NewAuthService(repository.Authorization),
	}
}
