package datasources

import (
	"context"

	"github.com/quyenphamkhac/skoppi/services/usersvc/entities"
)

type mockDatasource struct{}

func NewMockDatasource() *mockDatasource {
	return &mockDatasource{}
}

func (m *mockDatasource) GetUserById(ctx context.Context, id string) (*entities.User, error) {
	return &entities.User{
		FirstName: "Pham Khac",
		LastName:  "Quyen",
	}, nil
}

func (m *mockDatasource) GetUsers(ctx context.Context, search string) ([]entities.User, error) {
	return []entities.User{}, nil
}
