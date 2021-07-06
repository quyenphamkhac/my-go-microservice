package interactor

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/quyenphamkhac/skoppi/services/usersvc"
	"github.com/quyenphamkhac/skoppi/services/usersvc/entites"
	"github.com/quyenphamkhac/skoppi/services/usersvc/repo"
)

type interactor struct {
	repo   repo.UserRepo
	logger log.Logger
}

func NewInteractor(repo repo.UserRepo, logger log.Logger) usersvc.UserService {
	return &interactor{
		repo:   repo,
		logger: logger,
	}
}

func (i *interactor) GetById(ctx context.Context, id string) (*entites.User, error) {
	logger := log.With(i.logger, "method", "GetById")
	user, err := i.repo.GetUserById(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return &entites.User{}, err
	}
	return user, nil
}

func (i *interactor) GetAll(ctx context.Context, search string) ([]entites.User, error) {
	logger := log.With(i.logger, "method", "GetAll")
	users, err := i.repo.GetUsers(ctx, search)
	if err != nil {
		level.Error(logger).Log("err", err)
		return []entites.User{}, err
	}
	return users, nil
}
