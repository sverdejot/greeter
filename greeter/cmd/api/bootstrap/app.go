package bootstrap

import (
	"log"
	"net/http"
	"time"

	"github.com/sverdejot/greeter/greeter/internal/application"
	"github.com/sverdejot/greeter/greeter/internal/infrastructure/api"
	grpcclient "github.com/sverdejot/greeter/greeter/internal/infrastructure/grpc"
)

func Run() {
	us, _, err := grpcclient.NewGrpcUserService()
	if err != nil {
		log.Fatalf("cannot instantaite grpc service: %v", err)
	}

	app := application.NewApp(us)

	handler := api.AddRoutes(app)

	server := http.Server{
		Addr: ":8080",

		IdleTimeout:       30 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		ReadTimeout:       10 * time.Second,

		Handler: handler,
	}

	log.Println("starting app")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
