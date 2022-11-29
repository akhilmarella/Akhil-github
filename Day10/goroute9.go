package main

// import "fmt"

// func main() {
// 	chnl := make(chan int)
// 	go res(chnl)
// 	chase := <-chnl
// 	fmt.Println(chase)
// 	write := make(chan int)
// 	go product(write)
// 	for {
// 		v, ok := <-write
// 		if ok == false {
// 			break
// 		}
// 		fmt.Println(v, ok)
// 	}
// }
// func res(chnl chan int) {
// 	chase := 100
// 	chnl <- chase
// 	fmt.Println("res function")
// }
// func product(wrote chan int) {
// 	for i := 0; i <= 10; i++ {
// 		wrote <- i
// 	}
// 	close(wrote)
// }
