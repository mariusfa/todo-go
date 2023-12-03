package user

type UserServiceContract interface {
	GetAll() ([]User, error)
}

type UserService struct {
	userAdapter UserAdapterContract
}

func NewUserService(userAdapter UserAdapterContract) *UserService {
	return &UserService{userAdapter: userAdapter}
}

func (u *UserService) GetAll() ([]User, error) {
	return u.userAdapter.GetAll()
}

type UserServiceFake struct{}

func (u *UserServiceFake) GetAll() ([]User, error) {
	users := []User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
		{ID: 3, Name: "User 3"},
	}
	return users, nil
}
