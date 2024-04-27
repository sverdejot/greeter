package application

import (
	"fmt"

	"github.com/sverdejot/greeter/greeter/internal/domain/users"
)

type Greeter struct {
	repo users.UsersRepository
}

func NewGreeter(users users.UsersRepository) *Greeter {
	return &Greeter{users}
}

func (uc *Greeter) Greet(id int) (string, error) {
	if name, ok := uc.repo.GetUserName(id); ok {
		return fmt.Sprintf("Hello, %s!", name), nil
	}
	return "", fmt.Errorf("no user found for id %d", id)
}
