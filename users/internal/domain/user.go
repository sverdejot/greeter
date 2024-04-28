package domain

import (
	"context"
	"errors"
)
 
type UserStatus int

const (
	Undefined UserStatus = iota
	Active
	Inactive
	Suspended
	end
)

func parseStatus(status int) (UserStatus, error) {
	if (status >= int(end)) {
		return Undefined, errors.New("invalid user status")
	}
	return UserStatus(status), nil
}

type User struct {
	Id int
	Name string
	Mail string
	Age int
	Status UserStatus
}


var ErrInvalidAge = errors.New("an user's age cannot be below zero")

func NewUser(id int, name, mail string, age int, status int) (User, error) {
	if (age < 0) {
		return User{}, ErrInvalidAge
	}

	st, err := parseStatus(status)
	if err != nil {
		return User{}, err
	}

	return User{
		Id: id,
		Name: name,
		Mail: mail,
		Age: age,
		Status: st,
	}, nil
}

type UserRepository interface {
	Save(context.Context, User) error
	Find(context.Context, int) (User, bool)
}
