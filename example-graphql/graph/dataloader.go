package graph

import (
	"context"
	"errors"

	"github.com/graph-gophers/dataloader/v7"
	"github.com/minguu42/sandbox/example-graphql/graph/model"
	"github.com/minguu42/sandbox/example-graphql/graph/services"
)

type Loaders struct {
	UserLoader dataloader.Interface[string, *model.User]
}

func NewLoaders(Srv services.Services) *Loaders {
	ub := &userBatcher{Srv: Srv}
	return &Loaders{
		UserLoader: dataloader.NewBatchedLoader[string, *model.User](ub.BatchGetUsers),
	}
}

type userBatcher struct {
	Srv services.Services
}

func (b *userBatcher) BatchGetUsers(ctx context.Context, ids []string) []*dataloader.Result[*model.User] {
	results := make([]*dataloader.Result[*model.User], len(ids))
	for i := range results {
		results[i] = &dataloader.Result[*model.User]{
			Error: errors.New("not found"),
		}
	}

	// 検索条件であるIDが引数でもらったidsスライスの何番目のインデックスに格納されていたのかを検索できるようにマップ化する
	indexes := make(map[string]int, len(ids))
	for i, id := range ids {
		indexes[id] = i
	}

	users, err := b.Srv.ListUsersByID(ctx, ids)

	// 取得結果をresultsの適切な位置に配置する
	for _, user := range users {
		var rsl *dataloader.Result[*model.User]
		if err != nil {
			rsl = &dataloader.Result[*model.User]{
				Error: err,
			}
		} else {
			rsl = &dataloader.Result[*model.User]{
				Data: user,
			}
		}
		results[indexes[user.ID]] = rsl
	}
	return results
}
