package mock_ds

import (
	"context"

	"github.com/quyenphamkhac/skoppi/services/usersvc/entites"
)

type mockDatasource struct{}

func NewMockDatasource() *mockDatasource {
	return &mockDatasource{}
}

func (m *mockDatasource) GetById(ctx context.Context, id string) (*entites.User, error) {
	return &entites.User{}, nil
}

func (m *mockDatasource) GetAll(ctx context.Context, search string) ([]entites.User, error) {
	return []entites.User{}, nil
}
