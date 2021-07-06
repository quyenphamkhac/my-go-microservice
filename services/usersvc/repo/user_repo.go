package repo

import (
	"context"

	"github.com/quyenphamkhac/skoppi/services/usersvc/entities"
)

type UserRepo interface {
	GetUserById(ctx context.Context, id string) (*entities.User, error)
	GetUsers(ctx context.Context, search string) ([]entities.User, error)
}
