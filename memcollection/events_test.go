package memcollection_test

import (
	"testing"

	"github.com/kolson4282/bible-collection/biblecollection"
	"github.com/kolson4282/bible-collection/memcollection"
)

func TestEvents(t *testing.T) {
	var c biblecollection.BibleCollection = memcollection.NewMemoryCollection()

	t.Run("Get All Events", func(t *testing.T) {
		events := c.GetAllEvents()
		if len(events) == 0 {
			t.Error("No events returned")
		}
	})
}
