package dto

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/domain/entity"
)

type UserInput struct {
	Id           int          `json:"id"`
	Name         string       `json:"name" validate:"nonzero"`
	Age          int          `json:"age" validate:"nonzero"`
	Email        string       `json:"email" validate:"nonzero"`
	Password     string       `json:"password" validate:"nonzero"`
	AddressInput AddressInput `json:"address"`
}

type AddressInput struct {
	Id       int    `json:"id"`
	Address  string `json:"address" validate:"nonzero"`
	City     string `json:"city" validate:"nonzero"`
	County   string `json:"county"`
	Country  string `json:"country" validate:"nonzero"`
	Postcode string `json:"postcode" validate:"nonzero"`
}

func (this *UserInput) MapUserToEntity() *entity.User {
	user := entity.User{}

	user.Id = this.Id
	user.Name = this.Name
	user.Age = this.Age
	user.Email = this.Email
	user.Password = this.Password
	user.Address = *this.AddressInput.MapAddressToEntity()

	return &user
}

func (this *AddressInput) MapAddressToEntity() *entity.Address {
	address := entity.Address{}

	address.Id = this.Id
	address.Address = this.Address
	address.City = this.City
	address.County = this.County
	address.Country = this.Country
	address.Postcode = this.Postcode

	return &address
}
