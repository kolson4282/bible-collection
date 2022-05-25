package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kolson4282/bible-collection/biblecollection"
	"github.com/kolson4282/bible-collection/graph/generated"
)

func (r *bookResolver) Events(ctx context.Context, obj *biblecollection.Book) ([]*biblecollection.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

// Book returns generated.BookResolver implementation.
func (r *Resolver) Book() generated.BookResolver { return &bookResolver{r} }

type bookResolver struct{ *Resolver }
