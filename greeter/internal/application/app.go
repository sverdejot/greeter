package application

import "github.com/sverdejot/greeter/greeter/internal/domain/users"

type App struct {
	Greeter *Greeter
}

func NewApp(userService users.UsersRepository) *App {
	return &App{
		Greeter: NewGreeter(userService),
	}
}
