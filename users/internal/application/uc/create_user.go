package uc

import (
	"context"
	"fmt"

	"github.com/sverdejot/greeter/users/internal/application"
	"github.com/sverdejot/greeter/users/internal/domain"
)

type UseCaseCreateUser struct {
	repo domain.UserRepository
}

func NewUseCaseCreateUser(repo domain.UserRepository) *UseCaseCreateUser {
	return &UseCaseCreateUser{repo}
}

func (uc *UseCaseCreateUser) Run(ctx context.Context, id int, name, mail string, age int, status int) error {
	user, err := domain.NewUser(id, name, mail, age, status)
	if err != nil {
		return application.NewValidationError(err)
	}

	err = uc.repo.Save(ctx, user)
	if err != nil {
		return fmt.Errorf("cannot create user: %w", err)
	}

	return nil
}
