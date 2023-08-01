package QuizAppApi

type SubjectModel struct {
	ID          uint `gorm:"primarykey"`
	Name        string
	IconPath    string
	ColorString string
}

type SubjectResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	IconPath    string `json:"iconPath"`
	ColorString string `json:"color"`
}

type SubjectSingleResponse struct {
	SubjectResponse
	Questions []QuestionResponse `json:"questions"`
}

type QuestionModel struct {
	ID            uint `gorm:"primarykey"`
	SubjectId     uint
	QuestionId    uint
	Text          string
	Answer1       string
	Answer2       string
	Answer3       string
	Answer4       string
	CorrectAnswer string
	Level         int
}

type QuestionResponse struct {
	SubjectId     uint   `json:"sid"`
	QuestionId    uint   `json:"qid"`
	Text          string `json:"text"`
	Answer1       string `json:"answer1"`
	Answer2       string `json:"answer2"`
	Answer3       string `json:"answer3"`
	Answer4       string `json:"answer4"`
	CorrectAnswer string `json:"correctAnswer"`
	Level         int    `json:"level"`
}

// 0 - easy
// 1 - medium
// 2 - hard

type createNewSubjectQuestionsBody struct {
	Text          string `json:"text" binding:"required"`
	Answer1       string `json:"a1" binding:"required"`
	Answer2       string `json:"a2" binding:"required"`
	Answer3       string `json:"a3" binding:"required"`
	Answer4       string `json:"a4" binding:"required"`
	CorrectAnswer string `json:"ca" binding:"required"`
	Level         int    `json:"level" binding:"required"`
}
type CreateNewSubjectBody struct {
	Name        string                          `json:"name" binding:"required"`
	ColorString string                          `json:"color" binding:"required"`
	IconPath    string                          `json:"iconPath"`
	Questions   []createNewSubjectQuestionsBody `json:"questions" binding:"required"`
}

type CreateQuestionBody struct {
	SubjectId     uint   `json:"sid" binding:"true"`
	Text          string `json:"text" binding:"true"`
	Answer1       string `json:"a1" binding:"true"`
	Answer2       string `json:"a2" binding:"true"`
	Answer3       string `json:"a3" binding:"true"`
	Answer4       string `json:"a4" binding:"true"`
	CorrectAnswer string `json:"ca" binding:"true"`
	Level         int    `json:"level" binding:"true"`
}
