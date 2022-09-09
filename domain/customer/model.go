package customer

import "flyorg/entity"

type Customer struct{
	user *entity.User
}

func NewCustomer(user entity.User) (*Customer, error){
	return nil, nil
}