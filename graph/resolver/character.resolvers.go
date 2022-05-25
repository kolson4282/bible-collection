package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kolson4282/bible-collection/biblecollection"
	"github.com/kolson4282/bible-collection/graph/generated"
)

func (r *characterResolver) Gender(ctx context.Context, obj *biblecollection.Character) (biblecollection.Gender, error) {
	if biblecollection.Gender(obj.Gender).IsValid() {
		return biblecollection.GenderNone, nil
	}
	return biblecollection.Gender(obj.Gender), nil
}

func (r *characterResolver) Events(ctx context.Context, obj *biblecollection.Character) ([]*biblecollection.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *characterResolver) Chapters(ctx context.Context, obj *biblecollection.Character) ([]*biblecollection.Chapter, error) {
	return r.Collection.GetChaptersForCharacter(obj.ID)
}

// Character returns generated.CharacterResolver implementation.
func (r *Resolver) Character() generated.CharacterResolver { return &characterResolver{r} }

type characterResolver struct{ *Resolver }
