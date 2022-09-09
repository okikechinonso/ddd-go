package entity

import "time"


type User struct {
	FirstName string    `json:"first-name"`
	LastName  string    `json:"last-name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Verified  bool      `json:"verified"`
	CreatedAt time.Time `json:"created-at"`
	UpdateAt  time.Time `json:"update-at"`
}
