package main

import "fmt"

func res(marry chan int) {
	age := 25
	marry <- age
	fmt.Println("res")

}

func main() {
	marry := make(chan int)
	go res(marry)
	age := <-marry
	fmt.Println("main func:", age)
	carry := make(chan int)
	go res2(carry)
	for {
		v, ok := <-carry
		if ok == false {
			break
		}
		fmt.Println("receive:", v, ok)
	}
	write := make(chan int)
	go res1(write)
	for i := range write {
		fmt.Println("rec:", i)
	}
	read := make(chan int)
	go result(read)
	read <- 45

}
func res2(carry chan int) {
	for i := 0; i < 5; i++ {
		carry <- i
	}
	close(carry)
}
func res1(write chan<- int) {
	for i := 0; i <= 10; i += 2 {
		write <- i
	}
	close(write)
}
func result(read chan int) {
	age:=<-read
	fmt.Println(age)
}
