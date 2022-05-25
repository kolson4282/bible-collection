package memcollection

import "github.com/kolson4282/bible-collection/biblecollection"

func (mc *MemoryCollection) GetAllEvents() []*biblecollection.Event {
	return mc.events
}
