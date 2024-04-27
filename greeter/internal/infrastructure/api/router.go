package api

import (
	"net/http"

	"github.com/sverdejot/greeter/greeter/internal/application"
)

func AddRoutes(app *application.App) http.Handler {
	mux := http.NewServeMux()

	greeterHandler := NewGreeterHandler(app.Greeter)

	mux.HandleFunc("GET /hello/{id}", greeterHandler)

	return mux
}
