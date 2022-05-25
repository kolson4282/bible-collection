package memcollection

import (
	"github.com/kolson4282/bible-collection/biblecollection"
)

type MemoryCollection struct {
	characters []*biblecollection.Character
	books      []*biblecollection.Book
}

func NewMemoryCollection() *MemoryCollection {
	return &MemoryCollection{
		characters: []*biblecollection.Character{
			{
				ID:          "god",
				Name:        "God",
				Description: "God",
				Gender:      "none",
			},
			{
				ID:          "adam-1",
				Name:        "Adam",
				Description: "First Man",
				Gender:      "male",
			},
			{
				ID:          "eve-1",
				Name:        "Eve",
				Description: "First Woman",
				Gender:      "female",
			},
		},
		books: []*biblecollection.Book{
			{
				ID:       "gen",
				Name:     "Genesis",
				Chapters: 50,
			},
			{
				ID:       "exo",
				Name:     "Exodus",
				Chapters: 40,
			},
			{
				ID:       "Lev",
				Name:     "Leviticus",
				Chapters: 27,
			},
		},
	}
}
