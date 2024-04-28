package bootstrap

import (
	"log"
	"net"
	"net/http"
	"time"

	userv1 "github.com/sverdejot/greeter/gen/go/proto/user/v1"
	"github.com/sverdejot/greeter/users/internal/domain"
	"github.com/sverdejot/greeter/users/internal/infrastructure/api"
	grpcserver "github.com/sverdejot/greeter/users/internal/infrastructure/grpc"
	"github.com/sverdejot/greeter/users/internal/infrastructure/storage/memory"
	"google.golang.org/grpc"
)

// need to close both connection when signal
func Run() {
	lis, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	repo := memory.NewInMemoryUserRepository(map[int]domain.User{
		1: {
			Id: 1,
			Name: "Samuel",
			Mail: "sverdejot@gmail.com",
			Age: 26,
			Status: domain.Active,
		},
	})

	grpcServer := grpc.NewServer()
	service := grpcserver.NewGrpcUserService(repo)
	userv1.RegisterUserServiceServer(grpcServer, service)

	log.Println("starting grpc server")
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("unexpectedly stopped: %v", err)
	}

	mux := api.AddRoutes(repo)

	server := http.Server{
		Addr: ":8082",

		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		IdleTimeout:       30 * time.Second,

		Handler: mux,
	}

	log.Println("starting app")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
