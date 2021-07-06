package transports

import (
	"context"
	"errors"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/quyenphamkhac/skoppi/services/usersvc/endpoints"
)

var (
	ErrBadRouting = errors.New("router not found")
)

func NewHTTPTransport(svcEndpoints endpoints.Endpoints) http.Handler {
	var (
		r = mux.NewRouter()
	)
	r.Methods("GET").Path("/users/{id}").Handler(httptransport.NewServer(
		svcEndpoints.GetById,
		decodeGetByIdRequest,
		httptransport.EncodeJSONResponse,
	))
	return r
}

func decodeGetByIdRequest(ctx context.Context, r *http.Request) (req interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return endpoints.GetByIdRequest{ID: id}, nil
}
