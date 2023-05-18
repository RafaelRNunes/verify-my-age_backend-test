package dto

import "github.com/RafaelRNunes/verify-my-age_backend-test/domain/entity"

type UserInput struct {
	Name     string
	Age      int
	Email    string
	Password string
	Address  AddressInput
}

type AddressInput struct {
	Address  string
	City     string
	County   string
	Country  string
	Postcode string
}

func (this *UserInput) MapUser() *entity.User {
	user := entity.User{}

	user.Name = this.Name
	user.Age = this.Age
	user.Email = this.Email
	user.Password = this.Password
	user.Address = *this.Address.MapAddress()

	return &user
}

func (this *AddressInput) MapAddress() *entity.Address {
	address := entity.Address{}

	address.Address = this.Address
	address.City = this.City
	address.County = this.County
	address.Country = this.Country
	address.Postcode = this.Postcode

	return &address
}
