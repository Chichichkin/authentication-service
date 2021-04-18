package model

import "time"

type IUsers interface {
	CreateTableIfNotExists() error
	SelectById(id int64) (*User, error)
	SelectByEmail(email string) (*User, error)
	Insert(info *User) (*User, error)
	UpdatePassword(id int64, password string) (*User, error)
	UpdateEmail(id int64, email string) (*User, error)
	Delete(id int64) error
}

type User struct {
	Id        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Role      Role      `json:"role"`
	Status    Status    `json:"status"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
}

type Role int

const (
	Default Role = iota
	Moderator
)

func (r Role) String() string {
	return [...]string{"User", "Moderator"}[r]
}

type Status int

const (
	NotConfirmed Status = iota
	Confirmed
	Banned
)

func (s Status) String() string {
	return [...]string{"NotConfirmed", "Confirmed", "Banned"}[s]
}
