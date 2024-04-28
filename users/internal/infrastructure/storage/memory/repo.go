package memory

import (
	"context"
	"sync"

	"github.com/sverdejot/greeter/users/internal/domain"
)

var rwLock sync.RWMutex

type inMemoryUserRepository struct {
	users map[int]domain.User
}

func NewInMemoryUserRepository(users map[int]domain.User) *inMemoryUserRepository {
	return &inMemoryUserRepository{users}
}

func (r *inMemoryUserRepository) Save(_ context.Context, user domain.User) error {
	if lock := rwLock.TryLock(); lock {
		r.users[user.Id] = user
	}
	return nil
}

func (r *inMemoryUserRepository) Find(_ context.Context, id int) (domain.User, bool) {
	if lock := rwLock.TryRLock(); lock {
		if user, ok := r.users[id]; ok {
			return user, true
		}
	}
	return domain.User{}, false
}
