package mothers

import "github.com/sverdejot/greeter/users/internal/domain"

func GenerateUser(opts ...func(domain.User) domain.User) domain.User {
	var user domain.User
	for _, f := range opts {
		user = f(user)
	}
	return user
}

func WithId(id int) func(domain.User) domain.User {
	return func(u domain.User) domain.User {
		u.Id = id 
		return u
	}
}

func WithAge(age int) func(domain.User) domain.User {
	return func(u domain.User) domain.User {
		u.Age = age
		return u
	}
}

func WithName(name string) func(domain.User) domain.User {
	return func(u domain.User) domain.User {
		u.Name = name
		return u
	}
}

func WithMail(mail string) func(domain.User) domain.User {
	return func(u domain.User) domain.User {
		u.Mail = mail
		return u
	}
}

func WithStatus(status domain.UserStatus) func(domain.User) domain.User {
	return func(u domain.User) domain.User {
		u.Status = status
		return u
	}
}




