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

func (mc *MemoryCollection) GetEventsForCharacter(charID string) ([]*biblecollection.Event, error) {
	var events []*biblecollection.Event
	for _, event := range mc.GetAllEvents() {
		if event.ActorID == charID || event.ReceiverID == charID {
			events = append(events, event)
		}
	}
	return events, nil
}

func (mc *MemoryCollection) GetChaptersForCharacter(charID string) ([]*biblecollection.Chapter, error) {
	var chapters []*biblecollection.Chapter
	events, _ := mc.GetEventsForCharacter(charID)
	for _, event := range events {
		book, _ := mc.GetBookByID(event.BookID)
		chapter := event.Chapter
		chapters = mc.appendChapterIfDoesntExist(chapters, &biblecollection.Chapter{
			Book:    book,
			Chapter: chapter,
		})
	}
	return chapters, nil
}

func (mc *MemoryCollection) appendChapterIfDoesntExist(chapters []*biblecollection.Chapter, chapter *biblecollection.Chapter) []*biblecollection.Chapter {
	for _, chap := range chapters {
		if chap.Book.ID == chapter.Book.ID && chap.Chapter == chapter.Chapter {
			return chapters
		}
	}
	return append(chapters, chapter)
}
