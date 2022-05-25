package biblecollection

type Character struct {
	ID          int
	Name        string
	Description string
	Gender      string
}

type CharacterCollection interface {
	GetAllCharacters() []*Character
	GetCharacterByID(charID int) (*Character, error)
	GetCharacterByName(charName string) []*Character
}
