package biblecollection

type BibleCollection interface {
	GetAllCharacters() []*Character
	GetCharacterByID(charID int) (*Character, error)
	GetCharacterByName(charName string) []*Character
}

type Character struct {
	ID          int
	Name        string
	Description string
	Gender      string
}
