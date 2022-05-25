package biblecollection

type Character struct {
	ID          string
	Name        string
	Description string
	Gender      string
}

type CharacterCollection interface {
	GetAllCharacters() []*Character
	GetCharacterByID(charID string) (*Character, error)
	GetCharacterByName(charName string) []*Character
}
