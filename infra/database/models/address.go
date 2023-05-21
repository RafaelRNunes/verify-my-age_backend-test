package models

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/domain/entity"
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	Address  string
	City     string
	County   string
	Country  string
	Postcode string
	UserId   uint
}

func MapAddressToModel(address entity.Address) Address {
	return Address{
		Address:  address.Address,
		City:     address.City,
		County:   address.County,
		Country:  address.Country,
		Postcode: address.Postcode,
	}
}

func (this *Address) MapAddressToEntity() *entity.Address {
	return &entity.Address{
		Id:       int(this.ID),
		Address:  this.Address,
		City:     this.City,
		County:   this.County,
		Country:  this.Country,
		Postcode: this.Postcode,
	}
}
