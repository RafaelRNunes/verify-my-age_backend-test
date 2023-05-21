package models

import (
	"github.com/RafaelRNunes/verify-my-age_backend-test/domain/entity"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Email    string
	Password string
	Address  Address
}

func (this *User) MapUserToModel(user entity.User) {
	this.Name = user.Name
	this.Age = user.Age
	this.Email = user.Email
	this.Password = user.Password
	this.Address.MapAddressToModel(user.Address)
}

func (this *User) MapUserToEntity() *entity.User {
	return &entity.User{
		Id:       int(this.ID),
		Name:     this.Name,
		Age:      this.Age,
		Email:    this.Email,
		Password: this.Password,
		Address:  *this.Address.MapAddressToEntity(),
	}
}
