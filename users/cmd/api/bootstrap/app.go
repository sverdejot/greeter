package bootstrap

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	userv1 "github.com/sverdejot/greeter/gen/go/proto/user/v1"
	"github.com/sverdejot/greeter/users/internal/domain"
	"github.com/sverdejot/greeter/users/internal/infrastructure/api"
	grpcserver "github.com/sverdejot/greeter/users/internal/infrastructure/grpc"
	"github.com/sverdejot/greeter/users/internal/infrastructure/storage/memory"
	"google.golang.org/grpc"
)

var users map[int]domain.User = map[int]domain.User{
	1: {
		Id: 1,
		Name: "Samuel",
		Mail: "sverdejot@gmail.com",
		Age: 26,
		Status: domain.Active,
	},
}

// need to close both connection when signal
func Run() <- chan error {
	repo := memory.NewInMemoryUserRepository(users)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	errCh := make(chan error)
	lis, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("error opening tcp sock: %v", err)
	}

	grpcServer := newGrpcServer(repo)
	go func() {
		if err := grpcServer.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			log.Fatalf("error openning grpc server: %v", err)
		}
	}()

	server := newHttpServer(repo)
	go func() {
		log.Println("starting app")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// graceful shutdown
	go func() {
		<- ctx.Done()

		ctxTimeout, cancel := context.WithTimeout(ctx, 5 * time.Second)
		defer func() {
			lis.Close()
			stop()
			cancel()
			close(errCh)
		}()

		log.Println("gracefully shutting down services")

		grpcServer.Stop()
		if err := server.Shutdown(ctxTimeout); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}

		log.Println("succesfully closed all services")
	}()

	return errCh
}

func newGrpcServer(repo domain.UserRepository) *grpc.Server {
	grpcServer := grpc.NewServer()
	service := grpcserver.NewGrpcUserService(repo)
	userv1.RegisterUserServiceServer(grpcServer, service)

	return grpcServer
}

func newHttpServer(repo domain.UserRepository) *http.Server {
	mux := api.AddRoutes(repo)

	server := &http.Server{
		Addr: ":8082",

		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		IdleTimeout:       30 * time.Second,

		Handler: mux,
	}
	return server
}
