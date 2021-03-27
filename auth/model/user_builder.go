package model

type UserBuilder struct {
	User *User
}

func NewUserBuilder() *UserBuilder {
	u := &User{}
	return &UserBuilder{User: u}
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

func (b *UserBuilder) Build() User {
	return *b.User
}
