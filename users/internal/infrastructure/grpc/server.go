package grpcserver 

import (
	"context"
	"fmt"

	userv1 "github.com/sverdejot/greeter/gen/go/proto/user/v1"
	"github.com/sverdejot/greeter/users/internal/domain"
)

type userService struct {
	repo domain.UserRepository
}

func NewGrpcUserService(repo domain.UserRepository) *userService {
	return &userService{repo}
}

func (us *userService) GetUserName(ctx context.Context, req *userv1.GetUserNameRequest) (*userv1.GetUserNameResponse, error) {
	if user, ok := us.repo.Find(ctx, int(req.Id)); ok {
		return &userv1.GetUserNameResponse{User: &userv1.User{Name: user.Name}}, nil
	}
	return nil, fmt.Errorf("user w/ id[%d] not found", req.Id)
}

