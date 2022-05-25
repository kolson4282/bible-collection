package biblecollection

type Book struct {
	ID       string
	Name     string
	Chapters int
}

type BookCollection interface {
	GetAllBooks() []*Book
	GetBookByID(bookID string) (*Book, error)
}
