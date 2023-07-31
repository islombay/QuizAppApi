package repository

import (
	"QuizAppApi"
	"gorm.io/gorm"
	"log"
)

type SubjectPostgres struct {
	db *gorm.DB
}

func NewSubjectPostgres(db *gorm.DB) *SubjectPostgres {
	return &SubjectPostgres{db: db}
}

func (r *SubjectPostgres) GetAll() ([]QuizAppApi.SubjectModel, error) {
	var subjects []QuizAppApi.SubjectModel
	result := r.db.Find(&subjects).Where("")

	if result.Error != nil {
		return nil, result.Error
	}

	return subjects, nil
}

func (r *SubjectPostgres) GetSubject(sId int) (QuizAppApi.SubjectModel, error) {
	var subject QuizAppApi.SubjectModel
	result := r.db.First(&subject, sId)

	if result.Error != nil {
		return subject, result.Error
	}
	return subject, nil
}

func (r *SubjectPostgres) Create(model QuizAppApi.SubjectModel, qs []QuizAppApi.QuestionModel) (int, error) {
	tx := r.db.Begin()
	res := tx.Create(&model)
	if res.Error != nil {
		tx.Rollback()
		log.Fatalf("--- [DB] --- error creating a subject: %s", res.Error.Error())
	}

	for _, e := range qs {
		e.SubjectId = model.ID
		res = tx.Create(&e)
		if res.Error != nil {
			tx.Rollback()
			log.Fatalf("--- [DB] --- error creating questions of subject: %s", res.Error.Error())
		}
	}

	res = tx.Commit()
	if res.Error != nil {
		tx.Rollback()
		log.Fatalf("--- [DB] --- error commiting the transaction: %s", res.Error.Error())
	}
	return int(model.ID), nil
}

func (r *SubjectPostgres) DeleteSubject(f int) error {
	res := r.db.Delete(&QuizAppApi.SubjectModel{}, f)
	if res.Error != nil {
		return res.Error
	}
	res = r.db.Where("subject_id = ?", f).Delete(&QuizAppApi.QuestionModel{})
	if res.Error != nil {
		return res.Error
	}

	return nil
}
