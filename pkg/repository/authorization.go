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

func (r *AuthPostgres) AddAdmin(user, password string) error {
	newAdmin := QuizAppApi.AdminModel{
		User:     user,
		Password: password,
	}
	if res := r.db.Create(&newAdmin); res.Error != nil {
		return res.Error
	}
	return nil
}
