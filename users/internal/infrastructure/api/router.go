package api

import (
	"net/http"

	"github.com/sverdejot/greeter/users/internal/application/uc"
	"github.com/sverdejot/greeter/users/internal/domain"
)


func AddRoutes(repo domain.UserRepository) http.Handler {
	mux := http.NewServeMux()

	uc := uc.NewUseCaseCreateUser(repo)
	handler := NewCreateUserHandler(uc)
	mux.HandleFunc("GET /users/{id}", handler)

	return mux
}
