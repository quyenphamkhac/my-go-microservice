package repo

import (
	"context"

	"github.com/quyenphamkhac/skoppi/services/usersvc/entites"
)

type UserRepo interface {
	GetUserById(ctx context.Context, id string) (*entites.User, error)
	GetUsers(ctx context.Context) ([]entites.User, error)
}
