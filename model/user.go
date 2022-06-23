package model

type User struct {
	id       string
	password string
	secret   string
}

var users []*User = []*User{
	{
		id:       "hogehoge",
		password: "hogehoge", // password should be hashed
		secret:   "hello",    // secret should be encrypted
	},
}

func (u *User) First() *User {
	return users[0]
}
