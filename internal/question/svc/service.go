package svc

import (
	questionEntity "github.com/ghulammuzz/misterblast/internal/question/entity"
	"github.com/ghulammuzz/misterblast/internal/question/repo"
	"github.com/ghulammuzz/misterblast/pkg/app"
)

type QuestionService interface {
	AddQuestion(question questionEntity.SetQuestion) error
	AddQuizAnswer(question questionEntity.SetAnswer) error
	ListQuestions(filter map[string]string) ([]questionEntity.ListQuestionExample, error)
	ListQuizQuestions(filter map[string]string) ([]questionEntity.ListQuestionQuiz, error)
	DeleteQuestion(id int32) error
	DeleteAnswer(id int32) error
	ListAdmin(filter map[string]string, page, limit int) ([]questionEntity.ListQuestionAdmin, error)
	EditQuestion(id int32, question questionEntity.EditQuestion) error 
}

type questionService struct {
	repo repo.QuestionRepository
}

func NewQuestionService(repo repo.QuestionRepository) QuestionService {
	return &questionService{repo: repo}
}
func (s *questionService) AddQuizAnswer(question questionEntity.SetAnswer) error {
	return s.repo.AddQuizAnswer(question)
}

func (s *questionService) AddQuestion(q questionEntity.SetQuestion) error {
	exists, err := s.repo.Exists(q.SetID, q.Number)
	if err != nil {
		return err
	}
	if exists {
		return app.NewAppError(409, "question number already exists in this set")
	}

	return s.repo.Add(q)
}

func (s *questionService) ListQuestions(filter map[string]string) ([]questionEntity.ListQuestionExample, error) {
	return s.repo.List(filter)
}

func (s *questionService) DeleteQuestion(id int32) error {
	return s.repo.Delete(id)
}

// Quiz

func (s *questionService) ListQuizQuestions(filter map[string]string) ([]questionEntity.ListQuestionQuiz, error) {
	return s.repo.ListQuizQuestions(filter)
}

// admin

func (s *questionService) ListAdmin(filter map[string]string, page, limit int) ([]questionEntity.ListQuestionAdmin, error) {
	questions, err := s.repo.ListAdmin(filter, page, limit)
	if err != nil {
		return nil, err
	}
	return questions, nil
}

func (s *questionService) EditQuestion(id int32, question questionEntity.EditQuestion) error {
	return s.repo.Edit(id, question)
}

func (s *questionService) DeleteAnswer(id int32) error {
	return s.repo.DeleteAnswer(id)
}