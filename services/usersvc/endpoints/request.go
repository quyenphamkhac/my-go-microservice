package endpoints

import "github.com/quyenphamkhac/skoppi/services/usersvc/entities"

type GetByIdRequest struct {
	ID string `json:"id"`
}

type GetByIdResponse struct {
	User *entities.User `json:"data"`
	Err  error          `json:"error,omitempty"`
}

type GetAllRequest struct {
	Search string `json:"search"`
}

type GetAllResponse struct {
	Users []entities.User `json:"data"`
	Err   error           `json:"error,omitempty"`
}
