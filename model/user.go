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

func FindUserById(id string) *User {
	for _, u := range users {
		if u.id == id {
			return u
		}
	}
	return nil
}

func (u *User) EqualsPassword(pwd string) bool {
	return u.password == pwd
}

func (u *User) Secret() string {
	return u.secret
}
