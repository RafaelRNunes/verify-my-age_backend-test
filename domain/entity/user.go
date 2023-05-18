package entity

import (
	"github.com/google/uuid"
)

// User struct
type User struct {
	Id       uuid.UUID
	Name     string
	Age      int
	Email    string
	Password string
	Address  Address
}
