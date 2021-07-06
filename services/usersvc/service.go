package usersvc

import (
	"context"

	"github.com/quyenphamkhac/skoppi/services/usersvc/entites"
)

type UserService interface {
	GetById(ctx context.Context, id string) (*entites.User, error)
	GetAll(ctx context.Context, search string) ([]entites.User, error)
}
