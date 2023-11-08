package user

import "todo/api"

type UserAdapter struct{}

func (u *UserAdapter) GetAll() ([]User, error) {
	var users []UserDTO
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	if err := api.FetchJson(&users, "https://jsonplaceholder.typicode.com/users", headers); err != nil {
		return nil, err
	}
	return ToUsersDomain(users), nil
}

func NewUserAdapter() UserAdapter {
	return UserAdapter{}
}
