package entity

type SetAnswer struct {
	ID         int32   `json:"id"`
	QuestionID int32   `json:"question_id" validate:"required"`
	Code       string  `json:"code" validate:"required,oneof=a b c d esay"`
	Content    string  `json:"content" validate:"required"`
	ImgURL     *string `json:"img_url,omitempty"`
	IsAnswer   bool    `json:"is_answer"`
}

type ListAnswer struct {
	ID      int32   `json:"id"`
	Code    string  `json:"code"`
	Content string  `json:"content"`
	ImgURL  *string `json:"img_url"`
}
