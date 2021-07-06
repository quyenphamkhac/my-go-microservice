package usersvc

import (
	"context"

	"github.com/quyenphamkhac/skoppi/services/usersvc/entities"
)

type UserService interface {
	GetById(ctx context.Context, id string) (*entities.User, error)
	GetAll(ctx context.Context, search string) ([]entities.User, error)
}
