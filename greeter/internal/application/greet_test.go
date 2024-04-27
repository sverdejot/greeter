package application_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sverdejot/greeter/greeter/internal/application"
)

type StubUsersRepository struct {
	Users map[int]string
}

func (s *StubUsersRepository) GetUserName(id int) (string, bool) {
	name, ok := s.Users[id]
	return name, ok
}

func TestGreeter(t *testing.T) {
	// given
	cases := map[string]struct {
		got   int
		want  string
		err   string
		users map[int]string
	}{
		"greet user":                {1, "Hello, Samuel!", "", map[int]string{1: "Samuel"}},
		"return error when no user": {2, "", "no user found for id 2", map[int]string{}},
	}

	for name, test_case := range cases {
		t.Run(name, func(t *testing.T) {
			// given
			repo := &StubUsersRepository{test_case.users}
			uc := application.NewGreeter(repo)

			// when
			msg, err := uc.Greet(test_case.got)

			// then
			assert.Equal(t, msg, test_case.want)
			switch test_case.err {
			case "":
				assert.NoError(t, err)
			default:
				assert.EqualError(t, err, test_case.err)
			}
		})
	}

}
