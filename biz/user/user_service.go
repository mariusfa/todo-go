package user

type UserService struct {
	userAdapter UserAdapterContract
}

func NewUserService(userAdapter UserAdapterContract) *UserService {
	return &UserService{userAdapter: userAdapter}
}

func (u *UserService) GetAll() ([]User, error) {
	return u.userAdapter.GetAll()
}
