package entity

import "github.com/google/uuid"

type Address struct {
	Id       uuid.UUID
	Address  string
	City     string
	County   string
	Country  string
	Postcode string
}
