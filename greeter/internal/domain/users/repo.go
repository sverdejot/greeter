package users

type UsersRepository interface {
	GetUserName(id int) (string, bool)
}
