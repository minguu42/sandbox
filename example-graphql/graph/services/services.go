package services

import "github.com/volatiletech/sqlboiler/v4/boil"

type Services interface {
	UserService
	RepositoryService
}

type services struct {
	*userService
	*repositoryService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:       &userService{exec: exec},
		repositoryService: &repositoryService{exec: exec},
	}
}
