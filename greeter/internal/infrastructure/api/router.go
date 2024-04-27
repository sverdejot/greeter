package api

import (
	"net/http"

	"github.com/sverdejot/greeter/greeter/internal/application"
	services "github.com/sverdejot/greeter/greeter/internal/infrastructure/http"
)

func AddRoutes() http.Handler {
	mux := http.NewServeMux()


	us := services.NewUserService()

	uc := application.NewGreeter(us)
	greeterHandler := NewGreeterHandler(uc)

	mux.HandleFunc("GET /hello/{id}", greeterHandler)

	return mux
}
