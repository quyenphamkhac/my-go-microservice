package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/quyenphamkhac/skoppi/pkg/usersvc"
)

type Endpoints struct {
	GetById endpoint.Endpoint
	GetAll  endpoint.Endpoint
}

func MakeEndpoints(svc usersvc.UserService) Endpoints {
	return Endpoints{
		GetById: makeGetByIdEndpoint(svc),
		GetAll:  makeGetAllEndpoint(svc),
	}
}

func makeGetByIdEndpoint(svc usersvc.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetByIdRequest)
		res, err := svc.GetById(ctx, req.ID)
		if err != nil {
			return GetByIdResponse{}, err
		}
		return GetByIdResponse{User: res, Err: nil}, nil
	}
}

func makeGetAllEndpoint(svc usersvc.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetAllRequest)
		res, err := svc.GetAll(ctx, req.Search)
		return GetAllResponse{Users: res, Err: err}, err
	}
}
