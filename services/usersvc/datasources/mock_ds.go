package mock_ds

import (
	"context"

	"github.com/quyenphamkhac/skoppi/services/usersvc/entites"
)

type mockDatasource struct{}

func NewMockDatasource() *mockDatasource {
	return &mockDatasource{}
}

func (m *mockDatasource) GetUserById(ctx context.Context, id string) (*entites.User, error) {
	return &entites.User{}, nil
}

func (m *mockDatasource) GetUsers(ctx context.Context, search string) ([]entites.User, error) {
	return []entites.User{}, nil
}
