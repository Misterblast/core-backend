// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"database/sql"
	"github.com/ghulammuzz/misterblast/internal/email/handler"
	"github.com/ghulammuzz/misterblast/internal/email/repo"
	"github.com/ghulammuzz/misterblast/internal/email/svc"
	repo2 "github.com/ghulammuzz/misterblast/internal/user/repo"
	"github.com/go-playground/validator/v10"
)

// Injectors from wire.go:

func InitializedEmailService(sb *sql.DB, val *validator.Validate) *handler.EmailHandler {
	emailRepository := repo.NewEmailRepository(sb)
	userRepository := repo2.NewUserRepository(sb)
	otp := repo.NewOTPService()
	emailService := svc.NewEmailService(emailRepository, userRepository, otp)
	emailHandler := handler.NewEmailHandler(emailService, val)
	return emailHandler
}
