package user

type User struct {
	ID   int64
	Name string
}

func NewUser(id int64, name string) User {
	return User{
		ID:   id,
		Name: name,
	}
}
