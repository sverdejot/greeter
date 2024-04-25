package application

import (
	"fmt"
)

type Greeter struct {
	users map[int]string
}

func NewGreeter(users map[int]string) *Greeter {
	return &Greeter{users}
}

func (uc *Greeter) Greet(id int) (string, error) {
	if name, ok := uc.users[id]; ok {
		return fmt.Sprintf("Hello, %s!", name), nil
	}
	return "", fmt.Errorf("no user found for id %d", id)
}
