package repo

import (
	"context"

	"github.com/quyenphamkhac/skoppi/pkg/usersvc/entities"
)

type UserRepo interface {
	GetUserById(ctx context.Context, id string) (*entities.User, error)
	GetUsers(ctx context.Context, search string) ([]entities.User, error)
}
