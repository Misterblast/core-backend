// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"database/sql"
	"github.com/ghulammuzz/misterblast/internal/lesson/handler"
	"github.com/ghulammuzz/misterblast/internal/lesson/repo"
	"github.com/ghulammuzz/misterblast/internal/lesson/svc"
	"github.com/go-playground/validator/v10"
)

// Injectors from wire.go:

func InitializedLessonService(sb *sql.DB, val *validator.Validate) *handler.LessonHandler {
	lessonRepository := repo.NewLessonRepository(sb)
	lessonService := svc.NewLessonService(lessonRepository)
	lessonHandler := handler.NewLessonHandler(lessonService, val)
	return lessonHandler
}
