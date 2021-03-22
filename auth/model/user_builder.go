package model

type UserAddressBuilder struct {
	Address *Address
}

func (b *UserAddressBuilder) Name(v string) *UserAddressBuilder {
	b.Address.Name = v
	return b
}

func (b *UserAddressBuilder) Line1(v string) *UserAddressBuilder {
	b.Address.Line1 = v
	return b
}

func (b *UserAddressBuilder) Line2(v string) *UserAddressBuilder {
	b.Address.Line2 = v
	return b
}

func (b *UserAddressBuilder) Postcode(v string) *UserAddressBuilder {
	b.Address.Postcode = v
	return b
}

func (b *UserAddressBuilder) City(v string) *UserAddressBuilder {
	b.Address.City = v
	return b
}

func (b *UserAddressBuilder) CountryID(v uint) *UserAddressBuilder {
	b.Address.CountryID = v
	return b
}

type UserBuilder struct {
	User               *User
	UserAddressBuilder *UserAddressBuilder
}

func NewUserBuilder() *UserBuilder {
	u := &User{}
	return &UserBuilder{User: u, UserAddressBuilder: &UserAddressBuilder{}}
}

func (b *UserBuilder) Firstname(v string) *UserBuilder {
	b.User.Firstname = v
	return b
}

func (b *UserBuilder) Lastname(v string) *UserBuilder {
	b.User.Lastname = v
	return b
}

func (b *UserBuilder) Login(v string) *UserBuilder {
	b.User.Login = v
	return b
}

func (b *UserBuilder) Password(v string) *UserBuilder {
	b.User.SetPassword(v)
	return b
}

func (b *UserBuilder) Age(v uint8) *UserBuilder {
	b.User.Age = v
	return b
}

func (b *UserBuilder) AddAddress() *UserAddressBuilder {
	a := &Address{}
	b.User.Address = append(b.User.Address, a)
	b.UserAddressBuilder.Address = a

	return b.UserAddressBuilder
}

func (b *UserBuilder) Build() *User {
	return b.User
}
