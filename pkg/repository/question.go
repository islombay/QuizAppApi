package repository

import (
	"QuizAppApi"
	"gorm.io/gorm"
)

type QuestionPostgres struct {
	db *gorm.DB
}

func NewQuestionPostgres(db *gorm.DB) *QuestionPostgres {
	return &QuestionPostgres{db: db}
}

func (r *QuestionPostgres) GetQuestions(sId int) ([]QuizAppApi.QuestionModel, error) {
	var questions []QuizAppApi.QuestionModel
	result := r.db.Where("subject_id = ?", sId).Find(&questions)

	if result.Error != nil {
		return questions, result.Error
	}
	return questions, nil
}
