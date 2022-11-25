package main

import "fmt"

type Login struct {
	Email    string
	Password NewAcc
}
type NewAcc struct {
	Fname       string
	Lname       string
	Password    string
	ConfirmPass string
}
func main() {
	user := Login{}
	res1 := result(user)
	fmt.Println(res1)
	fmt.Println(&res1.Email)
}
func result(acc Login) Login {
	acc.Email = "qwerrty122"
	acc.Password.Lname = "babu"
	acc.Password.Fname = "akhil"
	acc.Password.ConfirmPass = "asdfgh1@123"
	return acc
}
