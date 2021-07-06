package endpoints

import "github.com/quyenphamkhac/skoppi/services/usersvc/entites"

type GetByIdRequest struct {
	ID string `json:"id"`
}

type GetByIdResponse struct {
	User *entites.User `json:"data"`
	Err  error         `json:"error,omitempty"`
}

type GetAllRequest struct {
	Search string `json:"search"`
}

type GetAllResponse struct {
	Users []entites.User `json:"data"`
	Err   error          `json:"error,omitempty"`
}
