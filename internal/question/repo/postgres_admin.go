package repo

import (
	"fmt"

	questionEntity "github.com/ghulammuzz/misterblast/internal/question/entity"
	"github.com/ghulammuzz/misterblast/pkg/app"
	"github.com/ghulammuzz/misterblast/pkg/log"
)

func (r *questionRepository) ListAdmin(filter map[string]string) ([]questionEntity.ListQuestionAdmin, error) {
	query := `
		SELECT q.id, q.number, q.type, q.content, q.is_quiz, q.set_id,
			   s.name AS set_name, l.name AS lesson_name, c.name AS class_name
		FROM questions q
		JOIN sets s ON q.set_id = s.id
		JOIN lessons l ON s.lesson_id = l.id
		JOIN classes c ON s.class_id = c.id
		WHERE 1=1
	`

	args := []interface{}{}
	argCounter := 1

	// Apply filters dynamically
	if isQuiz, exists := filter["is_quiz"]; exists {
		query += fmt.Sprintf(" AND q.is_quiz = $%d", argCounter)
		args = append(args, isQuiz)
		argCounter++
	}
	if lesson, exists := filter["lesson"]; exists {
		query += fmt.Sprintf(" AND l.name = $%d", argCounter)
		args = append(args, lesson)
		argCounter++
	}
	if class, exists := filter["class"]; exists {
		query += fmt.Sprintf(" AND c.name = $%d", argCounter)
		args = append(args, class)
		argCounter++
	}
	if set, exists := filter["set"]; exists {
		query += fmt.Sprintf(" AND s.name = $%d", argCounter)
		args = append(args, set)
		argCounter++
	}

	// Order by number
	query += " ORDER BY q.number"

	rows, err := r.db.Query(query, args...)
	if err != nil {
		log.Error("[Repo][ListAdmin] Error Query:", err)
		return nil, app.NewAppError(500, "failed to fetch admin questions")
	}
	defer rows.Close()

	var questions []questionEntity.ListQuestionAdmin

	for rows.Next() {
		var q questionEntity.ListQuestionAdmin
		err := rows.Scan(&q.ID, &q.Number, &q.Type, &q.Content, &q.IsQuiz, &q.SetID, &q.SetName, &q.LessonName, &q.ClassName)
		if err != nil {
			log.Error("[Repo][ListAdmin] Error Scan:", err)
			return nil, app.NewAppError(500, "failed to scan admin questions")
		}
		questions = append(questions, q)
	}

	return questions, nil
}
