package svc_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ghulammuzz/misterblast/internal/class/entity"
	"github.com/ghulammuzz/misterblast/internal/class/repo"
	"github.com/ghulammuzz/misterblast/internal/class/svc"
	"github.com/stretchr/testify/assert"
)

func TestAddClass(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repository := repo.NewClassRepository(mockDB)
	service := svc.NewClassService(repository)

	validClass := entity.SetClass{Name: "1"}
	invalidClass := entity.SetClass{Name: "Invalid"}

	mock.ExpectExec("INSERT INTO classes").
		WithArgs(validClass.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = service.AddClass(validClass)
	assert.NoError(t, err)

	err = service.AddClass(invalidClass)
	assert.Error(t, err)
	assert.Equal(t, "validation failed", err.Error())

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteClass(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repository := repo.NewClassRepository(mockDB)
	service := svc.NewClassService(repository)

	mock.ExpectExec("DELETE FROM classes WHERE id = ").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = service.DeleteClass(1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestListClass(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	repository := repo.NewClassRepository(mockDB)
	service := svc.NewClassService(repository)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "1").
		AddRow(2, "2")

	mock.ExpectQuery("SELECT id, name FROM classes").WillReturnRows(rows)

	classes, err := service.ListClasses()
	assert.NoError(t, err)
	assert.Len(t, classes, 2)
	assert.Equal(t, "1", classes[0].Name)
	assert.Equal(t, "2", classes[1].Name)
	assert.NoError(t, mock.ExpectationsWereMet())
}
