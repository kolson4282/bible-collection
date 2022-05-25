package biblecollection

type BibleCollection interface {
	GetAllCharacters() []Character
}

type Character struct {
	ID          int
	Name        string
	Description string
	Gender      string
}
