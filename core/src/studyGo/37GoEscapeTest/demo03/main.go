package main

type User struct {
	Username *string
	Password string
	Age      int
}

func Call1(u *User) int{
	name := "bbb"
	u.Username = &name
	return u.Age * 20
}


func main() {
	a := "aaa"
	u := &User{&a, "123", 12}
	Call1(u)
}

