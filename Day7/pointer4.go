package main

// import "fmt"

// type Login struct {
// 	Email string
// 	Pass  string
// }
// type Newacc struct {
// 	Fname string
// 	Lname string
// }


// func main() {
// 	name := "akhil"
// 	result(name)

// 	Naccount := Newacc{}
// 	res1 := result1(&Naccount)
// 	fmt.Println("address of fname", &res1.Fname)
// 	fmt.Println("name of fname", res1.Fname)
// 	fmt.Printf("address of lname  %s\n", res1.Lname)
// 	fmt.Println("name of lname ", &res1.Lname)
// 	res2 := result2()
// 	fmt.Println("addess of res2 ", &res2)
// 	fmt.Println("names of res2", res2)
// 	fmt.Println("address of res2(email & password)", &res2.Email, "\n", &res2.Pass)
// 	fmt.Println("name of res2(email & password)", res2.Email, "\n", res2.Pass)
// }
// func result(name string) {
// 	fmt.Println(name)
// }
// func result1(account *Newacc) *Newacc {
// 	account.Fname = "akhil"
// 	account.Lname = "babu"
// 	return account
// }
// func result2() Login {
// 	account := Login{}
// 	account.Email = "qwert@12"
// 	account.Pass = "asdf12"
// 	return account
// }
// func