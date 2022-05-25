package memcollection_test

import (
	"strings"
	"testing"

	"github.com/kolson4282/bible-collection/biblecollection"
	"github.com/kolson4282/bible-collection/memcollection"
)

func TestGetCharacters(t *testing.T) {
	var c biblecollection.BibleCollection = memcollection.NewMemoryCollection()

	t.Run("Get All Characters", func(t *testing.T) {
		chars := c.GetAllCharacters()
		if len(chars) == 0 {
			t.Error("No Characters Returned")
		}
		for i, char := range chars {
			if char.Name == "" {
				t.Errorf("Invalid character at postion %v", i)
			}
			if char.Description == "" {
				t.Errorf("Invalid description at postion %v", i)
			}
			if char.Gender == "" {
				t.Errorf("Invalid gender at postion %v", i)
			}
		}
	})

	t.Run("Get Character By ID", func(t *testing.T) {
		char, err := c.GetCharacterByID("adam-1")
		checkForError(t, err)
		if char.ID != "adam-1" {
			t.Errorf("Incorrect Character passed back. Wanted 'adam-1', got %s", char.ID)
		}
	})

	t.Run("Error if Character Not Found", func(t *testing.T) {
		_, err := c.GetCharacterByID("doesnt-exist")
		checkForNoError(t, err)
	})

	t.Run("Get Character By Name", func(t *testing.T) {
		characters := c.GetCharacterByName("adam")
		for i, char := range characters {
			if !strings.Contains(strings.ToLower(char.Name), "adam") {
				t.Errorf("Invalid customer. Wanted %s, got %s at position %d", "adam", char.Name, i)
			}
		}
	})

	t.Run("Get Character By Partial Name", func(t *testing.T) {
		characters := c.GetCharacterByName("ad")
		for i, char := range characters {
			if !strings.Contains(strings.ToLower(char.Name), "ad") {
				t.Errorf("Invalid customer. Wanted %s, got %s at position %d", "adam", char.Name, i)
			}
		}
	})
}
