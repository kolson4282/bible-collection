package memcollection_test

import (
	"testing"

	"github.com/kolson4282/bible-collection/biblecollection"
	"github.com/kolson4282/bible-collection/memcollection"
)

func TestGetCharacters(t *testing.T) {
	t.Run("Get All Characters", func(t *testing.T) {
		var c biblecollection.BibleCollection = memcollection.NewMemoryCollection()
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
		var c biblecollection.BibleCollection = memcollection.NewMemoryCollection()
		char, err := c.GetCharacterByID(1)
		if err != nil {
			t.Fatalf("Character not found")
		}
		if char.ID != 1 {
			t.Errorf("Incorrect Character passed back. Wanted 1, got %v", char.ID)
		}
	})

}
