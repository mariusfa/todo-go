package user

import "todo/biz/user"

type UserDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func FromDomainList(users []user.User) []UserDTO {
	var userDTOs []UserDTO
	for _, user := range users {
		userDTOs = append(userDTOs, FromDomain(user))
	}
	return userDTOs
}

func FromDomain(user user.User) UserDTO {
	return UserDTO{
		ID:   user.ID,
		Name: user.Name,
	}
}
