package dto

type UserOutput struct {
	Id       string
	Name     string
	Age      int
	Email    string
	Password string
	Address  AddressOutput
}

type AddressOutput struct {
	Id       string
	Address  string
	City     string
	County   string
	Country  string
	Postcode string
}
