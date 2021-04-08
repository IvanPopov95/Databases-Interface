package psqldb

import (
	"database/sql"
	"log"
	"projectttt/models"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var testUser = models.User{
	ID:   1,
	Name: "test",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("Cant open a stub database")
	}

	return db, mock
}

func TestGetUserWithID(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	query := "SELECT * FROM users WHERE id = $1"

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(testUser.ID, testUser.Name)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(testUser.ID).WillReturnRows(rows)
	user, err := GetUserWithID(db, testUser.ID)
	assert.NotEmpty(t, user)
	assert.NoError(t, err)
}

func TestGetUserWithIDError(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	query := "SELECT * FROM users WHERE id = $1"

	rows := sqlmock.NewRows([]string{"id", "name"})

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(testUser.ID).WillReturnRows(rows)

	user, err := GetUserWithID(db, testUser.ID)
	assert.Empty(t, user)
	assert.Error(t, err)
}

func TestGetUsersList(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	query := "SELECT id, name FROM users"

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(testUser.ID, testUser.Name)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	users, err := GetUsersList(db)
	assert.NotEmpty(t, users)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
}

func TestDelete(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	query := "DELETE FROM users WHERE id=\\$1"

	mock.ExpectExec(query).WithArgs(testUser.ID).WillReturnResult(sqlmock.NewResult(0, 1))

	err := DeleteUser(db, testUser.ID)
	assert.NoError(t, err)
}

func TestAddUser(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	query := "INSERT INTO users(name) values($1)"

	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(testUser.Name).WillReturnResult(sqlmock.NewResult(0, 1))
	err := AddUser(db, testUser)
	assert.NoError(t, err)
}
