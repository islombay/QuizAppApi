package repository

import (
	"QuizAppApi"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) GetAdmin(user, password string) (QuizAppApi.AdminModel, error) {
	var admin QuizAppApi.AdminModel
	if res := r.db.Where("\"user\" = ? AND \"password\" = ?", user, password).First(&admin); res.Error != nil {
		return admin, res.Error
	}

	return admin, nil
}
