package main

import "fmt"

func res(write chan bool) {
	fmt.Println("hello world")
	write <- true
}
func main() {
	write := make(chan bool)
	go res(write)
	<-write
	fmt.Println("main function")
	read := make(chan int)
	go result(read)
	age := <-read
	fmt.Println("main function:", age)
}        
func result(read chan int) {
	age := 45
	read <- age
	fmt.Println("result function")

}
func res2(activity chan string) {

}
