package user

type UserDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func ToUsersDomain(userDTOs []UserDTO) []User {
	var users []User
	for _, userDTO := range userDTOs {
		users = append(users, toUserDomain(userDTO))
	}
	return users
}

func toUserDomain(userDTO UserDTO) User {
	return NewUser(userDTO.ID, userDTO.Name)
}
