package main

import "fmt"

func main() {
	read := make(chan int)
	go result(read)
	<-read
	for {
		v, yes := <-read
		if yes == false {
			break
		}
		fmt.Println(v, yes)
	}
	write := make(chan string)
	go res(write)
	write <- "akhil"
	ch := make(chan string, 4)
	ch <- "ab"
	ch <- "asdf"
	ch <- "qwerty"
	ch <- "akhil babu"
	fmt.Println(len(ch))

}
func result(learn chan int) {
	for i := 0; i <= 10; i++ {
		learn <- i
	}
	close(learn)
}
func res(write chan string) {
	fmt.Println(<-write + "babu")
}
