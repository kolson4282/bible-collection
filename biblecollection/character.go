package biblecollection

import (
	"fmt"
	"io"
	"strconv"
)

type Character struct {
	ID          string
	Name        string
	Description string
	Gender      string
}

type CharacterCollection interface {
	GetAllCharacters() []*Character
	GetCharacterByID(charID string) (*Character, error)
	GetCharacterByName(charName string) []*Character
	GetChaptersForCharacter(charID string) ([]*Chapter, error)
}

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderNone   Gender = "none"
)

var AllGender = []Gender{
	GenderMale,
	GenderFemale,
	GenderNone,
}

func (e Gender) IsValid() bool {
	switch e {
	case GenderMale, GenderFemale, GenderNone:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
