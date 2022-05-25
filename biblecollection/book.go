package biblecollection

type Book struct {
	ID       string
	Name     string
	Chapters int
}

type BookCollection interface {
	GetAllBooks() []*Book
	GetBookByID(bookID string) (*Book, error)
	GetBookByName(bookName string) []*Book
}

type Chapter struct {
	Book    *Book `json:"book"`
	Chapter int   `json:"chapter"`
}
