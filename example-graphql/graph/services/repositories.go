package services

import (
	"context"

	"github.com/minguu42/sandbox/example-graphql/graph/db"
	"github.com/minguu42/sandbox/example-graphql/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func convertRepository(r *db.Repository) *model.Repository {
	return &model.Repository{
		ID: r.ID,
		Owner: &model.User{
			ID: r.Owner,
		},
		Name:      r.Name,
		CreatedAt: r.CreatedAt,
	}
}

type RepositoryService interface {
	GetRepositoryByNameAndOwner(ctx context.Context, name string, owner string) (*model.Repository, error)
}

type repositoryService struct {
	exec boil.ContextExecutor
}

func (s repositoryService) GetRepositoryByNameAndOwner(ctx context.Context, name string, owner string) (*model.Repository, error) {
	r, err := db.Repositories(
		qm.Select(db.RepositoryTableColumns.ID, db.RepositoryColumns.Name, db.RepositoryColumns.Owner, db.RepositoryColumns.CreatedAt),
		db.RepositoryWhere.Name.EQ(name),
		db.RepositoryWhere.Owner.EQ(owner),
	).One(ctx, s.exec)
	if err != nil {
		return nil, err
	}
	return convertRepository(r), nil
}
