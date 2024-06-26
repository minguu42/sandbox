package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.49

import (
	"context"
	"fmt"
	"strings"

	"github.com/friendsofgo/errors"
	"github.com/minguu42/sandbox/example-graphql/graph/model"
	"github.com/minguu42/sandbox/example-graphql/internal"
)

// Author is the resolver for the author field.
func (r *issueResolver) Author(ctx context.Context, obj *model.Issue) (*model.User, error) {
	thunk := r.Loaders.UserLoader.Load(ctx, obj.Author.ID)
	user, err := thunk()
	if err != nil {
		return nil, err
	}
	return user, nil
}

// AddProjectV2ItemByID is the resolver for the addProjectV2ItemById field.
func (r *mutationResolver) AddProjectV2ItemByID(ctx context.Context, input model.AddProjectV2ItemByIDInput) (*model.AddProjectV2ItemByIDPayload, error) {
	panic(fmt.Errorf("not implemented: AddProjectV2ItemByID - addProjectV2ItemById"))
}

// Owner is the resolver for the owner field.
func (r *projectV2Resolver) Owner(ctx context.Context, obj *model.ProjectV2) (*model.User, error) {
	return r.Srv.GetUserByID(ctx, obj.Owner.ID)
}

// Repository is the resolver for the repository field.
func (r *queryResolver) Repository(ctx context.Context, name string, owner string) (*model.Repository, error) {
	user, err := r.Srv.GetUserByName(ctx, owner)
	if err != nil {
		return nil, err
	}
	return r.Srv.GetRepoByFullName(ctx, user.ID, name)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, name string) (*model.User, error) {
	return r.Srv.GetUserByName(ctx, name)
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (model.Node, error) {
	nElems := strings.SplitN(id, "_", 2)
	nType, _ := nElems[0], nElems[1]
	switch nType {
	case "REPO":
		return r.Srv.GetRepoByID(ctx, id)
	case "ISSUE":
		return r.Srv.GetIssueByID(ctx, id)
	default:
		return nil, errors.New("invalid node type")
	}
}

// Owner is the resolver for the owner field.
func (r *repositoryResolver) Owner(ctx context.Context, obj *model.Repository) (*model.User, error) {
	return r.Srv.GetUserByID(ctx, obj.Owner.ID)
}

// Issue is the resolver for the issue field.
func (r *repositoryResolver) Issue(ctx context.Context, obj *model.Repository, number int) (*model.Issue, error) {
	return r.Srv.GetIssueByRepoAndNumber(ctx, obj.ID, number)
}

// Issues is the resolver for the issues field.
func (r *repositoryResolver) Issues(ctx context.Context, obj *model.Repository, after *string, before *string, first *int, last *int) (*model.IssueConnection, error) {
	return r.Srv.ListIssueInRepository(ctx, obj.ID, after, before, first, last)
}

// PullRequest is the resolver for the pullRequest field.
func (r *repositoryResolver) PullRequest(ctx context.Context, obj *model.Repository, number int) (*model.PullRequest, error) {
	panic(fmt.Errorf("not implemented: PullRequest - pullRequest"))
}

// PullRequests is the resolver for the pullRequests field.
func (r *repositoryResolver) PullRequests(ctx context.Context, obj *model.Repository, after *string, before *string, first *int, last *int) (*model.PullRequestConnection, error) {
	panic(fmt.Errorf("not implemented: PullRequests - pullRequests"))
}

// Issue returns internal.IssueResolver implementation.
func (r *Resolver) Issue() internal.IssueResolver { return &issueResolver{r} }

// Mutation returns internal.MutationResolver implementation.
func (r *Resolver) Mutation() internal.MutationResolver { return &mutationResolver{r} }

// ProjectV2 returns internal.ProjectV2Resolver implementation.
func (r *Resolver) ProjectV2() internal.ProjectV2Resolver { return &projectV2Resolver{r} }

// Query returns internal.QueryResolver implementation.
func (r *Resolver) Query() internal.QueryResolver { return &queryResolver{r} }

// Repository returns internal.RepositoryResolver implementation.
func (r *Resolver) Repository() internal.RepositoryResolver { return &repositoryResolver{r} }

type issueResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type projectV2Resolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type repositoryResolver struct{ *Resolver }
