package memcollection

import (
	"fmt"
	"strings"

	"github.com/kolson4282/bible-collection/biblecollection"
)

func (mc *MemoryCollection) GetAllBooks() []*biblecollection.Book {
	return mc.books
}

func (mc *MemoryCollection) GetBookByID(bookID string) (*biblecollection.Book, error) {
	for _, book := range mc.books {
		if strings.EqualFold(book.ID, bookID) {
			return book, nil
		}
	}
	return nil, fmt.Errorf("no book found with id %s", bookID)
}

func (mc *MemoryCollection) GetBookByName(bookName string) []*biblecollection.Book {
	var books []*biblecollection.Book
	for _, book := range mc.books {
		if strings.Contains(strings.ToLower(book.Name), bookName) {
			books = append(books, book)
		}
	}
	return books
}
