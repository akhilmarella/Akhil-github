package main

import "fmt"

func main() {
	write := make(chan int)
	go res(write)
	//	age := 25
	write <- 25
	fmt.Println("main function")
	read := make(chan int)
	go result(read)

	for {
		v, ok := <-read
		if ok == false {
			break

		}
		fmt.Println("receive:", v, ok)
	}
}

func res(write chan int) {
	age := <-write
	fmt.Println(age)
	close(write)
}
func result(read chan int) {
	for i := 0; i <= 100; i += 10 {
		read <- i

	}
	close(read)
}
