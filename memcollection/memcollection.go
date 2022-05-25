package memcollection

import (
	"fmt"
	"strings"

	"github.com/kolson4282/bible-collection/biblecollection"
)

type MemoryCollection struct {
	characters []*biblecollection.Character
}

func NewMemoryCollection() *MemoryCollection {
	return &MemoryCollection{
		characters: []*biblecollection.Character{
			{
				ID:          1,
				Name:        "God",
				Description: "God",
				Gender:      "none",
			},
			{
				ID:          2,
				Name:        "Adam",
				Description: "First Man",
				Gender:      "male",
			},
			{
				ID:          3,
				Name:        "Eve",
				Description: "First Woman",
				Gender:      "female",
			},
		},
	}
}

func (mc *MemoryCollection) GetAllCharacters() []*biblecollection.Character {
	return mc.characters
}

func (mc *MemoryCollection) GetCharacterByID(charID int) (*biblecollection.Character, error) {
	for _, char := range mc.characters {
		if char.ID == charID {
			return char, nil
		}
	}
	return nil, fmt.Errorf("character not found with id %d", charID)
}

func (mc *MemoryCollection) GetCharacterByName(charName string) []*biblecollection.Character {
	var characters []*biblecollection.Character
	for _, char := range mc.characters {
		if strings.Contains(strings.ToLower(char.Name), strings.ToLower(charName)) {
			characters = append(characters, char)
		}
	}
	return characters
}
