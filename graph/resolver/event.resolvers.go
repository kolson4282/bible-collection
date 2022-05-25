package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kolson4282/bible-collection/biblecollection"
	"github.com/kolson4282/bible-collection/graph/generated"
)

func (r *eventResolver) Actor(ctx context.Context, obj *biblecollection.Event) (*biblecollection.Character, error) {
	return r.Collection.GetCharacterByID(obj.ActorID)
}

func (r *eventResolver) Receiver(ctx context.Context, obj *biblecollection.Event) (*biblecollection.Character, error) {
	return r.Collection.GetCharacterByID(obj.ReceiverID)
}

func (r *eventResolver) Book(ctx context.Context, obj *biblecollection.Event) (*biblecollection.Book, error) {
	return r.Collection.GetBookByID(obj.BookID)
}

// Event returns generated.EventResolver implementation.
func (r *Resolver) Event() generated.EventResolver { return &eventResolver{r} }

type eventResolver struct{ *Resolver }
