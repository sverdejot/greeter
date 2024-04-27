package grpcclient

import (
	"context"
	"log"

	userv1 "github.com/sverdejot/greeter/gen/go/proto/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcUserService struct {
	client userv1.UserServiceClient
}

func (us *GrpcUserService) GetUserName(id int) (string, bool) {
	res, err := us.client.GetUserName(context.TODO(), &userv1.GetUserNameRequest{Id: int32(id)})
	if err != nil {
		log.Printf("cannot get user name: %v", err)
		return "", false
	}
	return res.GetUser().Name, true
}

func NewGrpcUserService() (*GrpcUserService, func(), error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:8081", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	client := userv1.NewUserServiceClient(conn)

	return &GrpcUserService{client}, func() { conn.Close() }, nil
}
