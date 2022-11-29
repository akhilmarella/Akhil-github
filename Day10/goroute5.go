package main

import "fmt"

func res(read chan bool) {
	read <- true
	fmt.Println("ress")
}

func main() {
	read := make(chan bool)
	go res(read)
	<-read
	fmt.Println("main function")
	name := make(chan string)
	go res1(name)
	Name := <-name
	fmt.Println(Name)
	write := make(chan int)
	go res2(write)
	for {
		v, ok := <-write
		if ok==false{
			break
		}
fmt.Println("receive:",v)
	}
}

func res1(name chan string) {
	Name := "marella"
	name <- Name
	fmt.Println("res1 function")
}
func res2(write chan int) {
	for i := 0; i < 10; i += 2 {
		write <- i
	}
	close(write)
}
