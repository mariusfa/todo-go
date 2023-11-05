package adapters

import "todo/biz/user"

type Adapters struct {
	UserAdapter user.UserAdapterContract
}

func NewAdapterFakes() *Adapters {
	return &Adapters{
		UserAdapter: &user.UserAdapterFake{},
	}
}
