package trasport

import (
	"context"
	"encoding/json"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/mohashari/kit-user/domain"
	userEndpoint "github.com/mohashari/kit-user/endpoint"
)

//MakeHandler ...
func MakeHandler(userEndpoint userEndpoint.UserEndpoint) http.Handler {
	r := mux.NewRouter()

	r.Methods("POST").Path("/api/v1/user").Handler(kithttp.NewServer(
		userEndpoint.CreateUserEndpoint,
		decodeUserRequest,
		endcode,
	))
	r.Methods("GET").Path("/api/v1/user").Handler(kithttp.NewServer(
		userEndpoint.GetUserEndpoint,
		decodeUserGetRequest,
		endcode,
	))

	return r
}

func decodeUserRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req domain.User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeUserGetRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return request, nil
}
func endcode(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
