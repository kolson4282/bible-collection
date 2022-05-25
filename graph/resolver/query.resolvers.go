package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kolson4282/bible-collection/biblecollection"
	"github.com/kolson4282/bible-collection/graph/generated"
)

func (r *queryResolver) Characters(ctx context.Context) ([]*biblecollection.Character, error) {
	return r.Collection.GetAllCharacters(), nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*biblecollection.Book, error) {
	return r.Collection.GetAllBooks(), nil
}

func (r *queryResolver) Events(ctx context.Context) ([]*biblecollection.Event, error) {
	return r.Collection.GetAllEvents(), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
