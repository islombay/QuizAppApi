package service

import (
	"QuizAppApi"
	"QuizAppApi/pkg/repository"
	"reflect"
)

type QuestionService struct {
	repo repository.Question
}

func NewQuestionService(repo repository.Question) *QuestionService {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) GetQuestions(sId int) ([]QuizAppApi.QuestionResponse, error) {
	questions, err := s.repo.GetQuestions(sId)
	if err != nil {
		return nil, err
	}

	var newQuestions []QuizAppApi.QuestionResponse
	for _, e := range questions {
		newQuestions = append(newQuestions, prepareQuestionResponse(e))
	}

	return newQuestions, nil
}

func (s *QuestionService) UpdateQuestion(qs QuizAppApi.QuestionResponse) error {
	var f QuizAppApi.QuestionModel
	reflectCommonFields(&qs, &f)
	return s.repo.UpdateQuestion(f)
}

func prepareQuestionResponse(source QuizAppApi.QuestionModel) QuizAppApi.QuestionResponse {
	var destination QuizAppApi.QuestionResponse

	v1 := reflect.ValueOf(source)
	v2 := reflect.ValueOf(&destination).Elem()

	for i := 0; i < v1.NumField(); i++ {
		field1 := v1.Type().Field(i)
		field2 := v2.FieldByName(field1.Name)

		if field2.IsValid() && field2.CanSet() {
			field2.Set(v1.Field(i))
		}
	}

	return destination
}

func (s *QuestionService) DeleteQuestion(sID, qID int) error {
	return s.repo.DeleteQuestion(sID, qID)
}

func (s *QuestionService) AddQuestion(su QuizAppApi.CreateQuestionBody) (uint, error) {
	var f QuizAppApi.QuestionModel
	reflectCommonFields(&su, &f)
	return s.repo.AddQuestion(f)
}
