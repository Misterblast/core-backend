// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"database/sql"
	"github.com/ghulammuzz/misterblast/internal/class/handler"
	"github.com/ghulammuzz/misterblast/internal/class/repo"
	"github.com/ghulammuzz/misterblast/internal/class/svc"
)

// Injectors from wire.go:

func InitializedClassService(sb *sql.DB) *handler.ClassHandler {
	classRepository := repo.NewClassRepository(sb)
	classService := svc.NewClassService(classRepository)
	classHandler := handler.NewClassHandler(classService)
	return classHandler
}
