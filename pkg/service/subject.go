package service

import (
	"QuizAppApi"
	"QuizAppApi/pkg/repository"
	"reflect"
)

type SubjectService struct {
	repo repository.Subject
}

func NewSubjectService(r repository.Subject) *SubjectService {
	return &SubjectService{repo: r}
}

func (s *SubjectService) GetAll() ([]QuizAppApi.SubjectResponse, error) {
	res, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	var response []QuizAppApi.SubjectResponse
	for _, e := range res {
		r := QuizAppApi.SubjectResponse{}
		reflectCommonFields(&e, &r)
		response = append(response, r)
		//response = append(response, prepareSubjectResponse(e))
	}
	return response, nil
}

func prepareSubjectResponse(source QuizAppApi.SubjectModel) QuizAppApi.SubjectResponse {
	var destination QuizAppApi.SubjectResponse

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

func reflectCommonFields(src interface{}, dest interface{}) {
	vSrc := reflect.ValueOf(src).Elem()
	vDest := reflect.ValueOf(dest).Elem()

	for i := 0; i < vSrc.NumField(); i++ {
		fieldSrc := vSrc.Field(i)
		fieldDest := vDest.FieldByName(vSrc.Type().Field(i).Name)

		if fieldDest.IsValid() && fieldDest.CanSet() {
			fieldDest.Set(fieldSrc)
		}
	}
}

func (s *SubjectService) Create(sb QuizAppApi.SubjectModel, qs []QuizAppApi.QuestionModel) (int, error) {
	return s.repo.Create(sb, qs)
}

func (s *SubjectService) ConvertCreate(sb QuizAppApi.CreateNewSubjectBody) (QuizAppApi.SubjectModel, []QuizAppApi.QuestionModel) {
	newSubject := QuizAppApi.SubjectModel{
		Name:        sb.Name,
		ColorString: sb.ColorString,
		IconPath:    sb.IconPath,
	}

	var questions []QuizAppApi.QuestionModel
	for i, e := range sb.Questions {
		questions = append(questions, QuizAppApi.QuestionModel{
			QuestionId:    uint(i + 1),
			Text:          e.Text,
			Answer1:       e.Answer1,
			Answer2:       e.Answer2,
			Answer3:       e.Answer3,
			Answer4:       e.Answer4,
			CorrectAnswer: e.CorrectAnswer,
			Level:         e.Level,
		})
	}

	return newSubject, questions
}

func (s *SubjectService) DeleteSubject(f int) error {
	return s.repo.DeleteSubject(f)
}

func (s *SubjectService) GetSubject(sId int) (QuizAppApi.SubjectSingleResponse, error) {
	subject, err := s.repo.GetSubject(sId)
	if err != nil {
		return QuizAppApi.SubjectSingleResponse{}, err
	}

	return QuizAppApi.SubjectSingleResponse{
		prepareSubjectResponse(subject),
		nil,
	}, nil
}

func (s *SubjectService) UpdateSubject(su QuizAppApi.SubjectResponse) error {
	var f QuizAppApi.SubjectModel
	reflectCommonFields(&su, &f)
	return s.repo.UpdateSubject(f)
}
