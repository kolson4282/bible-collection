package memcollection_test

import (
	"testing"

	"github.com/kolson4282/bible-collection/biblecollection"
	"github.com/kolson4282/bible-collection/memcollection"
)

func TestBook(t *testing.T) {
	var c biblecollection.BibleCollection = memcollection.NewMemoryCollection()
	t.Run("Get All Books", func(t *testing.T) {
		books := c.GetAllBooks()
		if len(books) == 0 {
			t.Fatal("No Books Returned")
		}
	})
	t.Run("Get book By ID", func(t *testing.T) {
		book, err := c.GetBookByID("gen")
		checkForError(t, err)
		if book.ID != "gen" {
			t.Errorf("Incorrect book passed back. Wanted gen, got %v", book.ID)
		}
	})

	t.Run("Error if Character Not Found", func(t *testing.T) {
		_, err := c.GetBookByID("Doesn't Exist")
		checkForNoError(t, err)
	})
}
