package memcollection

import "github.com/kolson4282/bible-collection/biblecollection"

type MemoryCollection struct {
	characters []biblecollection.Character
}

func NewMemoryCollection() *MemoryCollection {
	return &MemoryCollection{
		characters: []biblecollection.Character{
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

func (mc *MemoryCollection) GetAllCharacters() []biblecollection.Character {
	return mc.characters
}
