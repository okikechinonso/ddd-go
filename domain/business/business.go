package business

import (
	"flyorg/entity"
)

type Business struct {
	user *entity.User
}

func NewDriver(user entity.User) (*Business, error) {
	return nil, nil
}
