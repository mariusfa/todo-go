package user

type UserAdapterContract interface {
	GetAll() ([]User, error)
}
