package repository

import (
	"QuizAppApi"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg DBConfig) (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.Open(fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			cfg.Host,
			cfg.Username,
			cfg.Password,
			cfg.DBName,
			cfg.Port,
			cfg.SSLMode,
		)),
		&gorm.Config{},
	)

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&QuizAppApi.SubjectModel{}, &QuizAppApi.QuestionModel{}, &QuizAppApi.AdminModel{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
