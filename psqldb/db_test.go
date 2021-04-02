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

var testItem = models.Item{
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

func TestGetItemWithID(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	query := "SELECT * FROM items WHERE id = $1"

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(testItem.ID, testItem.Name)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(testItem.ID).WillReturnRows(rows)
	item, err := GetItemWithID(db, testItem.ID)
	assert.NotEmpty(t, item)
	assert.NoError(t, err)
}

func TestGetItemWithIDError(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()

	query := "SELECT * FROM items WHERE id = $1"

	rows := sqlmock.NewRows([]string{"id", "name"})

	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(testItem.ID).WillReturnRows(rows)

	item, err := GetItemWithID(db, testItem.ID)
	assert.Empty(t, item)
	assert.Error(t, err)
}

func TestGetItemsList(t *testing.T) {
	db, mock := NewMock()
	defer db.Close()
	query := "SELECT id, name FROM items"

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(testItem.ID, testItem.Name)

	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	items, err := GetItemsList(db)
	assert.NotEmpty(t, items)
	assert.NoError(t, err)
	assert.Len(t, items, 1)
}

// func TestAddItem(t *testing.T) {
// 	db, mock := NewMock()
// 	defer db.Close()
// 	query := "INSERT INTO items(name) values($1)"

// 	prep := mock.ExpectPrepare(regexp.QuoteMeta(query))
// 	res := prep.ExpectExec().WithArgs(testItem.Name).WillReturnResult(sqlmock.NewResult(1, 1))
// 	fmt.Println(res)
// 	err := AddItem(db, testItem)
// 	assert.NoError(t, err)

// }
