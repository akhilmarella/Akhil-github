package main

// import "fmt"

// func res(write chan bool) {
// 	write <- true
// 	fmt.Println("res function")
// }
// func main() {
// 	write := make(chan bool)
// 	go res(write)
// 	<-write
// 	fmt.Println("main function")
// 	name := make(chan string)
// 	go res1(name)
// 	names:=<-name
// 	fmt.Println(names)

// }
// func res1(name chan string) {
// 	fname := "akhil"
// 	lname := "babu"
// 	NAME := fname + lname
// 	name <- NAME
// 	fmt.Println("res1 function is:")
// }
