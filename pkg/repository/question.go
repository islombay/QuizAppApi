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

func (r *QuestionPostgres) UpdateQuestion(qs QuizAppApi.QuestionModel) error {
	var q QuizAppApi.QuestionModel
	if res := r.db.Where("subject_id = ? AND question_id = ?", qs.SubjectId, qs.QuestionId).First(&q); res.Error != nil {
		return res.Error
	}

	q.Text = qs.Text
	q.Answer1 = qs.Answer1
	q.Answer2 = qs.Answer2
	q.Answer3 = qs.Answer3
	q.Answer4 = qs.Answer4
	q.CorrectAnswer = qs.CorrectAnswer
	q.Level = qs.Level

	if res := r.db.Save(&q); res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *QuestionPostgres) DeleteQuestion(sID, qID int) error {
	if res := r.db.Where("subject_id = ? AND question_id = ?", sID, qID).Delete(&QuizAppApi.QuestionModel{}); res.Error != nil {
		return res.Error
	}
	if err := r.resetQuestionsID(sID); err != nil {
		return err
	}
	return nil
}

func (r *QuestionPostgres) AddQuestion(qs QuizAppApi.QuestionModel) (uint, error) {
	res := r.db.Create(&qs)
	if res.Error != nil {
		return 0, res.Error
	}

	if err := r.resetQuestionsID(int(qs.SubjectId)); err != nil {
		return 0, err
	}

	return qs.QuestionId, nil
}

func (r *QuestionPostgres) resetQuestionsID(sID int) error {
	questions, err := r.GetQuestions(sID)
	if err != nil {
		return err
	}

	for i, e := range questions {
		e.QuestionId = uint(i + 1)
		if res := r.db.Save(&e); res.Error != nil {
			return res.Error
		}
	}

	return nil
}
