package bootstrap

import (
	"context"
	"fmt"
	"log"
	"net"

	userv1 "github.com/sverdejot/greeter/gen/go/proto/user/v1"
	"google.golang.org/grpc"
)


type userService struct {
	users map[int]string
}

func (us *userService) GetUserName(_ context.Context, req *userv1.GetUserNameRequest) (*userv1.GetUserNameResponse, error) {
	if name, ok := users[int(req.GetId())]; ok {
		return &userv1.GetUserNameResponse{User: &userv1.User{Name: name}}, nil
	}
	return nil, fmt.Errorf("user w/ id[%d] not found", req.Id)
}


func ListenGrpc(address string) {
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer lis.Close()

	grpcServer := grpc.NewServer()
	userv1.RegisterUserServiceServer(grpcServer, &userService{})

	log.Println("starting app")
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("unexpectedly stopped: %v", err)
	}
}
