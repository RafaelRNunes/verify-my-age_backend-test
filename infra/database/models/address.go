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

func (this *Address) MapAddressToModel(address entity.Address) {
	this.Address = address.Address
	this.City = address.City
	this.County = address.County
	this.Country = address.Country
	this.Postcode = address.Postcode
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
