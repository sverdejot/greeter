package uc_test

import (
	"context"
	"testing"

	"github.com/sverdejot/greeter/users/internal/application/uc"
	"github.com/sverdejot/greeter/users/internal/application/uc/mothers"
	"github.com/sverdejot/greeter/users/internal/domain"
)

type StubUserRepository struct {
	
}

func (st *StubUserRepository) Save(context.Context, domain.User) error {
	return nil
}

func (st *StubUserRepository) Find(_ context.Context, id int) (domain.User, bool) {
	return domain.User{}, false
}

func TestCreateUser(t *testing.T) {
	// given
	stub := &StubUserRepository{}

	uc := uc.NewUseCaseCreateUser(stub)

	user := mothers.GenerateUser(
		mothers.WithId(1),
		mothers.WithAge(18),
		mothers.WithName("Samuel Verdejo de Toro"),
		mothers.WithMail("contacto@sverdejot.dev"),
		mothers.WithStatus(domain.Active),
	)

	// when
	got := uc.Run(
		context.TODO(),
		user.Id,
		user.Name,
		user.Mail,
		user.Age,
		int(user.Status),
	)

	// then
	if got != nil {
		t.Errorf("got %v want nothing", got)
	}
}
