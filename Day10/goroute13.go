package main

// import "fmt"

// func main() {
// 	marry := make(chan int)
// 	go res(marry)
// 	for v := range marry {
// 		fmt.Println("receiving:",v)
// 	}

// 	carry := make(chan int)
// 	go res1(carry)
// 	for v := range carry {
// 		fmt.Println("received:",v)
// 	}
// }
// func res(marry chan int) {
// 	for i := 0; i < 10; i++ {
// 		marry <- i
// 	}
// 	close(marry)
// }
// func res1(carry chan int) {
// 	for i := 0; i < 10; i++ {
// 		carry <- i
// 	}
// 	close(carry)
// }