package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"example_crud_graphql/graph/model"
)

func (r *queryResolver) InsertJenisBarang(ctx context.Context, jenisBarang string) (*model.ResultJenisBarang, error) {
	resultJenisBarang := model.ResultJenisBarang{
		Status: "",
		Code:   404,
	}

	return &resultJenisBarang, nil
}
