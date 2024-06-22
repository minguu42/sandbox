package services

import (
	"context"

	"github.com/minguu42/sandbox/example-graphql/graph/db"
	"github.com/minguu42/sandbox/example-graphql/graph/model"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func convertUser(user *db.User) *model.User {
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
}

func convertUserSlice(users db.UserSlice) []*model.User {
	result := make([]*model.User, 0, len(users))
	for _, user := range users {
		result = append(result, convertUser(user))
	}
	return result
}

type UserService interface {
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	ListUsersByID(ctx context.Context, ids []string) ([]*model.User, error)
}

type userService struct {
	exec boil.ContextExecutor
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user, err := db.Users(
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name),
		db.UserWhere.ID.EQ(id),
	).One(ctx, s.exec)
	if err != nil {
		return nil, err
	}
	return convertUser(user), nil
}

func (s *userService) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	user, err := db.Users(
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name),
		db.UserWhere.Name.EQ(name),
	).One(ctx, s.exec)
	if err != nil {
		return nil, err
	}
	return convertUser(user), nil
}

func (s *userService) ListUsersByID(ctx context.Context, ids []string) ([]*model.User, error) {
	users, err := db.Users(
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Name),
		db.UserWhere.ID.IN(ids),
	).All(ctx, s.exec)
	if err != nil {
		return nil, err
	}
	return convertUserSlice(users), nil
}
