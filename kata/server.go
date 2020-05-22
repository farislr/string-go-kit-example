package kata

import (
	"context"
	"net/http"

	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/uppercase").Handler(httpTransport.NewServer(
		endpoints.Uppercase,
		decodeUppercaseRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/count").Handler(httpTransport.NewServer(
		endpoints.Count,
		decodeCountRequest,
		encodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
