package user

type UserAdapterFake struct{}

func NewUserAdapterFake() *UserAdapterFake {
	return &UserAdapterFake{}
}

func (u *UserAdapterFake) GetAll() ([]User, error) {
	return []User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}, nil
}
