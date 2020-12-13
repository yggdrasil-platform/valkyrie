package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	_graphql "github.com/kieranroneill/valkyrie/pkg/graphql"
	"github.com/kieranroneill/valkyrie/pkg/model"
	"github.com/kieranroneill/valkyrie/pkg/service"
)

func (r *queryResolver) GetApplication(ctx context.Context) ([]*model.Application, error) {
	srv := service.NewApplicationService(r.Database)

	return srv.Get(), nil
}

func (r *queryResolver) GetApplicationByAlias(ctx context.Context, alias string) (*model.Application, error) {
	srv := service.NewApplicationService(r.Database)

	return srv.GetByAlias(alias), nil
}

func (r *queryResolver) GetApplicationByID(ctx context.Context, id int) (*model.Application, error) {
	srv := service.NewApplicationService(r.Database)

	return srv.GetById(id), nil
}

// Query returns _graphql.QueryResolver implementation.
func (r *Resolver) Query() _graphql.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
