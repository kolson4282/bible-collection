package memcollection

import (
	"fmt"
	"strings"

	"github.com/kolson4282/bible-collection/biblecollection"
)

func (mc *MemoryCollection) GetAllCharacters() []*biblecollection.Character {
	return mc.characters
}

func (mc *MemoryCollection) GetCharacterByID(charID string) (*biblecollection.Character, error) {
	for _, char := range mc.characters {
		if char.ID == charID {
			return char, nil
		}
	}
	return nil, fmt.Errorf("character not found with id %s", charID)
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

func (mc *MemoryCollection) GetChaptersForCharacter(charID string) ([]*biblecollection.Chapter, error) {
	return []*biblecollection.Chapter{}, nil
}
