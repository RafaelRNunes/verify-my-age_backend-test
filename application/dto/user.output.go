package dto

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/domain/entity"
)

type UserOutput struct {
	Id            int           `json:"id"`
	Name          string        `json:"name"`
	Age           int           `json:"age"`
	Email         string        `json:"email"`
	Password      string        `json:"password"`
	AddressOutput AddressOutput `json:"address"`
}

type AddressOutput struct {
	Id       int    `json:"id"`
	Address  string `json:"address"`
	City     string `json:"city"`
	County   string `json:"county"`
	Country  string `json:"country"`
	Postcode string `json:"postcode"`
}

func MapUserToDto(user *entity.User) *UserOutput {
	userOutput := UserOutput{}

	userOutput.Id = user.Id
	userOutput.Name = user.Name
	userOutput.Age = user.Age
	userOutput.Email = user.Email
	userOutput.Password = user.Password
	userOutput.AddressOutput = *MapAddressToDto(user.Address)

	return &userOutput
}

func MapAddressToDto(address entity.Address) *AddressOutput {
	addressOutput := AddressOutput{}

	addressOutput.Id = address.Id
	addressOutput.Address = address.Address
	addressOutput.City = address.City
	addressOutput.County = address.County
	addressOutput.Country = address.Country
	addressOutput.Postcode = address.Postcode

	return &addressOutput
}
