package repository

import (
	"QuizAppApi"
	"gorm.io/gorm"
)

type Subject interface {
	GetAll() ([]QuizAppApi.SubjectModel, error)
	Create(model QuizAppApi.SubjectModel, qs []QuizAppApi.QuestionModel) (int, error)
	DeleteSubject(f int) error

	GetSubject(sId int) (QuizAppApi.SubjectModel, error)

	UpdateSubject(su QuizAppApi.SubjectModel) error
}

type Authorization interface{}

type Question interface {
	GetQuestions(sId int) ([]QuizAppApi.QuestionModel, error)
}

type Repository struct {
	Subject
	Authorization
	Question
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Subject:  NewSubjectPostgres(db),
		Question: NewQuestionPostgres(db),
	}
}
