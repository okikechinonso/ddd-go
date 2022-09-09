package driver

import "flyorg/entity"

type Driver struct {
	user *entity.User
}

func NewDriver(user entity.User) (*Driver, error){
	return nil, nil
}