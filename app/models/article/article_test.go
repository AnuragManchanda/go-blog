package article

import (
	"fmt"
	dbClient "github.com/AnuragManchanda/go-blog/db"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

// a successful case
func TestAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectBegin()
	mock.ExpectExec("SELECT id, title, author, content FROM articles").WithArgs().WillReturnResult(sqlmock.NewResult(0,0))
	mock.ExpectCommit()

	dbClient.SetTestDb(db)

	// now we execute our method
	result := All();
	fmt.Print(result)
}
