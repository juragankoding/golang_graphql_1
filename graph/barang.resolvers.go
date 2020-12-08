package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example_crud_graphql/graph/generated"
	"example_crud_graphql/graph/model"
)

func (r *queryResolver) GetBarang(ctx context.Context) (*string, error) {
	keterangan := "belum di implementasikan"

	return &keterangan, nil
}

func (r *queryResolver) InsertBarang(ctx context.Context, id int, nama string, description string) (*model.Barang, error) {

	barang := model.Barang{
		ID:          id,
		Nama:        nama,
		Description: description,
	}

	return &barang, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
