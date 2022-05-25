package main_test

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
	})
}
