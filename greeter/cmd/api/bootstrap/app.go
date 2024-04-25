package bootstrap

import (
	"log"
	"net/http"
	"time"

	"github.com/sverdejot/greeter/internal/infrastructure/api"
)

func Run() {
	handler := api.AddRoutes()

	server := http.Server{
		Addr: ":8080",

		IdleTimeout: 30 * time.Second,
		WriteTimeout: 10 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		ReadTimeout: 10 * time.Second,

		Handler: handler,
	}

	log.Println("starting app")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
