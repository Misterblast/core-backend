package repo

import (
	"database/sql"
	"fmt"

	questionEntity "github.com/ghulammuzz/misterblast/internal/question/entity"
	"github.com/ghulammuzz/misterblast/pkg/app"
	"github.com/ghulammuzz/misterblast/pkg/log"
)

type QuestionRepository interface {
	Add(question questionEntity.SetQuestion) error
	List(filter map[string]string) ([]questionEntity.ListQuestionExample, error)
	Delete(id int32) error
	Exists(setID int32, number int) (bool, error)
	Edit(id int32, question questionEntity.EditQuestion) error
	AddQuizAnswer(answer questionEntity.SetAnswer) error
	ListQuizQuestions(filter map[string]string) ([]questionEntity.ListQuestionQuiz, error)
	ListAdmin(filter map[string]string, page, limit int) ([]questionEntity.ListQuestionAdmin, error)
	DeleteAnswer(id int32) error

}

func (r *questionRepository) DeleteAnswer(id int32) error {
	query := `DELETE FROM answers WHERE id = $1`
	res, err := r.db.Exec(query, id)
	if err != nil {
		log.Error("[Repo][DeleteAnswer] Error deleting answer:", err)
		return app.NewAppError(500, "failed to delete answer")
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return app.NewAppError(404, "answer not found")
	}

	return nil
}

type questionRepository struct {
	db *sql.DB
}

func (r *questionRepository) AddQuizAnswer(answer questionEntity.SetAnswer) error {
	query := `
		INSERT INTO answers (question_id, code, content, img_url, is_answer) 
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := r.db.Exec(query, answer.QuestionID, answer.Code, answer.Content, answer.ImgURL, answer.IsAnswer)
	if err != nil {
		log.Error("[Repo][AddQuizAnswer] Error inserting answer: ", err)
		return app.NewAppError(500, "failed to insert quiz answer")
	}

	return nil
}

func NewQuestionRepository(db *sql.DB) QuestionRepository {
	return &questionRepository{db: db}
}

func (r *questionRepository) Add(question questionEntity.SetQuestion) error {
	query := `INSERT INTO questions (number, type, content, is_quiz, set_id) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, question.Number, question.Type, question.Content, question.IsQuiz, question.SetID)
	if err != nil {
		log.Error("[Repo][AddQuestion] Error inserting question:", err)
		return app.NewAppError(500, err.Error())
	}
	return nil
}

func (r *questionRepository) List(filter map[string]string) ([]questionEntity.ListQuestionExample, error) {
	query := `SELECT id, number, type, content, set_id FROM questions WHERE 1=1`
	args := []interface{}{}
	argCounter := 1

	if setID, ok := filter["set_id"]; ok {
		query += fmt.Sprintf(" AND set_id = $%d", argCounter)
		args = append(args, setID)
		argCounter++
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		log.Error("[Repo][ListQuestions] Error Query: ", err)
		return nil, app.NewAppError(500, err.Error())
	}
	defer rows.Close()

	var questions []questionEntity.ListQuestionExample
	for rows.Next() {
		var question questionEntity.ListQuestionExample
		if err := rows.Scan(&question.ID, &question.Number, &question.Type, &question.Content, &question.SetID); err != nil {
			log.Error("[Repo][ListQuestions] Error Scan: ", err)
			return nil, app.NewAppError(500, "failed to scan question")
		}
		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		log.Error("[Repo][ListQuestions] Error Iterating Rows: ", err)
		return nil, app.NewAppError(500, "error iterating rows")
	}

	return questions, nil
}

func (r *questionRepository) Delete(id int32) error {
	query := `DELETE FROM questions WHERE id = $1`
	res, err := r.db.Exec(query, id)
	if err != nil {
		log.Error("[Repo][DeleteQuestion] Error deleting question:", err)
		return app.NewAppError(500, "failed to delete question")
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return app.NewAppError(404, "question not found")
	}

	return nil
}

func (r *questionRepository) Exists(setID int32, number int) (bool, error) {
	query := `SELECT COUNT(*) FROM questions WHERE set_id = $1 AND number = $2`
	var count int
	err := r.db.QueryRow(query, setID, number).Scan(&count)
	if err != nil {
		log.Error("[Repo][ExistsQuestion] Error checking question:", err)
		return false, app.NewAppError(500, "failed to check question existence")
	}

	return count > 0, nil
}

func (r *questionRepository) ListQuizQuestions(filter map[string]string) ([]questionEntity.ListQuestionQuiz, error) {
	query := `
		SELECT q.id, q.number, q.type, q.content, q.set_id,
			   COALESCE(a.id, 0) AS answer_id, COALESCE(a.code, '') AS code, 
			   COALESCE(a.content, '') AS answer_content, COALESCE(a.img_url, '') AS img_url
		FROM questions q
		LEFT JOIN answers a ON q.id = a.question_id
		WHERE q.is_quiz = true
	`
	args := []interface{}{}
	argCounter := 1

	if setID, exists := filter["set_id"]; exists {
		query += fmt.Sprintf(" AND q.set_id = $%d", argCounter)
		args = append(args, setID)
		argCounter++
	}
	if questionType, exists := filter["type"]; exists {
		query += fmt.Sprintf(" AND q.type = $%d", argCounter)
		args = append(args, questionType)
		argCounter++
	}
	if number, exists := filter["number"]; exists {
		query += fmt.Sprintf(" AND q.number = $%d", argCounter)
		args = append(args, number)
		argCounter++
	}

	query += " ORDER BY q.number, a.code"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		log.Error("[Repo][ListQuizQuestions] Error Query: ", err)
		return nil, app.NewAppError(500, "failed to fetch quiz questions")
	}
	defer rows.Close()

	questionsMap := make(map[int32]*questionEntity.ListQuestionQuiz)
	var questions []*questionEntity.ListQuestionQuiz

	for rows.Next() {
		var qID int32
		var number int
		var qType, content string
		var setID int32
		var aID int32
		var code, aContent string
		var imgURL string

		err := rows.Scan(&qID, &number, &qType, &content, &setID,
			&aID, &code, &aContent, &imgURL)
		if err != nil {
			log.Error("[Repo][ListQuizQuestions] Error Scan: ", err)
			return nil, app.NewAppError(500, "failed to scan quiz questions")
		}

		if _, exists := questionsMap[qID]; !exists {
			questionsMap[qID] = &questionEntity.ListQuestionQuiz{
				ID:      qID,
				Number:  number,
				Type:    qType,
				Content: content,
				SetID:   setID,
				Answers: []questionEntity.ListAnswer{},
			}
			questions = append(questions, questionsMap[qID])
		}

		if aID != 0 {
			answer := questionEntity.ListAnswer{
				ID:      aID,
				Code:    code,
				Content: aContent,
			}
			if imgURL != "" {
				answer.ImgURL = &imgURL
			}
			questionsMap[qID].Answers = append(questionsMap[qID].Answers, answer)
		}
	}

	finalQuestions := make([]questionEntity.ListQuestionQuiz, len(questions))
	for i, q := range questions {
		finalQuestions[i] = *q
	}

	return finalQuestions, nil
}

func (r *questionRepository) Edit(id int32, question questionEntity.EditQuestion) error {
	query := `
		UPDATE questions 
		SET number = $1, type = $2, content = $3, is_quiz = $4, set_id = $5 
		WHERE id = $6`
	
	_, err := r.db.Exec(query, question.Number, question.Type, question.Content, question.IsQuiz, question.SetID, id)
	if err != nil {
		log.Error("[Repo][EditQuestion] Error updating question:", err)
		return app.NewAppError(500, err.Error())
	}

	return nil
}