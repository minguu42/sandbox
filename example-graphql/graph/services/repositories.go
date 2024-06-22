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
	GetRepoByID(ctx context.Context, id string) (*model.Repository, error)
	GetRepoByFullName(ctx context.Context, owner, name string) (*model.Repository, error)
}

type repoService struct {
	exec boil.ContextExecutor
}

func (r *repoService) GetRepoByID(ctx context.Context, id string) (*model.Repository, error) {
	repo, err := db.FindRepository(ctx, r.exec, id,
		db.RepositoryColumns.ID, db.RepositoryColumns.Name, db.RepositoryColumns.Owner, db.RepositoryColumns.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return convertRepository(repo), nil
}

func (s *repoService) GetRepoByFullName(ctx context.Context, owner, name string) (*model.Repository, error) {
	repo, err := db.Repositories(
		qm.Select(
			db.RepositoryColumns.ID,        // レポジトリID
			db.RepositoryColumns.Name,      // レポジトリ名
			db.RepositoryColumns.Owner,     // レポジトリを所有しているユーザーのID
			db.RepositoryColumns.CreatedAt, // 作成日時
		),
		db.RepositoryWhere.Owner.EQ(owner),
		db.RepositoryWhere.Name.EQ(name),
	).One(ctx, s.exec)
	if err != nil {
		return nil, err
	}
	return convertRepository(repo), nil
}
