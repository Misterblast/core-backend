package entity

type SetQuestion struct {
	ID      int32  `json:"id"`
	Number  int    `json:"number" validate:"required,min=1"`
	Type    string `json:"type" validate:"required,oneof=C1 C2 C3 C4 C5 C6"`
	Content string `json:"content" validate:"required"`
	IsQuiz  bool   `json:"is_quiz"`
	SetID   int32  `json:"set_id" validate:"required"`
}

type EditQuestion struct {
	Number  int    `json:"number" validate:"required,min=1"`
	Type    string `json:"type" validate:"required,oneof=C1 C2 C3 C4 C5 C6"`
	Content string `json:"content" validate:"required"`
	IsQuiz  bool   `json:"is_quiz"`
	SetID   int32  `json:"set_id" validate:"required"`
}

type ListQuestionExample struct {
	ID      int32  `json:"id"`
	Number  int    `json:"number"`
	Type    string `json:"type"`
	Content string `json:"content"`
	SetID   int32  `json:"set_id"`
}

type ListQuestionQuiz struct {
	ID      int32        `json:"id"`
	Number  int          `json:"number"`
	Type    string       `json:"type"`
	Content string       `json:"content"`
	SetID   int32        `json:"set_id"`
	Answers []ListAnswer `json:"answers"`
}

type ListQuestionAdmin struct {
	ID         int32  `json:"id"`
	Number     int    `json:"number"`
	Type       string `json:"type"`
	Content    string `json:"content"`
	IsQuiz     bool   `json:"is_quiz"`
	SetID      int32  `json:"set_id"`
	SetName    string `json:"set_name"`
	LessonName string `json:"lesson_name"`
	ClassName  string `json:"class_name"`
}
