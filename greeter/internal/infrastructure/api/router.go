package api

import (
	"net/http"

	"github.com/sverdejot/greeter/internal/application"
)

func AddRoutes() http.Handler {
	mux := http.NewServeMux()

	users := map[int]string{
		1: "Samuel",
	}

	uc := application.NewGreeter(users)
	greeterHandler := NewGreeterHandler(uc)

	mux.HandleFunc("GET /hello/{id}", greeterHandler)

	return mux
}
