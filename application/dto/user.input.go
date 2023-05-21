package dto

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/domain/entity"
)

type UserInput struct {
	Id           int          `json:"id"`
	Name         string       `json:"name"`
	Age          int          `json:"age"`
	Email        string       `json:"email"`
	Password     string       `json:"password"`
	AddressInput AddressInput `json:"address"`
}

type AddressInput struct {
	Id       int    `json:"id"`
	Address  string `json:"address"`
	City     string `json:"city"`
	County   string `json:"county"`
	Country  string `json:"country"`
	Postcode string `json:"postcode"`
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
