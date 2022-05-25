package memcollection

type MemoryCollection struct {
	characters []string
}

func NewMemoryCollection() *MemoryCollection {
	return &MemoryCollection{
		characters: []string{
			"God",
			"Adam",
			"Eve",
		},
	}
}

func (mc *MemoryCollection) GetAllCharacters() []string {
	return mc.characters
}
